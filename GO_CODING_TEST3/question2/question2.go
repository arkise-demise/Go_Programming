package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readAndParseFile(myFile string) (map[string]string, error) {
	file, err := os.Open(myFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := make(map[string]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2)

		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid format in line: %s", line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		data[key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
}

func main() {
	myFile := "data.txt"

	data, err := readAndParseFile(myFile)

	if err != nil {
		panic(err)
	}
	fmt.Println("=========================")
	fmt.Println("key-value paris are :")
	fmt.Println("=========================")

	for key, value := range data {
		fmt.Printf("%s:%s\n", key, value)
	}
	fmt.Println("=========================")

}