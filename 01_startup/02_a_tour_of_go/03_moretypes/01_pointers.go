package main

import "fmt"

func main() {

  // 変数 T のポインタは *T型 ゼロ値はnull
  // & オペレータは、オペランドへのポインタを引き出す
  // * オペレータは、ポインタの指す先の変数を示します

  i, j := 42, 2701

  p := &i // point to i
  fmt.Println(*p) // read i through the pointer
  *p = 21 // set i through the pointer
  fmt.Println(i) // see the new value of i

  p = &j // point to j
  *p = *p / 37 // devide j through the pointer
  fmt.Println(j) // see the new value  of j
}
