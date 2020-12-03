package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	terrain := read_file_to_slice()
	// terrain := []string{
	// 	"..##.......",
	// 	"#...#...#..",
	// 	".#....#..#.",
	// 	"..#.#...#.#",
	// 	".#...##..#.",
	// 	"..#.##.....",
	// 	".#.#.#....#",
	// 	".#........#",
	// 	"#.##...#...",
	// 	"#...##....#",
	// 	".#..#...#.#",
	// }

	number_of_trees_hit := count_tree_hits_for_given_path(3, 1, terrain)
	fmt.Println("Trees hit in part one", number_of_trees_hit)

	// Part two
	trees_hit_slope_1 := count_tree_hits_for_given_path(1, 1, terrain)
	trees_hit_slope_2 := count_tree_hits_for_given_path(3, 1, terrain)
	trees_hit_slope_3 := count_tree_hits_for_given_path(5, 1, terrain)
	trees_hit_slope_4 := count_tree_hits_for_given_path(7, 1, terrain)
	trees_hit_slope_5 := count_tree_hits_for_given_path(1, 2, terrain)

	product := trees_hit_slope_1 * trees_hit_slope_2 * trees_hit_slope_3 * trees_hit_slope_4 * trees_hit_slope_5
	fmt.Println("Part two answer is", product)
}

func count_tree_hits_for_given_path(column_increment int, row_increment int, terrain []string) int {

	number_of_trees_hit := 0
	current_col := 0
	for current_row, line := range terrain {
		// is this a row we're stopping at? If row_increment divide evenly into this row number, then yes.
		if current_row%row_increment == 0 {
			if is_tree(current_col, line) {
				number_of_trees_hit = number_of_trees_hit + 1
			}
			current_col = current_col + column_increment
		}
	}
	return number_of_trees_hit
}

func is_tree(col_number int, line_of_terrain string) bool {
	row_width := len(line_of_terrain)
	// this use of modulo deals with the endless row problem
	this_space := line_of_terrain[col_number%row_width]
	return string(this_space) == "#"
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
