package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	text := read_file_to_slice()
	valid_counter := 0
	for _, line := range text {
		if is_valid(line) {
			valid_counter = valid_counter + 1
		} else {
			fmt.Println("found an invalid password")
		}
	}
	fmt.Println("I found this many valid passwords", valid_counter)
}

func is_valid(line string) bool {
	criteria := strings.Split(line, ":")[0]
	password := strings.Split(line, ": ")[1]
	specified_letter := strings.Split(criteria, " ")[1]
	count_range := strings.Split(criteria, " ")[0]
	lower_bound_as_str := strings.Split(count_range, "-")[0]
	upper_bound_as_str := strings.Split(count_range, "-")[1]

	lower_bound, _ := strconv.Atoi(lower_bound_as_str)
	upper_bound, _ := strconv.Atoi(upper_bound_as_str)

	// fmt.Println("Criteria is", criteria)
	// fmt.Println("Password is", password)
	// fmt.Println("specified_letter is", specified_letter)
	// fmt.Println("lower_bound is", lower_bound)
	// fmt.Println("upper_bound is", upper_bound)

	specified_letter_as_rune := []rune(specified_letter)[0]

	tally := count_letter_appearances(password, specified_letter_as_rune)
	if lower_bound <= tally && tally <= upper_bound {
		return true
	} else {
		return false
	}
}

func count_letter_appearances(str string, letter_to_count rune) int {
	// str := "wwwwwwbwwhww"
	appearances := 0

	for _, char := range str {
		if char == letter_to_count {
			appearances = appearances + 1
		}
	}
	return appearances
}

func read_file_to_slice() []string {
	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.
	file, err := os.Open("input")

	if err != nil {
		log.Fatalf("failed to open")

	}

	// The bufio.NewScanner() function is called in which the
	// object os.File passed as its parameter and this returns a
	// object bufio.Scanner which is further used on the
	// bufio.Scanner.Split() method.
	scanner := bufio.NewScanner(file)

	// The bufio.ScanLines is used as an
	// input to the method bufio.Scanner.Split()
	// and then the scanning forwards to each
	// new line using the bufio.Scanner.Scan()
	// method.
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	// The method os.File.Close() is called
	// on the os.File object to close the file
	file.Close()

	// and then a loop iterates through
	// and prints each of the slice values.
	return text
}

func slice_of_strings_to_slice_of_ints(slice_strings []string) []int {
	var slice_of_ints []int
	for _, str := range slice_strings {
		this_int, _ := strconv.Atoi(str)
		slice_of_ints = append(slice_of_ints, this_int)
	}
	return slice_of_ints
}
