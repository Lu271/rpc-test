package main

import (
	item "github.com/Lu271/rpc-test/example_shop/kitex_gen/example/shop/item/itemservice"
	"log"
)

func main() {
	svr := item.NewServer(new(ItemServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
