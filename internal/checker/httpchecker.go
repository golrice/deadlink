package checker

import (
	"net/http"
	"sync"
	"time"

	"github.com/golrice/deadlink/internal/models"
)

type HTTPChecker struct {
	client *http.Client
}

func NewLinkChecker() *HTTPChecker {
	return &HTTPChecker{
		client: &http.Client{
			Timeout: 10 * time.Second, // 设置超时时间为 10 秒
		},
	}
}

func (c *HTTPChecker) Check(links []string) []models.Link {
	var wg sync.WaitGroup
	results := make([]models.Link, len(links))

	for i, link := range links {
		wg.Add(1)
		go func(i int, link string) {
			curTime := time.Now()

			defer wg.Done()

			resp, err := c.client.Head(link) // 使用 HEAD 方法减少流量
			if err != nil {
				results[i] = models.Link{URL: link, Error: err}
				return
			}
			defer resp.Body.Close()

			results[i] = models.Link{
				URL:        link,
				StatusCode: resp.StatusCode,
				CheckedAt:  curTime,
			}
		}(i, link)
	}

	wg.Wait()
	return results
}
