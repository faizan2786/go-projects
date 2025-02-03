package services

import "fmt"

type Arith struct{}

type DivisionResult struct {
	Quo, Rem int // quotient and remainder
}

func (*Arith) Multiply(args *Args, result *int) error {
	*result = args.A * args.B
	return nil
}

func (*Arith) Divide(args *Args, result *DivisionResult) error {
	if args.B == 0 {
		return fmt.Errorf("divide by zero")
	}
	*result = DivisionResult{Quo: args.A / args.B, Rem: args.A % args.B}
	return nil
}
