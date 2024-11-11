package main

import (
	"fmt"
	"math/rand"
	"time"
)

type account struct {
	login    string
	password string
	url      string
}

func (acc *account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) string {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	acc.password = string(res)
}

func newAccount(login, password, url string) *account {
	return &account{login, password, url}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxezABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-!@#$%^&*()_+-=")

func main() {
	rand.Seed(time.Now().UnixNano())
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount := newAccount(login, password, url)
	myAccount.generatePassword(12)
	myAccount.outputPassword()
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var input string
	fmt.Scan(&input)
	return input
}
