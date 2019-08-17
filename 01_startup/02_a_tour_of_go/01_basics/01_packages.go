// Goのプログラムはパッケージで構成
// プログラムは main パッケージから開始
package main

import (
  "fmt"
  "math/rand"
)

func main(){
  fmt.Println("output comment",rand.Intn(10))
}
