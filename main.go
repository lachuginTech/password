package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"math/rand"
	"password/account"
	"password/encrypter"
	"password/files"
	"password/output"
	"strings"
	"time"
)

var menu = map[string]func(db *account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

var menuVariants = []string{
	"1. Создать аккаунт",
	"2. Найти аккаунт по url",
	"3. Найти аккаунт по логину",
	"4. Удалить аккаунт",
	"5. Выход",
	"Выберите вариант:",
}

func main() {
	fmt.Println("__Менеджер паролей__")
	err := godotenv.Load()
	if err != nil {
		output.PrintError("Не удалось найти env файл")
	}
	vault := account.NewVault(files.NewJsonDb("data.vault"), *encrypter.NewEncrypter())
Menu:
	for {
		variant := promptData(menuVariants...)
		menuFunc := menu[variant]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
		//switch variant {
		//case "1":
		//	createAccount(vault)
		//case "2":
		//	findAccount(vault)
		//case "3":
		//	deleteAccount(vault)
		//default:
		//	break Menu
		//}
	}
}

func findAccountByUrl(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска")
	accounts := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
	outputResult(&accounts)
}

func findAccountByLogin(vault *account.VaultWithDb) {
	login := promptData("Введите логин для поиска")
	accounts := vault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})
	outputResult(&accounts)
}

func outputResult(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		color.Blue("Аккаунт не найден")
	}
	for _, account := range *accounts {
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска")
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		output.PrintError("Не найдено")
	}
}

func createAccount(vault *account.VaultWithDb) {
	rand.Seed(time.Now().UnixNano())
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Неверный формат url")
		return
	}
	vault.AddAccount(*myAccount)
}

func promptData(prompt ...string) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v:", line)
		} else {
			fmt.Println(line)
		}
	}
	var input string
	fmt.Scanln(&input)
	return input
}
