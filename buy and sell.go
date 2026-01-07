package main

import "fmt"

type Trade struct {
	act   string
	name  string
	price int
}

func teller(trades <-chan Trade, results chan<- map[string]int) {
	buy := 0
	for tr := range trades {
		if tr.act == "buy" {
			buy += tr.price
		}
	}
	results <- map[string]int{
		"buy": buy,
	}
}

func main() {
	var n int
	fmt.Print("请输入购买的商品数量")
	fmt.Scan(&n)
	tradelist := make([]Trade, n)
	for i := 0; i < n; i++ {
		fmt.Printf("第%d笔交易", i+1)
		var action, name string
		var price int
		fmt.Scan(&action, &name, &price)
		if action != "buy" {
			fmt.Println("操作无效")
			continue
		}
		if price < 0 {
			fmt.Println("价格不合理！")
			continue
		}
		tradelist = append(tradelist, Trade{action, name, price})
	}
	trades := make(chan Trade, len(tradelist))
	result := make(chan map[string]int, 3)
	for i := 0; i < 3; i++ {
		go teller(trades, result)
	}
	for _, t := range tradelist {
		trades <- t
	}
	close(trades)
	total := 0
	for i := 0; i < 3; i++ {
		res := <-result
		total += res["buy"]
	}
	fmt.Println("交易完成！")
	fmt.Printf("总金额：%d\n", total)
}
