package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	seats := read_file_to_slice()
	fmt.Println("Answer to part one (highest_seat_id) is", solve_part_one(seats))
	fmt.Println("Answer to part two (my seat) is", solve_part_two(seats))
}

func solve_part_two(seats []string) int {
	seat_ids := []int{}
	for _, seat := range seats {
		seat_ids = append(seat_ids, get_seat_id(seat))
	}
	sort.Ints(seat_ids)
	for i, seat_id := range seat_ids[0:len(seat_ids)] {
		if seat_ids[i+1] == seat_id+2 {
			return seat_id + 1
		}
	}
	return 0
}

func solve_part_one(seats []string) int {
	// seat := "BBFFBBFRLL"
	highest_seat_id := 0
	for _, seat := range seats {
		this_seat_id := get_seat_id(seat)
		if this_seat_id > highest_seat_id {
			highest_seat_id = this_seat_id
		}
	}
	return highest_seat_id
}

func get_seat_id(seat string) int {
	lower := 0
	upper := 127
	for _, c := range seat[0:7] {
		mid := (upper + lower) / 2
		// if B, round up
		if c == 'B' {
			mid = mid + 1
			lower = mid
		} else if c == 'F' {
			upper = mid
		}
		// fmt.Println("New range is", lower, "to", upper)
	}
	row := upper

	// now find column
	lower = 0
	upper = 7
	for _, c := range seat[7:10] {
		mid := (upper + lower) / 2
		if c == 'R' {
			mid = mid + 1
			lower = mid
		} else if c == 'L' {
			upper = mid
		} else {
			// fmt.Println("BAD")
		}
		// fmt.Println("New range is", lower, "to", upper)
	}
	column := upper
	seat_id := row*8 + column
	return seat_id
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
