package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type EngineersData struct {
	Name       string `json:"name"`
	Profession string `json:"profession"`
	Age        string `json:"age"` // use int if you prefer: int `json:"age"`
}

type Payload struct {
	EngineersData EngineersData `json:"engineersData"`
}

func processFileData(filename string) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error while reading file", err)
		return
	}
	var data Payload
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("Error while unmarshalling json data", err)
		return
	}
	fmt.Printf("Data read from file %s : %+v\n", filename, data)

	writeBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error while marshalling json data", err)
		return
	}
	err = os.WriteFile("testWrite.json", writeBytes, 0777)
	if err != nil {
		fmt.Println("Error while writing to file", err)
		return
	}
}

func main() {
	processFileData("testRead.json")
}
