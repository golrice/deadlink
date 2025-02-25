package crawler

type Crawler interface {
	Crawl() ([]string, error)
}
