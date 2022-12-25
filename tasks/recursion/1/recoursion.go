package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getAccumulations(scanner *bufio.Scanner, capacity int) ([]int, error) {
	scanner.Scan()
	textStr := strings.Split(scanner.Text(), " ")
	accumulations := make([]int, capacity)
	var (
		n, k int
		err  error
	)

	for i := 0; i < len(textStr); i++ {
		curr := textStr[i]
		if curr == "" { // ебанутые присылают пустую строку тоже
			fmt.Println(curr, i, len(textStr))
			continue
		}

		n, err = strconv.Atoi(textStr[i])
		if err != nil {
			return nil, fmt.Errorf("%s: %w", textStr[i], err)
		}

		accumulations[k] = n
		k++

		if k == capacity {
			break
		}
	}

	return accumulations, nil
}

func getNumber(scanner *bufio.Scanner) (int, error) {
	scanner.Scan()
	return strconv.Atoi(scanner.Text())
}

func getParams(scanner *bufio.Scanner) (accumulations []int, bikePrice int, err error) {
	var days int
	days, err = getNumber(scanner)
	if err != nil {
		panic(fmt.Errorf("1: %w", err))
	}

	accumulations, err = getAccumulations(scanner, days)
	if err != nil {
		panic(fmt.Errorf("2: %w", err))
	}

	bikePrice, err = getNumber(scanner)
	if err != nil {
		panic(fmt.Errorf("3: %w", err))
	}

	return accumulations, bikePrice, err
}

func search(accumulations []int, expPrice, left, right int, found bool) (resp int, err error) {
	if left == right {
		return -1, nil
	}

	mid := (left + right) / 2
	midV := accumulations[mid]

	if midV >= expPrice {
		resp, err := search(accumulations, expPrice, left, mid, true)
		if err != nil {
			return -1, err
		}

		if resp == -1 {
			return mid + 1, nil
		}

		return resp, nil
	} else {
		return search(accumulations, expPrice, mid+1, right, found)
	}
}

func searchDaysWhenICanBuyTwoBikes(accumulations []int, bikePrice int) (
	whenICanBuyFirstBike, whenICanBuyTwoBike int, err error) {
	whenICanBuyFirstBike, err = search(accumulations, bikePrice, 0, len(accumulations), false)
	if whenICanBuyFirstBike <= 0 || err != nil {
		return -1, -1, err
	}

	whenICanBuyTwoBike, err = search(accumulations, bikePrice*2, 0, len(accumulations), false)
	return
}

func main() {
	const maxCapacity = 16 * 1024 * 1024
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, maxCapacity), maxCapacity)

	accumulations, twoBikePrice, err := getParams(scanner)
	if err != nil {
		panic(err)
	}

	firstBike, secondBike, err := searchDaysWhenICanBuyTwoBikes(accumulations, twoBikePrice)
	if err != nil {
		panic(err)
	}

	fmt.Println(firstBike, secondBike)
}
