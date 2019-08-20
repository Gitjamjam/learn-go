package main

import (
  "fmt"
)

func main(){
  a := f1(1,1)
  f2("m1","mm","mmm")
  c, c2 := f3()
  d, d2 := f4()
  fmt.Println(a,c,d,c2,d2)
}

// ノーマルな関数
func f1(a, b int) int {
  return a + b
}


// 可変長引数
func f2(s string, params ...string) {
  for _, p:=range params {
    fmt.Println(p)
  }
}

// 複数の戻り値
func f3() (int, string) {
  return 100, "hoge"
}

// 戻り値に名前をつけて扱う
func f4() (i int, t int) {
  i = 5
  t = 2
  return
}
