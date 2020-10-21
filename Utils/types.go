package types

type Config struct {
	MinDelay int    `json:"minDelay"`
	MaxDelay int    `json:"maxDelay"`
	Host     string `json:"host"`
	Nodes    []Node `json:"nodes"`
}

type Node struct {
	Number float64 `json:"number"`
	Port   float64 `json:"port"`
	Value  float64 `json:"value"`
}
