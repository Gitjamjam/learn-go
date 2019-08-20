package main

import (
  "fmt"
  "time"
)

var n = 0

func main() {
  fmt.Println("main実行")
  fmt.Println(&n)

  // チャネルの生成
  ch := make(chan int)

  go send(ch)

  // チャネルのよる値の受信
  c := <-ch
  fmt.Printf("値を受信：%d\n", c)

  // ごルーチンの生成
  go sub()

  // 3ミリ秒スリープ
  time.Sleep(time.Millisecond * 10)

}
func sub() {
  fmt.Println("sub実行")
  fmt.Println(&n)

  // 10ミリ秒スリープ
  time.Sleep(time.Millisecond * 20)

  fmt.Println("sub実行2")
}

func send(ch chan int) {
  fmt.Println("値を送信：1")
  ch <- 1
}
