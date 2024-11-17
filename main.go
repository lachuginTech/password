package main

import (
	"fmt"
	"math/rand"
	"password/account"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат url")
		return
	}
	myAccount.OutputPassword()
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var input string
	fmt.Scanln(&input)
	return input
}
