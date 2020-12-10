package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	file_as_string := read_file_as_string("input")
	// fmt.Println(file_as_string)
	array_of_familes := strings.Split(file_as_string, "\n\n")
	for _, family := range array_of_familes {
		fmt.Println("----------")
		fmt.Println(family)
	}
	sum := 0
	for _, family := range array_of_familes {
		sum = sum + count_unique_letters(family)
	}
	fmt.Println("Answer to part one is", sum)

	// on to part two
	sum = 0
	for _, family := range array_of_familes {
		for _, char := range array_of_familes[0] {
			if char != '\n' && is_letter_in_every_element(family, char) {
				sum++
			}
		}
	}
	fmt.Println("Answer for part two is", sum)
}

func is_letter_in_every_element(family string, letter rune) bool {
	// fmt.Println("received family is", strings.TrimSuffix(family, "\n"))
	family = strings.TrimSuffix(family, "\n")
	fam_len := len(family)

	letter_appearances := 0
	for _, person := range strings.Split(family, "\n") {
		if contains_letter([]rune(person), letter) {
			letter_appearances++
			continue
		}
	}
	fmt.Println("I thinkevery person in", family, "contains", letter)
	return letter_appearances == fam_len
}

func count_unique_letters(str string) int {
	seen_letters := []rune{}
	for _, char := range str {
		if !contains_letter(seen_letters, char) && char != '\n' {
			seen_letters = append(seen_letters, char)
		}
	}
	return len(seen_letters)
}

func contains_letter(slice []rune, target rune) bool {
	if len(slice) == 0 {
		return false
	}
	for _, char := range slice {
		if char == target {
			return true
		}
	}
	return false
}

func read_file_as_string(file_path string) string {
	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile(file_path)
	if err != nil {
		return ""
	}

	// Convert []byte to string and print to screen
	text := string(content)
	return text
}
