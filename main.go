package main

import (
	"fmt"

	"github.com/whitewolf185/check_parser/pkg/check"

	log "github.com/sirupsen/logrus"
)

func main() {
	parser := check.MakeParser()

	fmt.Printf("Введите название файла\n>")
	var fileName string
	_, err := fmt.Scanf("%s", &fileName)
	if err != nil {
		panic(err)
	}

	if err := parser.GetScv(fileName); err != nil {
		log.Fatal(err)
	}
}
