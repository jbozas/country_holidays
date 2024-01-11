package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"country_holidays/model"
	"country_holidays/services"
)

func GetHolidaysHandler(year string, countryCode string) ([]model.Holiday, error) {

	holidayService := services.HolidayService{
		BaseURL: "https://date.nager.at/api/v3",
	 }
  
	 holidaysResponse, err := holidayService.GetHolidays(year, countryCode)
	 if err != nil {
		 fmt.Println("Error:", err)
		 return nil, err
	 }
	 defer holidaysResponse.Body.Close()
 
	 body, err := ioutil.ReadAll(holidaysResponse.Body)
	 if err != nil {
		 fmt.Println("Error reading response body:", err)
		 return nil, err
	 }
 
	 var holidays []model.Holiday
	 err = json.Unmarshal(body, &holidays)
	 if err != nil {
		 fmt.Println("Error decoding JSON:", err)
		 return nil, err
	 }
	
	return holidays, nil
}
