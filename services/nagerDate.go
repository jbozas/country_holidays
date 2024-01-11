package services

import (
	"fmt"
	"net/http"
)

type HolidayService struct {
	BaseURL string
}


func (s *HolidayService) GetHolidays(year, countryCode string)(*http.Response, error) {
	apiURL := fmt.Sprintf("%s/publicholidays/%s/%s", s.BaseURL, year, countryCode)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return resp, nil
}
