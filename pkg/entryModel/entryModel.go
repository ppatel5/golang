package entryModel

//defines each entry from json file
type Entry struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	LoadAmount string `json:"load_amount"`
	Time       string `json:"time"`
}


//maintains weekly amount
type WeeklyEntry struct {
	WeekAmount float64
}


//maintains daily amount
type DailyEntry struct {
	ID               string
	EntryDate	     string
	EntryCount       int
	EntryTotal       float64
}


//defines trasaction validity output message
type OutputMsg struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	Accepted   bool   `json:"accepted"`
}
