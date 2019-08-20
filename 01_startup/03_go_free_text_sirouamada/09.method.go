package main

import "fmt"
// goのメソッドは、
// 特定の型に関連つけられた関数

type Calc struct {
  a, b int
}

type Mystruct struct {
  Calc
}

type (p Calc) Add() int {
  return p.a + p.b
}

type MyInt int

// メソッド Calc構造体に紐づく関数
func (p Calc) Add() int {
  return p.a + p.b
}

// メソッド MyInt型に紐づく関数
func (m MyInt) Add(n int) Myint {
  return m + MyInt(n)
}
func main() {
  p := Calc{3,2}
  var m MyInt = 1

  fmt.Println(p.Add())
  fmt.Println(m.Adds(2))

  var s Mystruct
  s.a = 5
  s.b = 4
  fmt.Println(s.Add())
}
