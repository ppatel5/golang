
//To run main function, 
//	Open terminal-> open cmd folder-> run command " go run main.go input.txt "
//To run test cases
//	Open terminal-> open pkg/handlerFunc folder-> run command " go test "


// entryModel package- this package defines the required struct types for input data

// 	Entry is struct type define to describe json input data
// 	WeeklyEntry struct type describes week amount
// 	DailyEntry struct type describes daily amount for customer
// 	OutputMsg struct type describes the format of output message


// handlerFunc package- package defines functions to process the input data

//function CheckIfValidTrans, is function called by main package to validate the transaction. Function takes input 
//	as EntryModel.Entry, seperates information like year/week of trasaction, date of trasaction and customer id. ISOWeek 
//	inbuilt function is used to extract year and week from date.Helper functions checkDailyValid , checkWeeklyValid 
//	are used to validate trasaction and returns the output message. 
//function checkDailyValid is used to validate daily trasaction amount and count. It keeps track of amount and trasaction 
//	count using map of customerId and entry date. Returns true if amount<5000 && count<=3, returns false if it exceeds
//	daily amount/count limit.
//function checkWeeklyValid is used to validate weekly limits of transaction amount. It keeps track of weekly amount for each 
//	customer using map of customerid and year-week of trasaction. Returns false if exceeds weekly limit of 20000, else returns true


//handlerFunc_test covers the test scenarios

//main package- define within cmd folder, contains main() function 

//	It takes input file name as argument. 
//	Json Unmarshal used to decode inout file to entryModel.entry for each trasaction. If same transaction id is
//	repeated for customer, all but first trasactions are ignored. To keep track of this scenario, map of load id 
//	and customer id is maintain. 
//	handlerFunc package is used to validate each transaction,and output message is written to output file created by main function.




 
	





