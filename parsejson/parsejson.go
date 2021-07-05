package parsejson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Person struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

func ParseJSONFile() {
	// jsonFile, _ := os.Open("data.json")

	// defer jsonFile.Close()

	jsonFile, _ := ioutil.ReadFile("data.json")

	var people []Person

	err := json.Unmarshal([]byte(jsonFile), &people)

	if err != nil {
		return
	}

	fmt.Printf("People : %v\n", people)
}
