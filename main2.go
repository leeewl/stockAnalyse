package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 普通理财利率3%
	const normalRate float64 = 0.03
	var filename string
	filename = os.Args[1]
	// 初始现金
	initMoney, _ := strconv.ParseFloat(os.Args[2], 32)
	// 剩余现金
	money := initMoney
	// 股价
	price, _ := strconv.ParseFloat(os.Args[3], 32)
	// 分红后是否重新买入股票 0 不买  1 买
	rebuy, _ := strconv.ParseInt(os.Args[4], 10, 32)
	// 每万元的打新收益
	pio, _ := strconv.ParseFloat(os.Args[5], 32)

	fmt.Printf("pio %f\n", pio)
	fmt.Printf("%s\n", filename)
	fmt.Printf("money %f\n", money)
	fmt.Printf("init price %f\n", price)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}

	// 股票数量
	var stock int64
	// 买成股票
	stock = int64(money / (price * 100))
	fmt.Printf("stock %d\n", stock)
	// 剩余现金
	money -= float64(stock*100) * price
	fmt.Printf("money %f\n", money)

	// 总资产
	var total float64

	fmt.Print("日期\t\t总资产\t\t复合年化收益率\n")
	for index, line := range strings.Split(string(data), "\n") {
		rowData := strings.Split(string(line), " ")
		date := rowData[0]
		// 每股分红
		dividend, _ := strconv.ParseFloat(rowData[1], 32)
		// 分红后的开盘价
		price, _ := strconv.ParseFloat(rowData[2], 32)
		//fmt.Printf("price %f\n", price)
		// 现金普通理财一年
		money = money * (1 + normalRate)
		// 分红后的现金
		money += dividend * float64(stock*100)
		// 加上打新的盈利
		money += pio * (price * float64(stock*100) / 10000)
		//fmt.Printf("pio %f\n", pio*(price*float64(stock*100)/10000))
		if rebuy > 0 {
			// 增加的股票(单位：手)
			stockAdd := int64(money / (price * 100))
			stock += stockAdd
			// 剩余现金
			money -= float64(stockAdd*100) * price
			//fmt.Printf("rebuy stock %d\n", stock)
			//fmt.Printf("rebuy money %f\n", money)
		}
		total = price*float64(stock*100) + money
		// 计算年化收益率
		annual := math.Pow((total/initMoney), 1/float64(index+1)) - 1
		fmt.Printf("%s\t%.2f\t%.2f%%\n", date, total, annual*float64(100))
		/*
			if index > 0 {
				//fmt.Printf("stock %d\n", stock)
				//fmt.Printf("money %f\n", money)
				fmt.Printf("%s\t%.2f\t%.2f%%\n", date, total, annual*float64(100))
			}
		*/
	}
	fmt.Printf("total %f\n", total)
}
