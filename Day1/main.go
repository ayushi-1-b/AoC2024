package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var list1 []int
	var list2 []int


	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()

		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Invalid line format:", line)
			continue
		}

		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error converting first number:", err)
			continue
		}
		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Error converting second number:", err)
			continue
		}

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println(distanceTotal(list1,list2))
	fmt.Println(similarityTotal(list1,list2))
}


func sumWithForLoop(numbers []int) int {
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    return sum
}

func distanceTotal(list1 []int, list2 []int) int {

	var distances []int

	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})

	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

	for i:= range list1 {
		if list1[i] < list2[i] {
			distance := list2[i] - list1[i]
			distances = append(distances, distance)
		} else {
			distance := list1[i] - list2[i]
			distances = append(distances, distance)
		}
	}
	return(sumWithForLoop(distances))
}

func similarityTotal(list1 []int, list2 []int) int {

	total:=0
	var countMap map[int]int
	countMap = make(map[int]int)
	for _, item := range list2{
		value, err := countMap[item]
		if (err) {
			countMap[item] = value + 1
		} else {
			countMap[item] = 1
		}
	}

	for _, item1:= range list1{
		value, erro := countMap[item1]
		if (erro) {
			total = total + (item1*value)
		}
	}
	return total
}