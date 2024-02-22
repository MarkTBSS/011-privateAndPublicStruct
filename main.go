package main

import (
	"fmt"

	"github.com/MarkTBSS/011-privateAndPublicStruct/models"
)

type privateStructNameBook struct {
	id     string
	name   string
	author string
}

func main() {
	// Method 1: Create and initialize using a struct literal
	bookWay1 := privateStructNameBook{
		id:     "1",
		name:   "bookWay1",
		author: "bookAuthor1",
	}
	// Method 2: Create and initialize using positional values
	bookWay2 := privateStructNameBook{"2", "bookWay2", "bookAuthor2"}
	// Method 3: Create an empty instance and then set values
	var bookWay3 privateStructNameBook
	bookWay3.id = "3"
	bookWay3.name = "bookWay3"
	bookWay3.author = "bookAuthor3"
	// Print the books
	fmt.Println("Book 1:", bookWay1)
	fmt.Println("Book 2:", bookWay2)
	fmt.Println("Book 3:", bookWay3)

	bookWay4 := models.PublicStructNameBook{
		Id:     "4",
		Name:   "bookWay4",
		Author: "bookAuthor4",
	}
	bookWay5 := models.PublicStructNameBook{Id: "5", Name: "bookWay5", Author: "bookAuthor5"}
	var bookWay6 models.PublicStructNameBook
	bookWay6.Id = "6"
	bookWay6.Name = "bookWay6"
	bookWay6.Author = "bookAuthor6"

	fmt.Println("Book 4:", bookWay4)
	fmt.Println("Book 5:", bookWay5)
	fmt.Println("Book 6:", bookWay6)
}
