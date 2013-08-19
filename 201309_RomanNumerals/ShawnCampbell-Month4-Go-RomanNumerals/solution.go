package main

import (
	"os"
	"bufio"
	"bytes"
	"io"
	"fmt"
	//"strings"
	"strconv"
	"time"
)
func readLines(path string) (err error) {
	var (
		file *os.File
		part []byte
		prefix bool
		)
	values := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	numerals := [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))
	count := 0
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			convert_num_to_roman(count, buffer.String(), values, numerals)
			count++
			//lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}

func convert_num_to_roman(count int, line string, values [13]int, numerals [13]string) {
	//fmt.Println(line)
	res := ""
	number, err := strconv.Atoi(line)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 13; i++ {
		for number > values[i] {
			number -= values[i]
			res += numerals[i]
		}
	}
	fmt.Print(count)
	fmt.Println(": " + res)
}

func main() {
	start := time.Now()
	err := readLines(os.Args[1])
	if err != nil {
		fmt.Println("Error: %s\n", err)
		return
	}
	end := time.Now()
	diff := end.Nanosecond()%1e9/1e6 - start.Nanosecond()%1e9/1e6
	//miliSeconds := (diff % 1e9) / 1e6
	fmt.Print("It took me: ")
	fmt.Print(diff)
	fmt.Println(" milliseconds")

}