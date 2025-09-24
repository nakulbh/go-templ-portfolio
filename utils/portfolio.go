package utils

import (
	"encoding/json"
	"os"

	"github.com/nakulbh/go-templ-portfolio/models"
)

func LoadPortfolioData(filepath string) (*models.PortfolioData, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var portfolio models.PortfolioData
	if err := json.Unmarshal(data, &portfolio); err != nil {
		return nil, err
	}

	return &portfolio, nil
}
