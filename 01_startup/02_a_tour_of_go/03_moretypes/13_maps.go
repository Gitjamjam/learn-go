package main

import "fmt"

type Vertex struct {
  Lat, Long float64
}

var m map[string]Vertex

// map リテラルは structリテラルに似ているが、
// キーが必要
var m2 = map[string]Vertex{
  "Bell Labs": Vertex{
    40.68433, -74.39967,
  },
  "Google": Vertex{
    37.42202, -122.08408
  },
}

// もし map に渡すトプレベルの型がシンプルな型名なら
// リテラル要素から推定できるので、省略可能
var m3 = map[string]Vertex {
  "Bell Labs": {40.68433, -74.39967},
  "Google": {37.42202, -122.08408},
}

func main() {
  // マップはキーと値を関連付けします
  // マップのゼロ値はnil
  m = make(map[string]Vertex)
  m["Bell Labs"] = Vertex{
    40.68433, -74.39967,
  }
  fmt.Println(m["Bell Labs"])

  fmt.Println(m2)

  fmt.Println(m3)
}
