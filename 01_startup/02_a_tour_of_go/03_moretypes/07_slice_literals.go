package main

import "fmt"

func main() {
  s := []int{2, 3, 5, 7, 11, 13}

  // スライスするとき、上限または下限を省略できる
  // 下限の既定値は 0
  // 上限の既定値はスライスの長さ
  s = s[1:4]
  fmt.Println(s)

  s = s[:2]
  fmt.Println(s)

  s = s[1:]
  fmt.Println(s)

}
