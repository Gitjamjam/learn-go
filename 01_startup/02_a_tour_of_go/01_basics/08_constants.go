package main

import (
  "fmt"
)
// 定数は const を使って変数と同じように宣言
// 定数は文字(character)、文字(string)
// boolean、数値(numeric)のみ使用可能
// 定数は := は使えない
const Pi = 3.14

func main() {
  const World = "世界"
  fmt.Println("Hello", World)
  fmt.Println("Happy", Pi, "Day")

  const Truth = true
  fmt.Println("Go rules?", Truth)
}
