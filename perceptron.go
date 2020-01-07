package golangML

//single-layer perceptron
//bias = -0.3
func Perceptron(x ...[]float64) []float64 {
	len_x := len(x)
	num_x := len(x[0])
	learning_rate := 0.05
	bias := -0.3
	var w []float64
	for i := 1; i < len_x; i++ {
		w = append(w, 1/float64(len_x-1))
	}
	for {
		before_w := make([]float64, len(w), cap(w))
		copy(before_w, w)
		for i := 0; i < num_x; i++ {
			sum, y := 0., 0.
			for index, xi := range x {
				if index == len_x-1 {
					sum += bias
					if sum <= 0 {
						y = 0.
					} else {
						y = 1.
					}
					if xi[i] != y {
						for idx, xj := range x {
							if idx == len_x-1 {
								continue
							}
							w[idx] = w[idx] + learning_rate*xj[i]*(xi[i]-y)
						}
					}
				} else {
					sum += xi[i] * w[index]
				}
			}
		}
		for i := 0; i < len_x-1; i++ {
			if before_w[i] != w[i] {
				break
			}
			if i == len_x-2 {
				return w
			}
		}
	}
}
