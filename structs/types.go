package types

type Config struct {
	MinDelay         int      `json:"minDelay"`
	MaxDelay         int      `json:"maxDelay"`
	Host             string   `json:"host"`
	DefaultPort      string   `json:"defaultPort"`
	NumOfFaultyNodes int      `json:"numOfFaultyNodes"`
	FaultyNodes      []string `json:"faultyNodes"`
	Nodes            []Node   `json:"nodes"`
}

type Node struct {
	Number string  `json:"number"`
	Port   string  `json:"port"`
	Value  float64 `json:"value"`
}

type Message struct {
	Round  string
	Value  float64
	IsExit bool
}
