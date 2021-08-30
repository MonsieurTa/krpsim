package entity

type Stocks map[string]int

type Config struct {
	Stocks       Stocks
	Processes    []*Process
	Goals        []string
	OptimizeTime bool
}

type Process struct {
	Name    string
	Needs   Stocks
	Results Stocks
	Delay   int
}
