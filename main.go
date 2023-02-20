package main

import (
	"github.com/go-recipes/basics"
	"github.com/go-recipes/json"
	"github.com/go-recipes/structs"
	"github.com/go-recipes/text"
	"github.com/go-recipes/time"
)

func main() {
	basics.TestBasics()
	time.TestTime()
	text.TestText()
	structs.TestStructs()
	json.TestJson()
}
