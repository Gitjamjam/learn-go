package main

import (
  "fmt"
  "os"
  "errors"
)

func test() error {
  return errors.New("errors.Newでエラー生成")
}

func main() {
  ee := test()
  fmt.Println(ee.Error())

  // os.Open関数はファイルを開くための関数
  // 存在しないファイルを開こうとしてエラー
  _, e := os.Open("./hoge,txt")
  if (e != nil) {
    fmt.Println(e.Error())
    return
  }
}
