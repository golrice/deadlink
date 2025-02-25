package utils

import (
	"net/url"
	"strings"
)

// NormalizeURL 规范化 URL 地址
func NormalizeURL(baseURL, href string) string {
	if strings.HasPrefix(href, "javascript:") || strings.HasPrefix(href, "mailto:") {
		return ""
	}

	if strings.HasPrefix(href, "/") {
		baseURL = strings.TrimRight(baseURL, "/")
		return baseURL + href
	}

	parsedURL, err := url.Parse(href)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return ""
	}

	return href
}

// ExtractDomain 提取 URL 的域名部分
func ExtractDomain(urlStr string) string {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return ""
	}
	return parsedURL.Hostname()
}
