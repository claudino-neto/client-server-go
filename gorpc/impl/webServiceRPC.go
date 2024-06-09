package impl

import (
	"errors"
)

type Args struct {
	A float32
	B float32
}

type Calculator struct{}

func (c *Calculator) Sum(args *Args, reply *float32) error {
	*reply = args.A + args.B
	return nil
}

func (c *Calculator) Sub(args *Args, reply *float32) error {
	*reply = args.A - args.B
	return nil
}

func (c *Calculator) Mul(args *Args, reply *float32) error {
	*reply = args.A * args.B
	return nil
}

func (c *Calculator) Div(args *Args, reply *float32) error {
	if args.B == 0 {
		var err error = errors.New("division by zero")
		return err
	}
	*reply = args.A / args.B
	return nil

}
