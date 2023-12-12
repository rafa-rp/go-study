package main

import (
	"bank-go/accounts"
	"fmt"
)

func main() {

	//How instance in a different way and atribute values
	firstAccoutn := accounts.CurrentAccount{owner: "Rafael", balance: 100.5}
	secondAccount := accounts.CurrentAccount{"Mariana", 111, 123456, 200.0}
	firstAccoutn.agency = 123

	fmt.Println(firstAccoutn)
	fmt.Println(secondAccount)

	//How intance in a third way, less usual in go
	var thirdAccount *accounts.CurrentAccount
	thirdAccount = new(accounts.CurrentAccount)
	thirdAccount.owner = "Guilherme"
	fmt.Println(thirdAccount)
	fmt.Println(*thirdAccount)

	//Comparative exercise
	var thirdAccount2 *accounts.CurrentAccount
	thirdAccount2 = new(accounts.currCurrentAccountentAccount)
	thirdAccount2.owner = "Guilherme"
	secondAccount2 := accounts.CurrentAccount{"Mariana", 111, 123456, 200.0}

	fmt.Println(secondAccount == secondAccount2)
	fmt.Println(thirdAccount == thirdAccount2)
	fmt.Println(*thirdAccount == *thirdAccount2)

	//Using the function withdraw
	fourthAccount := accounts.CurrentAccount{balance: 600}
	fmt.Println(fourthAccount.withdraw(100))
	fmt.Println(fourthAccount.balance)

	//Using function with more than 1 return
	fifthAccount := accounts.CurrentAccount{owner: "Giovanna", balance: 300.0}
	status, balance := fifthAccount.deposit(500)
	fmt.Println(status, balance)

	//Using transfer function
	sixthAccount := accounts.CurrentAccount{owner: "Paulo", balance: 500.0}
	seventhAccount := accounts.CurrentAccount{owner: "Marcela", balance: 600.0}

	result := sixthAccount.transfer(400, &seventhAccount)
	fmt.Println(result)
	fmt.Println(sixthAccount)
	fmt.Println(seventhAccount)
}
