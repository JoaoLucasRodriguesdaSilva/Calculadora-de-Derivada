package derivada

import "fmt"

var h float64 = 0.00001

func f(x float64) float64 {
  return x*x*x + x*x + x + 1 //3x*x + 2x + 1
}

func CalcDer(x float64) {
  ddx := (f(x + h) - f(x - h)) / (2 * h)
  fmt.Printf("f(x) = %.7f, f'(x) = %.7f" , f(x), ddx)
}