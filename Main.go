package main

import (
	"errors"
	"fmt"
	"github.com/atotto/clipboard"
	"strings"
)

var MsgIsEmpty = errors.New("clipboard msg is empty")
var ColumnCountNotMatch = errors.New("column count not match")

func splitting(msg string) ([][]string, error) {
	var level2SplitMsg [][]string
	if len(strings.Trim(msg, " ")) == 0 {
		return level2SplitMsg, MsgIsEmpty
	}

	splitMsg := strings.Split(msg, "\n")
	for i, m := range splitMsg {
		// last row with empty data
		if len(splitMsg)-1 == i && m == ""{
			continue
		} else{
			level2SplitMsg = append(level2SplitMsg, strings.Split(strings.TrimRight(m, "\r"), "\t"))
		}
	}

	return level2SplitMsg, nil
}
func validation(splitMsg [][]string) error {
	var columnCount int = -1
	for i, cols := range splitMsg {
		if columnCount == -1 {
			columnCount = len(cols)
		} else if columnCount != len(cols) {
			fmt.Printf("Column count not match: %d", i)
			return ColumnCountNotMatch
		} else {
			continue
		}
	}
	return nil
}

func getColumnWidth(splitMsg [][]string) ([]int, error) {
	var arr []int

	err := validation(splitMsg)
	if err != nil {
		return arr, err
	}
	for _, _ = range splitMsg[0] {
		arr = append(arr, 0)
	}

	for _,row := range splitMsg {
		for j, value := range row {
			if arr[j] < len(value){
				arr[j] = len(value)
			}
		}
	}

	return arr, nil

}

func pad(msg string, targetWidth int, padWith string) string{
	msgLen := len(msg)
	if msgLen >= targetWidth {
		return msg
	} else {
		diff := targetWidth - msgLen

		var left = true
		for i:=0; i < diff; i++ {
			if left {
				msg = padWith+msg
			} else {
				msg = msg+padWith
			}
			left = !left
		}
		return msg
	}
}

func packForHeader(row []string, columnWidths []int) string{
	var row1 = "|"
	var row2 = "|"
	for i, columnValue := range row {
		row1 = row1+" "+pad(columnValue,columnWidths[i], " ")+" |"
		row2 = row2+"-"+pad("",columnWidths[i],"-")+"-|"
	}
	return row1+"\n"+row2
}

func packForRow(row []string, columnWidths []int) string {
	var row1 = "|"
	for i, columnValue := range row {
		row1 = row1+" "+pad(columnValue,columnWidths[i]," ")+" |"
	}
	return row1
}

func main() {
	msg, err := clipboard.ReadAll()
	if err != nil {
		panic(err)
	}
	splitMsg, err := splitting(msg)
	if err != nil {
		panic(err)
	}
	if err := validation(splitMsg); err != nil {
		panic(err)
	}

	columnWidth, err := getColumnWidth(splitMsg)
	if err != nil {
		panic(err)
	}

	var finalStr = ""

	for i, row := range splitMsg {
		if i==0 {
			finalStr += packForHeader(row, columnWidth) + "\n"
		} else {
			finalStr += packForRow(row, columnWidth) + "\n"
		}
	}

	err = clipboard.WriteAll(finalStr)
	if err != nil {
		panic(err)
	}
}
