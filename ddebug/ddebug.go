package ddebug

import "fmt"

func PrintParam(data interface{}) {
	fmt.Print("\r\n ============ \r\n")
	fmt.Print(data)
	fmt.Print("\r\n ============  \r\n")
}
