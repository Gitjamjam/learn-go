package main

import (
  "fmt"
)

func main() {

  // int 型のポインタ変数
  var p *int

  // int型の変数
  n := 10

  // 変数nのアドレスを取得
  p = &n

  fmt.Println(p)
  fmt.Println(*p)

  // 型を指定してメモリを割り当てる
  var mem *int = new(int)
  fmt.Println(mem)
  fmt.Println(*mem)

}
