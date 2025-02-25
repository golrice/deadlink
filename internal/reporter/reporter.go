package reporter

import "github.com/golrice/deadlink/internal/models"

type Reporter interface {
	Generate(result *models.CheckResult, outputPath string) error // 生成报告
}
