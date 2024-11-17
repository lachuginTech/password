package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
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

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	acc.password = string(res)
}

func newAccount(login, password, urlString string) (*account, error) {
	if login == "" {
		return nil, errors.New("invalid login")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("invalid url")
	}
	newAcc := &account{
		login:    login,
		password: password,
		url:      urlString,
	}
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-!@#$%^&*()_+-=")

func main() {
	rand.Seed(time.Now().UnixNano())
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := newAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат url")
		return
	}
	myAccount.outputPassword()
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var input string
	fmt.Scanln(&input)
	return input
}
