package main

import "fmt"

func main() {
  sum := 0
  // for ループはセミコロン; で3箇所
  // 初期化; 条件式; 後処理;
  // 条件式が false となった場合にループを停止
  for i := 0; i < 10; i++ {
    sum += i
  }
  fmt.Println(sum)

  sum = 1
  // for ループの 初期化と後処理の記述は任意
  // 条件式は、ないと無限ループになるため必須
  for ; sum < 1000; {
    sum += sum
    fmt.Println(sum)
  }
  fmt.Println("sum = ",sum)

  // セミコロン; を省略できる
  // C言語などにある while は Go では forだけで実現
　sum = 1
  for sum < 1000 {
    sum += sum
  }
  fmt.Println("sum = ",sum)

  // ループ条件を省略すれば無限ループになる
  // infinite loop
  for {
    
  }

}
