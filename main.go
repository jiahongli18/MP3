package main

import (
	"fmt"
	"math/rand"
	"time"
	"./Utils"
)

/*var n int
var f int
var states []float64*/

//Set the number of nodes and maximum faulty nodes
func SetNodes()(int, int){
	fmt.Println("Specify n (3<=n<=6) and f (n>2f)")
	_, err := fmt.Scanf("%d %d", &Utils.N, &Utils.F)
	if err != nil {
		return 0,0
	}
	return Utils.N, Utils.F
}

func SetStates(n int)(s []float64){
	for i:=0; i < n; i++{
		rand.Seed(time.Now().UnixNano())
		Utils.States = append(Utils.States, rand.Float64())
	}
	return Utils.States
}

func Exports()(n1 int, f1 int){
	//n,f :=SetNodes()
	//states = SetStates(n)
	//fmt.Print(n,f, states)
	//return n,f, states
	return 1,1
}
func main() {
	SetNodes()
	//states = SetStates(n)
	//fmt.Print(n,f)
	//fmt.Print(states)

}


