package m31

import "math"

const BIT_MOVE_TIMES = 10
const MULTIPLIER_FACTOR = 1 << BIT_MOVE_TIMES

type FInt struct {
	scaledValue int
}

func (p *FInt) GetValue() int {
	return p.scaledValue
}

type IFInt interface {
	GetValue() int
}

func ConvertFloat64ToFInt(value float64) IFInt {
	newValue := int(math.Round(value))
	return &FInt{
		scaledValue: newValue * MULTIPLIER_FACTOR,
	}
}

func ConvertIntToFInt(value int) IFInt {
	return &FInt{
		scaledValue: value * MULTIPLIER_FACTOR,
	}
}
