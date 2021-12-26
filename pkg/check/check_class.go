package check

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
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

type Parser struct {
	Separator rune
	ResFold   string
	JsonFold  string
}

func MakeParser() *Parser {
	var result Parser
	_, b, _, _ := runtime.Caller(0)
	Root := filepath.ToSlash(filepath.Join(filepath.Dir(b), "../.."))
	log.Printf("Root now is %s", Root)

	result.Separator = '\t'
	result.ResFold = Root + "/cmd/results/"
	result.JsonFold = Root + "/cmd/JSONs/"
	return &result
}

func (obj *Parser) getTimeStamp() string {
	result := time.Now().Format("02 Jan 06 15:04 MST")
	result = strings.Replace(result, " ", "_", -1)
	result = strings.Replace(result, ":", "-", -1)

	return result
}

func (obj *Parser) GetScv(fileName string) error {
	var items Items
	file, err := ioutil.ReadFile(obj.JsonFold + fileName)
	if err != nil {
		newErr := fmt.Sprintf("File read err. Error : %s", err.Error())
		return errors.New(newErr)
	}
	err = json.Unmarshal(file, &items)
	if err != nil {
		return errors.New(fmt.Sprintf("Ошибка unmarchal. Error: %s", err.Error()))
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

	csvFile, err := os.Create(obj.ResFold + obj.getTimeStamp() + ".csv")
	if err != nil {
		return errors.New(fmt.Sprintf("Failed creating file: %e", err))
	}
	defer func() {
		err := csvFile.Close()
		if err != nil {
			panic(err)
		}
	}()

	csvwriter := csv.NewWriter(csvFile)
	csvwriter.Comma = obj.Separator

	if err := csvwriter.Write([]string{"Название", "Цена за ед", "Количество", "Сумма"}); err != nil {
		return errors.New(fmt.Sprintf("Ошибка записи в csb writer"))
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
			return errors.New(fmt.Sprintf("Ошибка записи в csb writer"))
		}
	}
	csvwriter.Flush()

	return nil
}
