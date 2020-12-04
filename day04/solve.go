package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Passport struct {
	byr string // (Birth Year)
	iyr string // (Issue Year)
	eyr string // (Expiration Year)
	hgt string // (Height)
	hcl string // (Hair Color)
	ecl string // (Eye Color)
	pid string // (Passport ID)
	cid string // (Country ID)
}

func main() {
	raw_list_of_passports := read_file_to_slice()
	// var passport_structs []Passport
	passport_structs := []Passport{}
	var this_passport_raw string
	for _, line := range raw_list_of_passports {
		// this logic doesn't read the _Last_ passport, since there isn't an empty line at the end of the input file
		if line != "" {
			this_passport_raw = this_passport_raw + " " + line
		} else {
			// this_passport_raw is done
			// fmt.Println("a raw passport:", this_passport_raw)
			this_passport_struct := build_passport_object(this_passport_raw)
			passport_structs = append(passport_structs, this_passport_struct)

			// empty out this variable
			this_passport_raw = ""
		}
	}
	// Oh God, this is ATROCIOUS, but given that there isn't an empty line at the end of
	// the input file, the loop logic above misses the very last passport
	// sooo... we'll just do it now real quick
	this_passport_struct := build_passport_object(this_passport_raw)
	passport_structs = append(passport_structs, this_passport_struct)

	// Now count valid passports
	number_of_valid_passports := 0
	for _, passport_struct := range passport_structs {
		if passport_is_valid_for_part_one(passport_struct) {
			number_of_valid_passports++
		}
		// fmt.Println(passport_struct, "is", passport_is_valid_for_part_one(passport_struct))
	}

	fmt.Println("I found this many valid passports for part one:", number_of_valid_passports) // 260

	// Part 2
	number_of_valid_passports = 0
	for _, passport_struct := range passport_structs {
		if passport_is_valid_for_part_two(passport_struct) {
			number_of_valid_passports++
		}
	}

	fmt.Println("I found this many valid passports for part two:", number_of_valid_passports) // 155 is too high
}

func build_passport_object(raw_str string) Passport {
	raw_as_slice := strings.Split(raw_str, " ")
	var this_passport Passport

	for _, field := range raw_as_slice {
		category := strings.Split(field, ":")[0]
		val := ""
		if len(strings.Split(field, ":")) == 2 {
			val = strings.Split(field, ":")[1]
		}
		if category == "byr" {
			this_passport.byr = val
		} else if category == "iyr" {
			this_passport.iyr = val
		} else if category == "eyr" {
			this_passport.eyr = val
		} else if category == "hgt" {
			this_passport.hgt = val
		} else if category == "hcl" {
			this_passport.hcl = val
		} else if category == "ecl" {
			this_passport.ecl = val
		} else if category == "pid" {
			this_passport.pid = val
		} else if category == "cid" {
			this_passport.cid = val
		}
	}
	return this_passport
}

func passport_is_valid_for_part_one(p Passport) bool {
	// don't care what p.cid is!
	return p.byr != "" && p.iyr != "" && p.eyr != "" && p.hgt != "" && p.hcl != "" && p.ecl != "" && p.pid != ""
}

func passport_is_valid_for_part_two(p Passport) bool {
	// is_valid := true
	byr, err := strconv.Atoi(p.byr)
	if err != nil || !(1920 <= byr && byr <= 2002) {
		return false
	}
	iyr, err := strconv.Atoi(p.iyr)
	if err != nil || !(2010 <= iyr && iyr <= 2020) {
		return false
	}
	eyr, err := strconv.Atoi(p.eyr)
	if err != nil || !(2020 <= eyr && eyr <= 2030) {
		return false
	}

	if strings.Contains(p.hgt, "cm") {
		measurement_in_cm, err := strconv.Atoi(p.hgt[0 : len(p.hgt)-2])
		if err != nil || !(150 <= measurement_in_cm && measurement_in_cm <= 193) {
			return false
		}
	} else if strings.Contains(p.hgt, "in") {
		measurement_in_in, err := strconv.Atoi(p.hgt[0 : len(p.hgt)-2])
		if err != nil || !(59 <= measurement_in_in && measurement_in_in <= 76) {
			return false
		}
	} else { // height is not present or unlabeled
		return false
	}

	if !is_valid_color(p.hcl) {
		return false
	}

	// Eye color
	if !(p.ecl == "amb" || p.ecl == "blu" || p.ecl == "brn" || p.ecl == "gry" || p.ecl == "grn" || p.ecl == "hzl" || p.ecl == "oth") {
		return false
	}

	if !is_valid_pid(p.pid) {
		return false
	}

	// don't care what p.cid is!

	// made it through all the gates, so must be valid
	return true
}

func is_valid_color(c string) bool {
	if c == "" || string(c[0]) != "#" || len(c) != 7 {
		return false
	}
	for _, char := range c[1:len(c)] {
		// fmt.Printf("This char is %v\n", char)
		if !((48 <= char && char <= 57) || (97 <= char && char <= 102)) {
			return false
		}
	}
	return true
}

func is_valid_pid(c string) bool {
	if c == "" || len(c) != 9 {
		return false
	}
	for _, char := range c {
		if !(48 <= char && char <= 57) {
			return false
		}
	}
	return true
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
