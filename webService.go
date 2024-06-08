// package service

// import (
// 	"errors"
// )

// type Calculator struct{}

// func (c *Calculator) Sum(numList ...int32) int32 {
// 	var total int32 = 0
// 	for _, n := range numList {
// 		total += n
// 	}
// 	return total
// }

// func (c *Calculator) Sub(numList ...int32) int32 {
// 	var total int32 = 0
// 	for _, n := range numList {
// 		total -= n
// 	}
// 	return total
// }

// func (c *Calculator) Mul(numA int32, numB int32) int64 {
// 	var product int64 = int64(numA) * int64(numB)
// 	return product
// }

// func (c *Calculator) Div(dividend int32, divisor int32) (float32, error) {
// 	if divisor == 0 {
// 		var err error = errors.New("Division by Zero")
// 		return float32(divisor), err
// 	}
// 	var quocient float32 = float32(dividend) / float32(divisor)
// 	return float32(quocient), nil
// }
