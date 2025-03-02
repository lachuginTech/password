package main

import (
	"fmt"
	"math/rand"
	"password/account"
	"password/files"
	"time"
)

func main() {
	fmt.Println("__Менеджер паролей__")
Menu:
	for {
		var variant = getMenu()
		switch variant {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		default:
			break Menu
		}
	}
}

func getMenu() int {
	var variant int
	fmt.Println("Выберите вариант:")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	fmt.Scan(&variant)
	return variant
}

func findAccount() {

}

func deleteAccount() {

}

func createAccount() {
	rand.Seed(time.Now().UnixNano())
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат url")
		return
	}
	file, err := myAccount.ToBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать данные")
		return
	}
	files.WriteFile(file, "data.json")
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var input string
	fmt.Scanln(&input)
	return input
}
