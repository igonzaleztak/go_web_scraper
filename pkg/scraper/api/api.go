package api

import (
	"encoding/json"
	"fmt"
	"intelygenz/pkg/config"
	"intelygenz/pkg/logs"
	"io"
	"net/http"
)

func ScrapFromAPI() (any, error) {
	logs.Logger.Debugf("[scraper.ScrapFromAPI] starting API scraper")
	api := config.AppConfig.APIURL
	logs.Logger.Debugf("[scraper.ScrapFromAPI] got hacker news API host: %s", api)

	param := "newstories"
	uri := fmt.Sprintf("%s/%s.json?print=pretty", api, param)
	fmt.Println(uri)

	resp, err := http.Get(uri)
	if err != nil {
		return nil, fmt.Errorf("failed to execute get to '%s': %v", uri, err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	response := make([]int, 0)
	_ = json.Unmarshal(body, &response)

	latest30 := response[0:30]
	fmt.Println(len(latest30))
	return nil, nil
}
