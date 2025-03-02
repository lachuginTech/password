package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"net/url"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-!@#$%^&*()_+-=")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc *Account) OutputPassword() {
	color.Cyan(acc.Login)
	fmt.Println(acc.Password)
	fmt.Println(acc.Url)
}

func (acc *Account) ToBytes() ([]byte, error) {
	file, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	acc.Password = string(res)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("invalid login")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("invalid url")
	}
	newAcc := &Account{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Login:     login,
		Password:  password,
		Url:       urlString,
	}

	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}
