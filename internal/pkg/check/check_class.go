package check

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/whitewolf185/check_parser/internal/pkg/domain"
)

type Parser struct {
	Separator rune
	ResFold   string
	JsonFold  string
}

func MakeParser() *Parser {
	var result Parser
	_, b, _, _ := runtime.Caller(0)
	Root := filepath.ToSlash(filepath.Join(filepath.Dir(b), "../../.."))
	log.Printf("Root now is %s", Root)

	result.Separator = '\t'
	result.ResFold = Root + "/doc/results/"
	result.JsonFold = Root + "/doc/JSONs/"
	return &result
}

func (obj *Parser) getTimeStamp() string {
	result := time.Now().Format("02 Jan 06 15:04 MST")
	result = strings.Replace(result, " ", "_", -1)
	result = strings.Replace(result, ":", "-", -1)

	return result
}

func (obj *Parser) GetCSV(fileName string) error {
	var receipt domain.FullReceipt
	file, err := ioutil.ReadFile(obj.JsonFold + fileName)
	if err != nil {
		return fmt.Errorf("File read err: %w", err)
	}
	err = json.Unmarshal(file, &receipt)
	if err != nil {
		return fmt.Errorf("Ошибка unmarchal: %w", err)
	}

	if len(receipt) >= 2 {
		log.Warnf("receipt has len %d", len(receipt))
	}

	result := obj.countDoubles(receipt[0].Ticket.Document.Receipt.Items)

	return obj.writeToCSV(result)
}

func (obj *Parser) countDoubles(items domain.Items) map[string]domain.Item {
	result := make(map[string]domain.Item)

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

	return result
}

func (obj *Parser) writeToCSV(items map[string]domain.Item) error {
	// TODO вставлять время из чека
	csvFile, err := os.Create(obj.ResFold + obj.getTimeStamp() + ".csv")
	if err != nil {
		return fmt.Errorf("Failed creating file: %w", err)
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
		return fmt.Errorf("Ошибка записи в csv writer: %w", err)
	}

	for _, item := range items {
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
			return fmt.Errorf("Ошибка записи в csv writer при записи данных: %w", err)
		}
	}
	csvwriter.Flush()

	return nil
}
