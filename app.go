package main

import (
	"homework010204/crawler"
	"homework010204/database"
	"log"
)

func main() {

	// arr := []int{3, 4, 1, 9, 5, 7, 10}
	// newArr := sorts.QuickSort(arr)
	// log.Println(newArr)

	// for i := 0; i < 10; i++ {
	// 	log.Printf("%d", fibonacci.Fibonacci(i))
	// }

	// parsejson.ParseJSONFile()
	// api.HandleRequest()

	// file := readtxt.ReadFile("thermopylae.txt")
	// readtxt.FindMinMaxAverage(file)
	// readtxt.CheckPrime(file)
	// fmt.Println(readtxt.IsExist(3, file))

	db, err := database.DbConnection() //Khởi tạo biến conection
	if err != nil {                    //Catch error trong quá trình thực thi
		log.Printf("Error %s when getting db connection", err)
		return
	}
	defer db.Close()
	log.Printf("Successfully connected to database")

	crawler.Crawler(db)

}
