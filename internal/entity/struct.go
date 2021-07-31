package entity

type Config struct {
	Stocks    []*Stock
	Processes []*Process
	Optimize  []string
}

type Stock struct {
	Key   string
	Value int
}

type Process struct {
	Name    string
	Needs   []*Stock
	Results []*Stock
	Delay   int
}
