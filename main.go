package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("--- Weather Station ---")

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
			weather.get()
			continue
		case "clear":
			weather.clear()
			continue
		default:
			id, value := getIdAndValue(line)
			weather.update(id, value)
			continue
		}

	}

}

func getIdAndValue(line string) (int, *float64) {
	parts := strings.Split(line, ",")
	if len(parts) != 2 {
		return 0, nil
	}

	id, _ := strconv.Atoi(parts[0])

	if parts[1] == "NULL" {
		return id, nil
	}

	number, _ := strconv.ParseFloat(parts[1], 64)
	return id, &number
}