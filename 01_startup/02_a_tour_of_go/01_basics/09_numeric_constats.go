package main

import "fmt"

// 数値の定数は高精度な値です
// intで足りないときなどにconst数値型を活用
const (
  Big = 1 << 100
  Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1}
func needFloat(x float64) float64 {
  return x * 0.1
}

func main() {
  fmt.Println(needInt(Small))
  fmt.Println(needFloat(Small))
  fmt.Println(needFloat(Big))
}
