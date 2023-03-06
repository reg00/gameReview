package main

import (
	"fmt"

	"github.com/Reg00/gameReview/cmd"
)

func main() {
	//cmd.Execute()
	err := cmd.StartServer()
	if err != nil {
		fmt.Println(err)
	}
}
