package golangML

import (
	"errors"
	"fmt"
)

func Linear(x, y []float64) error {
	len_x := len(x)
	len_y := len(y)
	if len_x != len_y {
		return errors.New("len_x differ to len_y")
	}
	W := 0.0
	b := 0.0

	epochs := 5000
	learning_rate := 0.01
	var cost, gradient_w, gradient_b float64

	for i := 0; i < epochs; i++ {
		cost = 0
		gradient_w = 0
		gradient_b = 0
		for j := 0; j < len_x; j++ {
			hypothesis := x[j]*W + b
			cost += ((hypothesis - y[j]) * (hypothesis - y[j]))
			gradient_w += (x[j]*W - y[j] + b) * 2 * x[j]
			gradient_b += (x[j]*W - y[j] + b) * 2
		}
		cost /= float64(len_x)
		gradient_w /= float64(len_x)
		gradient_b /= float64(len_x)
		W -= learning_rate * gradient_w
		b -= learning_rate * gradient_b

		if i%100 == 0 {
			fmt.Printf("Epoch(%5d/ 5000)  cost : %10f W: %10f b: %10f\n", i, cost, W, b)
		}
	}
	return nil
}
