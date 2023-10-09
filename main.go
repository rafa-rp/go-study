package main

import "fmt"

type currentAccount struct {
	owner         string
	agency        int
	accountNumber int
	balance       float64
}

func main() {

	firstAccoutn := currentAccount{owner: "Rafael", balance: 100.5}
	secondAccount := currentAccount{"Mariana", 111, 123456, 200.0}

	fmt.Println(firstAccoutn)
	fmt.Println(secondAccount)
}
