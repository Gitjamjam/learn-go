package main

import "fmt"

// var ステートメントは変数宣言
// 関数の引数リストと同様に、複数変数の同型は最後だけにまとめられる
var c, python, java bool

// var宣言では変数ごとに初期値を付与できる
// 初期値はあるときは型を省略できる
var i, j int = 1, 2

func main() {
  var i int
  fmt.Println(i, c, python, java)

  // 関数の中では var宣言の代わりに
  // := の代入文を使って暗黙的型宣言が可能
  // 関数外では使用できない
  var i2, j2 int = 1, 2
  k := 3
  c2, python2, java2 := true, false, "NO"
  fmt.Println(i2, j2, k, c2, python2, java2)

}
