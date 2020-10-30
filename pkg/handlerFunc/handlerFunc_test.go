package handlerFunc

import ( "testing"
		"../entryModel"
 )


// If daily trasanction amount exceeds 5000
func TestInvalidDaily(t *testing.T){

	entry := entryModel.Entry{}
	entry.ID = "1470"
	entry.CustomerID = "440"
	entry.LoadAmount = "$5774.12"
	entry.Time = "2000-01-05T15:28:58Z"
	
	result := CheckIfValidTrans(entry)
	
	expectedResult := createMsg(entry.ID, entry.CustomerID, false);
	
	if expectedResult != result {
		t.Errorf("failed expected %s got %s", expectedResult, result) 
	}
}


// If multiple daily trasanction amount exceeds 5000
func TestInvalidMultipleDaily(t *testing.T){

	entry := entryModel.Entry{}
	entry.ID = "1471"
	entry.CustomerID = "441"
	entry.LoadAmount = "$2000.00"
	entry.Time = "2000-01-05T12:28:58Z"
	CheckIfValidTrans(entry)
	
	entry.LoadAmount = "$3000.00"
	entry.Time = "2000-01-05T14:28:58Z"
	CheckIfValidTrans(entry)
	
	entry.LoadAmount = "$74.12"
	entry.Time = "2000-01-05T15:28:58Z"
	
	result := CheckIfValidTrans(entry)
	
	expectedResult := createMsg(entry.ID, entry.CustomerID, false);
	
	if expectedResult != result {
		t.Errorf("failed expected %s got %s", expectedResult, result) 
	}
}


//Check for valid transactions
func TestValidMultipleDaily(t *testing.T){

	entry := entryModel.Entry{}
	entry.ID = "1478"
	entry.CustomerID = "445"
	entry.LoadAmount = "$2000.00"
	entry.Time = "2000-01-05T12:28:58Z"
	CheckIfValidTrans(entry)
	
	entry.LoadAmount = "$3000.00"
	entry.Time = "2000-01-05T14:28:58Z"

	
	result := CheckIfValidTrans(entry)
	
	expectedResult := createMsg(entry.ID, entry.CustomerID, true);
	
	if expectedResult != result {
		t.Errorf("failed expected %s got %s", expectedResult, result) 
	}
}

//check trasanction declined when exceeds daily count
func TestInvalidDailyTransCount(t *testing.T){

	entry := entryModel.Entry{}
	entry.ID = "1477"
	entry.CustomerID = "442"
	entry.LoadAmount = "$774.12"
	entry.Time = "2000-01-05T15:28:58Z"
	CheckIfValidTrans(entry)
	
	entry.LoadAmount = "$774.12"
	entry.Time = "2000-01-05T15:30:58Z"
	CheckIfValidTrans(entry)
	
	entry.LoadAmount = "$774.12"
	entry.Time = "2000-01-05T16:28:58Z"
	CheckIfValidTrans(entry)
	
	entry.LoadAmount = "$774.12"
	entry.Time = "2000-01-05T16:30:58Z"
	result := CheckIfValidTrans(entry)
	
	expectedResult := createMsg(entry.ID, entry.CustomerID, false);
	
	if expectedResult != result {
		t.Errorf("failed expected %s got %s", expectedResult, result) 
	}
}

//check transaction declined when exceeds weekly amount
func TestInvalidWeekly(t *testing.T) {
	
	entry := entryModel.Entry{}
	entry.ID = "1473"
	entry.CustomerID = "443"
	entry.LoadAmount = "$5000"
	entry.Time = "2020-10-26T15:28:58Z"
	CheckIfValidTrans(entry)
	
	entry.LoadAmount = "$5000"
	entry.Time = "2020-10-27T15:28:58Z"
	CheckIfValidTrans(entry)
	
	entry.LoadAmount = "$5000"
	entry.Time = "2020-10-28T15:28:58Z"
	CheckIfValidTrans(entry)
	
	entry.LoadAmount = "$5000"
	entry.Time = "2020-10-29T15:28:58Z"
	CheckIfValidTrans(entry)	
	
	entry.LoadAmount = "$2000"
	entry.Time = "2020-10-29T15:28:58Z"
	result := CheckIfValidTrans(entry)
	
	expectedResult := createMsg(entry.ID, entry.CustomerID, false);
	
	
	if expectedResult != result {
		t.Errorf("failed expected %s got %s", expectedResult, result) 
	}

}

//check trasaction aceepted when meets weekly limit
func TestValidWeekly(t *testing.T) {
	
	entry := entryModel.Entry{}
	entry.ID = "1477"
	entry.CustomerID = "444"
	entry.LoadAmount = "$5000"
	entry.Time = "2020-10-26T15:28:58Z"
	CheckIfValidTrans(entry)
	
	entry.LoadAmount = "$5000"
	entry.Time = "2020-10-27T15:28:58Z"
	CheckIfValidTrans(entry)
	
	entry.LoadAmount = "$5000"
	entry.Time = "2020-10-28T15:28:58Z"
	CheckIfValidTrans(entry)
	
	entry.LoadAmount = "$5000"
	entry.Time = "2020-10-29T15:28:58Z"

	result := CheckIfValidTrans(entry)
	
	expectedResult := createMsg(entry.ID, entry.CustomerID, true);	
	
	if expectedResult != result {
		t.Errorf("failed expected %s got %s", expectedResult, result) 
	}

}


//when amount is lees than 0
func TestValidAmount(t *testing.T) {

	entry := entryModel.Entry{}
	entry.ID = "1477"
	entry.CustomerID = "444"
	entry.LoadAmount = "$-1"
	entry.Time = "2020-10-26T15:28:58Z"
	CheckIfValidTrans(entry)
	
	
	result := CheckIfValidTrans(entry)
	
	expectedResult := "Not valid amount"
	
	if expectedResult != result {
		t.Errorf("failed expected %s got %s", expectedResult, result) 
	}
}

