package main

import (
	"log"
	stack "rpc-test/example_shop/kitex_gen/example/shop/stack/stockservice"
)

func main() {
	svr := stack.NewServer(new(StockServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
