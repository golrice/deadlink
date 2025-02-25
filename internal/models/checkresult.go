package models

import (
	"strings"
	"time"
)

type CheckResult struct {
	Links         []Link `json:"links"`          // 所有链接的检查结果
	BrokenLinks   []Link `json:"broken_links"`   // 死链列表
	ValidLinks    []Link `json:"valid_links"`    // 有效链接列表
	TotalLinks    int    `json:"total_links"`    // 总链接数
	BrokenCount   int    `json:"broken_count"`   // 死链数量
	ValidCount    int    `json:"valid_count"`    // 有效链接数量
	CheckDuration string `json:"check_duration"` // 检查耗时
}

func NewCheckResult(links []Link, duration time.Duration) *CheckResult {
	result := &CheckResult{
		Links:         links,
		TotalLinks:    len(links),
		CheckDuration: duration.String(),
	}

	for _, link := range links {
		if link.IsBroken() {
			result.BrokenLinks = append(result.BrokenLinks, link)
		} else {
			result.ValidLinks = append(result.ValidLinks, link)
		}
	}
	result.BrokenCount = len(result.BrokenLinks)
	result.ValidCount = len(result.ValidLinks)

	return result
}

func (r *CheckResult) Format() string {
	var sb strings.Builder

	sb.WriteString("Dead Link Check Report\n")
	sb.WriteString("=======================\n")
	sb.WriteString("Total Links: " + string(r.TotalLinks) + "\n")
	sb.WriteString("Valid Links: " + string(r.ValidCount) + "\n")
	sb.WriteString("Broken Links: " + string(r.BrokenCount) + "\n")
	sb.WriteString("Check Duration: " + r.CheckDuration + "\n\n")

	if r.BrokenCount > 0 {
		sb.WriteString("Broken Links:\n")
		for _, link := range r.BrokenLinks {
			sb.WriteString("- " + link.String() + "\n")
		}
	} else {
		sb.WriteString("No broken links found.\n")
	}

	return sb.String()
}
