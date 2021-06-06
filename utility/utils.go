package utility

import "fmt"

func PrintKeyValue(key int64, value string) {
	fmt.Println("<" + fmt.Sprintf("%d", key) + " " + value + ">")
}
