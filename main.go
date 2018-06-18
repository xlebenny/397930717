package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type myTable struct {
	gorm.Model
	MyColumn1 string
	MyColumn2 Annotation
}

func main() {
	connectionString := fmt.Sprintf("myconnectionString")
	database, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	data := myTable{
		MyColumn1: "MyColumn1Value1",
		MyColumn2: Annotation{Values: map[string]string{
			"key1": "value1",
			"key2": "value2",
		}},
	}
	database.AutoMigrate(&myTable{})
	database.Create(&data)
}
