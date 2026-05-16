package model

type StastResponse struct {
	TotalRevenue 			int64
	TotalTransactions		int32
	Stocks					int32
	BestSeller				string
	DataChart				[]DataChart
}

type DataChart struct {
	Date		string // YYYY-MM-DD
	In			int
	Out			int
}