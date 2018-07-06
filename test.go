package main

import (
	"fmt"
)

func main() {
	jj := "123"
	hi := fmt.Sprint(
`{
	userID: `, jj,`
	age: `, "21", `
}`)
	fmt.Println(hi)
}