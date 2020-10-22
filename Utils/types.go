package types

type Config struct {
	MinDelay         int    `json:"minDelay"`
	MaxDelay         int    `json:"maxDelay"`
	Host             string `json:"host"`
	DefaultPort      string `json:"defaultPort"`
	NumOfFaultyNodes int    `json:"numOfFaultyNodes"`
	Nodes            []Node `json:"nodes"`
}

type Node struct {
	Number float64 `json:"number"`
	Port   float64 `json:"port"`
	Value  float64 `json:"value"`
}

type Message struct {
	Round    int
	Value    float64
	IsFaulty bool
	IsExit   bool
}
