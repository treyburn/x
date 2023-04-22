package facts

import (
	"errors"
	"math"
)

type Op int

func (o Op) Resolve() (OpFunc, error) {
	f, ok := opFuncs[o]
	if !ok {
		return nil, errors.New("invalid op code")
	}
	return f, nil
}

const (
	Add Op = iota
	Subtract
	Multiply
	Divide
)

type OpFunc func(float64, float64) float64

var opFuncs = map[Op]OpFunc{
	Add:      add,
	Subtract: sub,
	Multiply: mlt,
	Divide:   div,
}

func add(x, y float64) float64 {
	return toFixed(x+y, 2)
}

func sub(x, y float64) float64 {
	return toFixed(x-y, 2)
}

func mlt(x, y float64) float64 {
	return toFixed(x*y, 2)
}

func div(x, y float64) float64 {
	return toFixed(x/y, 2)
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
