package section3

import "math"

func absValue(a *float64) interface{}{

	if a == nil {
		return nil
	}
	return math.Abs(*a)
}
