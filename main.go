package main

import (
	"llouis/memoryDB/DB"
)

func main() {
	runner := DB.NewRunner()
	runner.Run()
}