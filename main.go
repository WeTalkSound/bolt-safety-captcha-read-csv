package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// Highscore struct
type Highscore struct {
	Name string
	Handle string
	SocialMedia string
	Country string
	Score int
	Category string
	Time string
} 

func main()  {
	jsonDataFromFile, err := ioutil.ReadFile("./scores.json")

	if err != nil {
		fmt.Println(err)
	}

	var jsonData []Highscore
	err = json.Unmarshal([]byte(jsonDataFromFile), &jsonData)

	if err != nil {
		fmt.Println(err)
	}

	csvFile, err := os.Create("./scores.csv")

	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)

	var row []string
	row = append(row, "Name")
	row = append(row, "Handle")
	row = append(row, "Social Media")
	row = append(row, "Country")
	row = append(row, "Score")
	row = append(row, "Category")
	row = append(row, "Time")
	writer.Write(row)

	for _, usance := range jsonData {
		var row []string
		row = append(row, usance.Name)
		row = append(row, usance.Handle)
		row = append(row, usance.SocialMedia)
		row = append(row, usance.Country)
		row = append(row, strconv.Itoa(usance.Score))
		row = append(row, usance.Category)
		row = append(row, usance.Time)
		writer.Write(row)
	}

	writer.Flush()
}