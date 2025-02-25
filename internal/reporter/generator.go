package reporter

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/golrice/deadlink/internal/models"
)

type ReportGenerator struct{}

func NewReportGenerator() *ReportGenerator {
	return &ReportGenerator{}
}

func (r *ReportGenerator) Generate(result *models.CheckResult, outputPath string) error {
	if outputPath == "" {
		// 输出到控制台
		fmt.Println(result.Format())
		return nil
	}

	switch {
	case strings.HasSuffix(outputPath, ".csv"):
		return r.exportToCSV(result, outputPath)
	case strings.HasSuffix(outputPath, ".json"):
		return r.exportToJSON(result, outputPath)
	default:
		return fmt.Errorf("unsupported file format: %s", outputPath)
	}
}

func (r *ReportGenerator) exportToCSV(result *models.CheckResult, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"URL", "Status Code", "Error", "Checked At"}
	if err := writer.Write(headers); err != nil {
		return err
	}

	for _, link := range result.Links {
		row := []string{
			link.URL,
			fmt.Sprintf("%d", link.StatusCode),
			fmt.Sprintf("%v", link.Error),
			link.CheckedAt.String(),
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}

func (r *ReportGenerator) exportToJSON(result *models.CheckResult, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}
