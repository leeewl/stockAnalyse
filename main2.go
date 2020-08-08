package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	//counts := make(map[string]int)
	var filename string
	filename = os.Args[1]
	m, _ := strconv.ParseFloat(os.Args[2], 32)
	initP, _ := strconv.ParseFloat(os.Args[3], 32)
	fmt.Printf("%s\n", filename)
	fmt.Printf("m %f\n", m)
	fmt.Printf("initP %f\n", initP)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}

	var s int64
	s = int64(m / (initP * 100))
	fmt.Printf("s %d\n", s)
	m -= float64(s*100) * initP
	fmt.Printf("m %f\n", m)

	var total float64

	for _, line := range strings.Split(string(data), "\n") {
		rowData := strings.Split(string(line), " ")
		dividend, _ := strconv.ParseFloat(rowData[0], 32)
		price, _ := strconv.ParseFloat(rowData[1], 32)
		m += dividend * float64(s)
		newS := int64(m / (price * 100))
		s += newS
		m -= float64(newS*100) * price
		fmt.Printf("row 0 %f\n", dividend)
		fmt.Printf("row 1 %f\n", price)
		total = price*float64(s*100) + m
	}
	fmt.Printf("total %f\n", total)

}
