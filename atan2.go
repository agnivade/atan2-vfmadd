package main

import "math"

func myatan2(y, x float64) float64 {
	// special cases
	switch {
	case math.IsNaN(y) || math.IsNaN(x):
		return math.NaN()
	case y == 0:
		if x >= 0 && !math.Signbit(x) {
			return math.Copysign(0, y)
		}
		return math.Copysign(math.Pi, y)
	case x == 0:
		return math.Copysign(math.Pi/2, y)
	case math.IsInf(x, 0):
		if math.IsInf(x, 1) {
			switch {
			case math.IsInf(y, 0):
				return math.Copysign(math.Pi/4, y)
			default:
				return math.Copysign(0, y)
			}
		}
		switch {
		case math.IsInf(y, 0):
			return math.Copysign(3*math.Pi/4, y)
		default:
			return math.Copysign(math.Pi, y)
		}
	case math.IsInf(y, 0):
		return math.Copysign(math.Pi/2, y)
	}

	// Call atan and determine the quadrant.
	q := atan(y / x)
	if x < 0 {
		if q <= 0 {
			return q + math.Pi
		}
		return q - math.Pi
	}
	return q
}

func atan(x float64) float64 {
	if x == 0 {
		return x
	}
	if x > 0 {
		return satan(x)
	}
	return -satan(-x)
}

// satan reduces its argument (known to be positive)
// to the range [0, 0.66] and calls xatan.
func satan(x float64) float64 {
	const (
		Morebits = 6.123233995736765886130e-17 // pi/2 = PIO2 + Morebits
		Tan3pio8 = 2.41421356237309504880      // tan(3*pi/8)
	)
	if x <= 0.66 {
		return xatan(x)
	}
	if x > Tan3pio8 {
		return math.Pi/2 - myxatan(1/x) + Morebits
	}
	return math.Pi/4 + myxatan((x-1)/(x+1)) + 0.5*Morebits
}
