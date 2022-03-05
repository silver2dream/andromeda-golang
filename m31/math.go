package m31

import (
	"fmt"
	"math"
)

const BIT_MOVE_TIMES = 10
const MULTIPLIER_FACTOR = 1 << BIT_MOVE_TIMES

type IFInt interface {
	GetValue() int
	Add(value IFInt) IFInt
	Sub(value IFInt) IFInt
	Mul(value IFInt) IFInt
	Div(value IFInt) IFInt
	Neg() IFInt
	Equal(value IFInt) bool
	NotEqual(value IFInt) bool
	GreaterThan(value IFInt) bool
	GreaterThanOrEqual(value IFInt) bool
	LessThan(value IFInt) bool
	LessThanOrEqual(value IFInt) bool
	LeftShift(moveTimes int) IFInt
	RightShift(moveTimes int) IFInt
	ToString() string
}

type FInt struct {
	scaledValue int
}

func (p *FInt) GetValue() int {
	return p.scaledValue
}

func (p *FInt) Add(value IFInt) IFInt {
	p.scaledValue = p.scaledValue + value.GetValue()
	return p
}

func (p *FInt) Sub(value IFInt) IFInt {
	p.scaledValue = p.scaledValue - value.GetValue()
	return p
}

func (p *FInt) Mul(value IFInt) IFInt {
	newValue := (p.scaledValue * value.GetValue()) >> BIT_MOVE_TIMES
	p.scaledValue = newValue
	return p
}

func (p *FInt) Div(value IFInt) IFInt {
	if value.GetValue() == 0 {
		panic("divisor by zero")
	}

	newValue := p.scaledValue << BIT_MOVE_TIMES / value.GetValue()
	p.scaledValue = newValue
	return p
}

func (p *FInt) Neg() IFInt {
	p.scaledValue = -p.scaledValue
	return p
}

func (p *FInt) Equal(value IFInt) bool {
	if value == nil {
		return false
	}

	return p.scaledValue == value.GetValue()
}

func (p *FInt) NotEqual(value IFInt) bool {
	if value == nil {
		return false
	}

	return p.scaledValue != value.GetValue()
}

func (p *FInt) GreaterThan(value IFInt) bool {
	if value == nil {
		return false
	}

	return p.scaledValue > value.GetValue()
}

func (p *FInt) GreaterThanOrEqual(value IFInt) bool {
	if value == nil {
		return false
	}

	return p.scaledValue >= value.GetValue()
}

func (p *FInt) LessThan(value IFInt) bool {
	if value == nil {
		return false
	}

	return p.scaledValue < value.GetValue()
}

func (p *FInt) LessThanOrEqual(value IFInt) bool {
	if value == nil {
		return false
	}

	return p.scaledValue <= value.GetValue()
}

func (p *FInt) LeftShift(moveTimes int) IFInt {
	p.scaledValue = p.scaledValue << moveTimes
	return p
}

func (p *FInt) RightShift(moveTimes int) IFInt {
	p.scaledValue = p.scaledValue >> moveTimes
	return p
}

func (p *FInt) RawFloat() float64 {
	return float64(p.scaledValue) / float64(MULTIPLIER_FACTOR)
}

func (p *FInt) RawInt() int {
	return p.scaledValue >> BIT_MOVE_TIMES
}

func (p *FInt) ToString() string {
	return fmt.Sprintf("%f", p.RawFloat())
}

func NewFInt(value any) IFInt {
	var iFInt IFInt
	switch v := value.(type) {
	case int32:
	case int64:
	case int:
		iFInt = &FInt{
			scaledValue: v * MULTIPLIER_FACTOR,
		}
	case float32:
	case float64:
		newValue := int(math.Round(v * MULTIPLIER_FACTOR))
		iFInt = &FInt{
			scaledValue: newValue,
		}
	}
	return iFInt
}
