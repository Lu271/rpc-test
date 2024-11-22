package main

import (
	"context"
	stack "rpc-test/example_shop/kitex_gen/example/shop/stack"
)

// StockServiceImpl implements the last service interface defined in the IDL.
type StockServiceImpl struct{}

// GetItemStock implements the StockServiceImpl interface.
func (s *StockServiceImpl) GetItemStock(ctx context.Context, req *stack.GetItemStockReq) (resp *stack.GetItemStockResp, err error) {
	// TODO: Your code here...
	return
}