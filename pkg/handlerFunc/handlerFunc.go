package handlerFunc

import (
	"encoding/json"
	"strconv"
	"time"

	"../entryModel"
)


//maps to keep track of customer weekly and daily transaction amount
var (
	weekEntry = make(map[string]map[string]entryModel.WeeklyEntry)
	dailyEntry  = make(map[string]map[string]entryModel.DailyEntry)
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}


/*
Helper function that validates daily trasaction
uses map to keep track of daily amount for each customer
returns false if exceeds 5000, or more than 3 transaction count
*/
func checkDailyValid(entryData entryModel.Entry, entryDate string, amount float64) bool {
	if _,ok := dailyEntry[entryData.CustomerID]; ok {
			if currData, ok := dailyEntry[entryData.CustomerID][entryDate]; ok {
				if (currData.EntryCount + 1) > 3 || (currData.EntryTotal + amount) > 5000 {
					return false;
				}else{
					currData.EntryCount += 1
					currData.EntryTotal += amount
					dailyEntry[entryData.CustomerID][entryDate] = currData			
				}	
			}else{
					currData.EntryDate = entryDate
					currData.EntryCount =1
					currData.EntryTotal = amount
					dailyEntry[entryData.CustomerID][entryDate] = currData
			}
	}else{
		dailyEntry[entryData.CustomerID] = make(map[string]entryModel.DailyEntry)
		currData := entryModel.DailyEntry{}
		currData.ID = entryData.ID
		currData.EntryDate = entryDate
		currData.EntryCount = 1
		currData.EntryTotal = amount
		dailyEntry[entryData.CustomerID][entryDate] = currData
	}
	return true;
}


/*
Helper function to validate weekly amount
uses map to keep track of amount for each customer per week
Declines if exceeds 20000
*/
func checkWeeklyValid(entryData entryModel.Entry, yearWeek string, amount float64) bool {
	if _,ok := weekEntry[entryData.CustomerID]; ok {
			if currData,ok := weekEntry[entryData.CustomerID][yearWeek]; ok {
				if currData.WeekAmount + amount > 20000 {
					return false;
				}else {				
					currData.WeekAmount += amount
				}		
			}else{
				currData.WeekAmount = amount
				weekEntry[entryData.CustomerID][yearWeek] = currData
			}
	}else{
		weekEntry[entryData.CustomerID] = make(map[string]entryModel.WeeklyEntry)
		currData := entryModel.WeeklyEntry{}
		currData.WeekAmount = amount
		weekEntry[entryData.CustomerID][yearWeek] = currData
	}
		return true;
}


//formats output message for output file
func createMsg(id string, customerID string, accepted bool) string {
	_response := entryModel.OutputMsg{
		ID:         id,
		CustomerID: customerID,
		Accepted:   accepted}
	response, _ := json.Marshal(_response)
	return string(response)
}


/*
main function used by service to validate transactions
gets model data and retrives information needed to be sent to helper functions
*/
func CheckIfValidTrans(entryData entryModel.Entry) string {

	amount, err := strconv.ParseFloat(entryData.LoadAmount[1:], 32 )
	Check(err)
	if amount<=0 {
		return "Not valid amount"
	}
	if amount > 5000 {
		return createMsg(entryData.ID, entryData.CustomerID, false)
	}else {
		dt,err := time.Parse("2006-01-02T15:04:05Z07:00", entryData.Time)
		Check(err)            
		year,week := dt.ISOWeek()
		yearWeekRune:= rune(year) + rune(week)
		yearWeek := string(yearWeekRune)
		entryDate := entryData.Time[:10]
		
		daily := checkDailyValid(entryData, entryDate, amount);
		weekly := checkWeeklyValid(entryData, yearWeek, amount);
		
		if daily && weekly {
			return createMsg(entryData.ID, entryData.CustomerID, true)
		} else {
			return createMsg(entryData.ID, entryData.CustomerID, false)
		}
	}

}

