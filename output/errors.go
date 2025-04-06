package output

import "fmt"

func PrintError(value any) {
	intValue, ok := value.(int)
	if ok {
		fmt.Println(intValue)
		return
	}
	strValue, ok := value.(string)
	if ok {
		fmt.Println(strValue)
		return
	}
	errorValue, ok := value.(error)
	if ok {
		fmt.Println(errorValue)
		return
	}
}
