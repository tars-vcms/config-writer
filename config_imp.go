package main

import (
	"context"
)

// configImp servant implementation
type configImp struct {
}

// Init servant init
func (imp *configImp) Init() error {
	//initialize servant here:
	//...
	return nil
}

// Destroy servant destory
func (imp *configImp) Destroy() {
	//destroy servant here:
	//...
}

func (imp *configImp) Add(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your function
	//...
	*c = a + b
	return 0, nil
}
func (imp *configImp) Sub(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//Doing something in your functionpackage main
	//
	//import (
	//	"context"
	//)
	//
	//// configImp servant implementation
	//type configImp struct {
	//}
	//
	//// Init servant init
	//func (imp *configImp) Init() error {
	//	//initialize servant here:
	//	//...
	//	return nil
	//}
	//
	//// Destroy servant destory
	//func (imp *configImp) Destroy() {
	//	//destroy servant here:
	//	//...
	//}
	//
	//func (imp *configImp) Add(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//	//Doing something in your function
	//	//...
	//	*c = a+b;
	//	return 0, nil
	//}
	//func (imp *configImp) Sub(ctx context.Context, a int32, b int32, c *int32) (int32, error) {
	//	//Doing something in your function
	//	//...
	//	return 0, nil
	//}
	//...
	return 0, nil
}
