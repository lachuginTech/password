package main

import "fmt"

type account struct {
	login    string
	password string
	url      string
}

func main() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount := account{login, password, url}

	outputPassword(&myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var input string
	fmt.Scan(&input)
	return input
}

func outputPassword(acc *account) {
	fmt.Println(acc.login, acc.password, acc.url)
}
