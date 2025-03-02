package files

import (
	"fmt"
	"os"
)

func ReadFile() {
	date, err := os.ReadFile("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(date))
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
