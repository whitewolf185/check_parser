package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/whitewolf185/check_parser/internal/pkg/check"
)

func main() {
	parser := check.MakeParser()

	fmt.Printf("Введите название файла\n>")
	var fileName string
	_, err := fmt.Scanf("%s", &fileName)
	if err != nil {
		panic(err)
	}

	if err := parser.GetCSV(fileName); err != nil {
		log.Fatal(err)
	}
}
