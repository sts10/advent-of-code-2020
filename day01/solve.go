package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	text := read_file_to_slice()
	text_as_int := slice_of_strings_to_slice_of_ints(text)

	part_one, err := solve_part_one(text_as_int)
	// is this a good way to handle this result?
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Answer to part 1 is", part_one)
	}

	part_two, err := solve_part_two(text_as_int)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Answer to part 2 is", part_two)
	}
}

func solve_part_one(text []int) (int, error) {
	for _, a := range text {
		for _, b := range text {
			if a+b == 2020 {
				product := a * b
				return product, nil
			}
		}
	}
	// Don't love that I return a 0 here, even with the error...
	// But I don't see another choice
	return 0, fmt.Errorf("Didn't find a solution to part one")
}

func solve_part_two(text []int) (int, error) {
	for _, a := range text {
		for _, b := range text {
			for _, c := range text {
				if a+b+c == 2020 {
					product := a * b * c
					return product, nil
				}
			}
		}
	}
	return 0, fmt.Errorf("Didn't find a solution to part two")
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
