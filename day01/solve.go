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
    part_one := solve_part_one(text)
    fmt.Println("Answer to part 1 is", part_one)
    part_two := solve_part_two(text)
    fmt.Println("Answer to part 2 is", part_two)
}

func solve_part_one(text []string) int {
    for _, n := range text {
        n_as_int, _ := strconv.Atoi(n)
        for _, m := range text {
            m_as_int, _ := strconv.Atoi(m)
            if n_as_int + m_as_int == 2020 {
                product := n_as_int * m_as_int
                // fmt.Println("Answer is ", product)
                return product
            }
        }
    }
    return 0
}

func solve_part_two(text []string) int {
    for _, a := range text {
        a_as_int, _ := strconv.Atoi(a)
        for _, b := range text {
            b_as_int, _ := strconv.Atoi(b)
            for _, c := range text {
                c_as_int, _ := strconv.Atoi(c)
                if a_as_int + b_as_int + c_as_int == 2020 {
                    product := a_as_int * b_as_int * c_as_int
                    // fmt.Println("Answer is ", product)
                    return product
                }
            }
        }
    }
    return 0
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
