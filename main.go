package main

import (
	"encoding/xml"
	"fmt"
)

type Book struct {
	XMLName xml.Name `xml:"Book"`
	Name    string   `xml:"name"`
	Author  string   `xml:"author"`
}

func main() {

	source :=
		`<?xml version='1.0' encoding='UTF-8'?>
		<Book>
			<name>아내를 모자로 착각한 남자11</name>
			<author>올리버 색스</author>
			<isbn13>9791159920257</isbn13>
			<isbn10>1159920257</isbn10>
		</Book>`

	book := Book{}
	err := xml.Unmarshal([]byte(source), &book)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("name : " + book.Name)
	fmt.Println("author : " + book.Author)

}