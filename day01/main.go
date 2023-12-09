package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetNumberMap() map[string]string {
	return map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
}

func getInput() ([]byte, error) {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	return content, err
}

func GetFirstNumber(s string) (string, int) {
	firstIndex := strings.IndexAny(s, "0123456789")
	if firstIndex != -1 {
		return string(s[firstIndex]), firstIndex
	}
	return "", -1
}

func GetLastNumber(s string) (string, int) {
	lastIndex := strings.LastIndexAny(s, "0123456789")
	if lastIndex != -1 {
		return string(s[lastIndex]), lastIndex
	}
	return "", -1
}

func GetFirstCharInteger(s string) (string, int) {
	numberMap := GetNumberMap()

	finalIndex := -1
	foundInt := ""
	for k, _ := range numberMap {
		wordIndex := strings.Index(s, k)
		if wordIndex != -1 && finalIndex == -1 {
			foundInt = numberMap[k]
			finalIndex = wordIndex
		} else if wordIndex != -1 && wordIndex < finalIndex {
			foundInt = numberMap[k]
			finalIndex = wordIndex
		}
	}
	return foundInt, finalIndex
}

func GetLastCharInteger(s string) (string, int) {
	numberMap := GetNumberMap()

	finalIndex := -1
	foundInt := ""
	for k, _ := range numberMap {
		wordIndex := strings.LastIndex(s, k)
		if finalIndex == -1 && wordIndex != -1 {
			foundInt = numberMap[k]
			finalIndex = wordIndex
		} else if wordIndex != -1 && wordIndex > finalIndex {
			foundInt = numberMap[k]
			finalIndex = wordIndex
		}
	}
	return foundInt, finalIndex
}

func main() {

	content, err := getInput()
	if err != nil {
		log.Fatal(err)
	}

	splitStrings := strings.Split(string(content), "\n")
	finalValue := 0
	for _, s := range splitStrings {

		// Get the first integer and index in the string
		firstInt, firstIntIndex := GetFirstNumber(s)

		// Get the first character integer and index in the string
		firstCharacterInt, firstCharacterIntIndex := GetFirstCharInteger(s)
		if firstCharacterIntIndex != -1 && firstCharacterIntIndex < firstIntIndex {
			firstInt = firstCharacterInt
		}

		// Get the last integer and index in the string
		lastInt, lastIntIndex := GetLastNumber(s)

		// Get the last character integer and index in the string
		lastCharacterInt, lastCharacterIntIndex := GetLastCharInteger(s)
		if lastCharacterIntIndex != -1 && lastCharacterIntIndex > lastIntIndex {
			lastInt = lastCharacterInt
		}

		if firstInt != "" && lastInt != "" {
			stringInt, _ := strconv.Atoi(firstInt + lastInt)
			finalValue += stringInt
		}
	}
	fmt.Println("finalValue:", finalValue)
}
