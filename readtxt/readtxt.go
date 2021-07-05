package readtxt

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func ReadFile(file string) []int {
	data, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}
	content := string(data)

	arr := strings.Split(content, ",")

	newArr := []int{}

	for _, a := range arr {
		num, _ := strconv.Atoi(a)
		newArr = append(newArr, num)
	}

	return newArr
}

func FindMinMaxAverage(arr []int) {
	var sum int
	min := arr[0]
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}
	fmt.Printf("Average: %f\n", float64(sum)/float64(len(arr)))
	fmt.Printf("Min number: %d\n", min)
	fmt.Printf("Max number: %d\n", max)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; float64(i) <= math.Sqrt(float64(n)); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func CheckPrime(arr []int) {
	for i := 0; i < len(arr); i++ {
		if isPrime(arr[i]) {
			fmt.Printf("%d is a prime number \n", arr[i])
		} else {
			fmt.Printf("%d is a not prime number \n", arr[i])
		}
	}
}

func IsExist(num int, arr []int) bool {
	var count int
	for _, a := range arr {
		if num == a {
			count++
		}
	}
	if count == 0 {
		return false
	} else {
		return true
	}
}
