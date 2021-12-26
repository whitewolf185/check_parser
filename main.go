package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Item struct {
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Sum      float32 `json:"sum"`
	Quantity float32 `json:"quantity"`
}

type Items struct {
	Items []Item `json:"items"`
}

func main() {
	fmt.Printf("Введите название файла\n>")
	var fileName string
	_, err := fmt.Scanf("%s", &fileName)
	if err != nil {
		panic(err)
	}

	var items Items
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &items)
	if err != nil {
		log.Fatalf("Ошибка unmarchal. Error: %s", err.Error())
	}

	result := make(map[string]Item)

	for _, item := range items.Items {
		res, ok := result[item.Name]
		if !ok {
			result[item.Name] = item
		} else {
			res.Sum += item.Sum
			res.Quantity += item.Quantity
			result[item.Name] = res
		}
	}

	csvFile, err := os.Create("result.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer func() {
		err := csvFile.Close()
		if err != nil {
			panic(err)
		}
	}()

	csvwriter := csv.NewWriter(csvFile)
	csvwriter.Comma = '\t'

	if err := csvwriter.Write([]string{"Название", "Цена за ед", "Количество", "Сумма"}); err != nil {
		log.Fatalf("Ошибка записи в csb writer")
	}

	for _, item := range result {
		var stringWriter []string
		stringWriter = append(stringWriter, item.Name)

		price := fmt.Sprintf("%.2f", item.Price/100)
		price = strings.Replace(price, ".", ",", -1)
		stringWriter = append(stringWriter, price)

		quant := fmt.Sprintf("%.2f", item.Quantity)
		quant = strings.Replace(quant, ".", ",", -1)
		stringWriter = append(stringWriter, quant)

		sum := fmt.Sprintf("%.2f", item.Sum/100)
		sum = strings.Replace(sum, ".", ",", -1)
		stringWriter = append(stringWriter, sum)

		if err := csvwriter.Write(stringWriter); err != nil {
			log.Fatalf("Ошибка записи в csb writer")
		}
	}
	csvwriter.Flush()
}
