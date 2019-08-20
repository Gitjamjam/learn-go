package main

import "fmt"

func main () {

  // deferステートメントは deferへ渡した関数の実行を、
  // 呼び出し元の関数の終わりまで遅延させるもの
  // deferへ渡した関数の引数はすぐに評価されますが、
  // その関数自体は呼び出し元の関数がreturnするまで実行されない
  defer fmt.Prinln("world")
  fmt.Println("hello")

  fmt.Println("counting")

  // deferへ渡した関数が複数あると stack される
  // 呼び出し元の関数がreturnするとき
  // deferへ渡した関数は LIFO で実行
  for i := 0; i < 10; i++ {
    defer fmt.Println(i)
  }
  fmt.Println("done")
}
