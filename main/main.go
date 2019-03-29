package main

import (
	"fmt"

	"github.com/bogertron/bb-consumer/file_consumer/retrosheet"
)

func main() {
	fmt.Println("Starting consumer service")
	games, err := retrosheet.ParseFile("/home/bogertron/Downloads/2018eve/2018SDN.EVN")
	if err == nil {
		fmt.Println("Completed with ", len(games), " games")
	} else {
		fmt.Println("Error ", err)
	}
}
