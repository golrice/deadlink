package models

import (
	"fmt"
	"time"
)

type Link struct {
	URL         string    `json:"url"`          // 链接地址
	StatusCode  int       `json:"status_code"`  // HTTP 响应状态码
	Error       error     `json:"error"`        // 错误信息（如网络错误）
	CheckedAt   time.Time `json:"checked_at"`   // 检查时间
	RedirectURL string    `json:"redirect_url"` // 如果有重定向，记录最终目标 URL
}

func (l *Link) IsBroken() bool {
	return l.StatusCode >= 400 || l.Error != nil
}

func (l *Link) String() string {
	if l.IsBroken() {
		if l.Error != nil {
			return l.URL + " - Error: " + l.Error.Error()
		}
		return l.URL + " - Status Code: " + fmt.Sprint(l.StatusCode)
	}
	return l.URL + " - OK"
}

func UniqueLinks(links []string) []string {
	uniqueMap := make(map[string]struct{})
	var uniqueList []string

	for _, link := range links {
		if _, exists := uniqueMap[link]; !exists {
			uniqueMap[link] = struct{}{}
			uniqueList = append(uniqueList, link)
		}
	}
	return uniqueList
}
