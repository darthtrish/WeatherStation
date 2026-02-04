package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {

	fmt.Println(--- Weather Station ---)

	weather := &WeatherData{}
	reader := bufio.NewReader(os.Stdin)

	for {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)


		switch line {
		case "exit":
			fmt.Println("Exiting...")
			return
		case "get":
			weather.print()
			continue
		case "clear":
			weather.clear()
			continue
		default:
			id, value := getIdAndValue(input)
			weather.update(id, value)
			continue
		}
		

	}
	

}

func 