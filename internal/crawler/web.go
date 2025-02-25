package crawler

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/golrice/deadlink/internal/utils"
)

type WebCrawler struct {
	StartURL   string              // 起始 URL
	DomainName string              // 需要递归查询的域名
	MaxDepth   int                 // 最大爬取深度
	Visited    map[string]struct{} // 已访问的链接集合
	Mutex      sync.Mutex          // 用于并发控制
}

func NewCrawler(startURL string, maxDepth int) *WebCrawler {
	return &WebCrawler{
		StartURL: startURL,
		MaxDepth: maxDepth,
		Visited:  make(map[string]struct{}),
	}
}

func (c *WebCrawler) Crawl() ([]string, error) {
	meta, err := url.Parse(c.StartURL)
	if err != nil {
		return nil, utils.WrapError(err, "parse start url")
	}
	c.DomainName = meta.Hostname()

	var links []string
	var wg sync.WaitGroup

	wg.Add(1)
	go c.crawlPage(c.StartURL, 0, &links, &wg)

	wg.Wait()

	return links, nil
}

func (c *WebCrawler) crawlPage(url string, depth int, links *[]string, wg *sync.WaitGroup) {
	defer wg.Done()

	if !strings.HasSuffix(url, c.DomainName) {
		return
	}

	if depth > c.MaxDepth {
		return
	}

	c.Mutex.Lock()
	if _, ok := c.Visited[url]; ok {
		c.Mutex.Unlock()
		return
	}
	c.Visited[url] = struct{}{}
	c.Mutex.Unlock()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to fetch %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Printf("Failed to parse %s: %v\n", url, err)
		return
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
			return
		}

		fullURL := utils.NormalizeURL(url, href)
		if fullURL == "" {
			return
		}

		*links = append(*links, fullURL)

		wg.Add(1)
		go c.crawlPage(fullURL, depth+1, links, wg)
	})
}
