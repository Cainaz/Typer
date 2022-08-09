package main

import (
	"fmt"

	"github.com/Cainaz/Go/utils"
)

func main() {
	r := utils.StringInList("a", []string{"a", "b"})
	fmt.Printf("String in list = %v \n", r)
}
