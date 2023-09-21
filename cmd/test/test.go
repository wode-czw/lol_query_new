package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Address   struct {
		Street  string `json:"street"`
		City    string `json:"city"`
		ZipCode string `json:"zip_code"`
	} `json:"address"`
}

func main() {
	// 假设我们从某个地方得到了以下JSON数据：
	jsonData := `{
		"first_name": "John",
		"last_name": "Doe",
		"age": 30,
		"address": {
			"street": "123 Main St",
			"city": "Springfield",
			"zip_code": "12345"
		}
	}`

	// 解析JSON数据到我们的Person结构体
	var person Person
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印解析后的数据
	fmt.Println("Person Details:")
	fmt.Println("First Name:", person.FirstName)
	fmt.Println("Last Name:", person.LastName)
	fmt.Println("Age:", person.Age)
	fmt.Println("Street:", person.Address.Street)
	fmt.Println("City:", person.Address.City)
	fmt.Println("Zip Code:", person.Address.ZipCode)
}
