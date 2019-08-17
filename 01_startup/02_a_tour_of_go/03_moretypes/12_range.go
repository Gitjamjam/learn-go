package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {

  // for ループに使用する rangeは スライスやmapを
  // 一つずつ反復処理するために使用
  // スライスをrangeで繰り返す場合、
  // rangeは反復ごとに2つの変数を返します
  // 1つ目はインデックス
  // 2つ目はインデックスの場所の要素のコピー
  for i, v := range pow {
    fmt.Printf("2**%d = %d\n",i, v)
  }

  pow2 := make([]int, 10)
  for i := range pow2 {
    pow2[i] = 1 << uint(i)
  }
  // インデックスや値は、アンダーバー_ へ
  // 代入して捨てることができます

  // インデックスだけ必要なのなら
  // ", value"を省略しまうｓ
  for _, value := range pow2 {
    fmt.Printf("%d\n", value)
  }
}
