package scraper

import (
	"intelygenz/pkg/logs"
	"intelygenz/pkg/scraper/api"
)

func StartScraperProcess() {
	logs.Logger.Info("[scraper] starting scraping function ")
	api.ScrapFromAPI()
}
