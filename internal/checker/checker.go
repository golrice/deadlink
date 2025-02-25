package checker

import "github.com/golrice/deadlink/internal/models"

type LinkChecker interface {
	Check(links []string) []models.Link
}
