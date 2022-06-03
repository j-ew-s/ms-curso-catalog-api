package main

import (
	"fmt"

	"github.com/j-ew-s/ms-curso-catalog-api/userService"
)

func main() {

	res, _ := userService.GetByUserId("1")
	fmt.Println(res)
}
