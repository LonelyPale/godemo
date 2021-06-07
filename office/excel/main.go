package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {
	file := "/Users/wyb/project/github/godemo/office/excel/鸿学院测试问卷-1 csv-demo-由出题组填写提供.xlsx"
	f, err := excelize.OpenFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 获取工作表中指定单元格的值
	cell, err := f.GetCellValue("鸿学院测试问卷-1 csv-demo", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)

	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("鸿学院测试问卷-1 csv-demo")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
