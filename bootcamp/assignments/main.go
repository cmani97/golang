package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/assignments/models"
)

var slots []models.ParkingLot

func main() {
	fileName := "parking_lot_file_inputs.txt"
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		}

	}
}
