package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"parkingLot/parkinglot"
	"strings"
)

func init() {
	filename := os.Args
	if len(filename) < 2{
		return
	}
	file, err := ioutil.ReadFile(filename[1])

	if err == nil{
		for _, line := range strings.Split(string(file), "\n") {
			msg, err := parkinglot.Process(line)

			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(msg)
			}

		}
		os.Exit(1)

	}
	fmt.Printf("Didn't find file %v\n", filename[1])
}


func main()  {
	reader := bufio.NewReader(os.Stdin)


	fmt.Println("Welcome to my parking Lot.\n Type `help` to get available commands")

	for {

		input, _ := reader.ReadString('\n')


		input = strings.Replace(input, "\n", "", -1)

		msg, err := parkinglot.Process(input)

		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(msg)
	}
}