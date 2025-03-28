package files

import (
	"fmt"
	"os"
)

func ReadFile(name string) ([]byte, error) {
	date, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return date, nil
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	_, err = file.Write(content)
	defer file.Close()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("Запись успешна")
}
