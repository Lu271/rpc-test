package main

import (
	stack "github.com/Lu271/rpc-test/example_shop/kitex_gen/example/shop/stack/stockservice"
	"log"
)

func main() {
	svr := stack.NewServer(new(StockServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
