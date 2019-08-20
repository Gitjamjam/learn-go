package main

import "fmt"

func main() {

  // スライスは組み込みのmake 関数を使用して作成

  // make関数は ゼロ化された配列を割り当て、
  // その配列を指すスライスを返します
  a := make([]int, 5)
  printSlice("a", a)

  // make の 3番目の引数に容量を設定
  b := make([]int, 0, 5)
  printSlice("b", b)

  c := b[:2]
  printSlice("c", c)

  d := c[2:5]
  printSlice("d", d)

  // a len=5 cap=5 [0 0 0 0 0]
  // b len=0 cap=5 []
  // c len=2 cap=5 [0 0]
  // d len=3 cap=3 [0 0 0]
  
}

func printSlice(s string, x []int){
  fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
