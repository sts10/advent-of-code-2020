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
	solve_part_one(text)
	solve_part_two(text)
}

func solve_part_one(text []string) {
	valid_counter := 0
	for _, line := range text {
		if is_valid_for_part_one(line) {
			valid_counter = valid_counter + 1
		}
	}
	fmt.Println("I found this many valid passwords for part one:", valid_counter)
}

func is_valid_for_part_one(line string) bool {
	criteria := strings.Split(line, ":")[0]
	password := strings.Split(line, ": ")[1]
	specified_letter := strings.Split(criteria, " ")[1]
	count_range := strings.Split(criteria, " ")[0]
	lower_bound_as_str := strings.Split(count_range, "-")[0]
	upper_bound_as_str := strings.Split(count_range, "-")[1]

	// do some type transformations
	lower_bound, _ := strconv.Atoi(lower_bound_as_str)
	upper_bound, _ := strconv.Atoi(upper_bound_as_str)
	specified_letter_as_rune := []rune(specified_letter)[0]

	// count the number of times specified_letter_as_rune shows up
	tally := count_letter_appearances(password, specified_letter_as_rune)
	return lower_bound <= tally && tally <= upper_bound
}

func count_letter_appearances(str string, letter_to_count rune) int {
	appearances := 0

	for _, char := range str {
		if char == letter_to_count {
			appearances = appearances + 1
		}
	}
	return appearances
}

func solve_part_two(text []string) {
	valid_counter := 0
	for _, line := range text {
		if is_valid_for_part_two(line) {
			valid_counter = valid_counter + 1
		}
	}
	fmt.Println("I found this many valid passwords for part two:", valid_counter)
}
func is_valid_for_part_two(line string) bool {
	criteria := strings.Split(line, ":")[0]
	password := strings.Split(line, ": ")[1]
	specified_letter := strings.Split(criteria, " ")[1]
	count_range := strings.Split(criteria, " ")[0]
	slot_one_as_str := strings.Split(count_range, "-")[0]
	slot_two_as_str := strings.Split(count_range, "-")[1]

	slot_one, _ := strconv.Atoi(slot_one_as_str)
	slot_two, _ := strconv.Atoi(slot_two_as_str)

	specified_letter_as_rune := []rune(specified_letter)[0]
	rune_in_slot_one := []rune(password)[slot_one-1]
	rune_in_slot_two := []rune(password)[slot_two-1]

	if rune_in_slot_one == specified_letter_as_rune && rune_in_slot_two == specified_letter_as_rune {
		// if specificed letter is in BOTH slots, it's an invalid password
		return false
	} else if rune_in_slot_one == specified_letter_as_rune || rune_in_slot_two == specified_letter_as_rune {
		return true
	} else {
		return false
	}
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

	return text
}
