package main

import "fmt"

//　インターフェースの宣言
type Human interface {
  hello()
  walk()
}
// 構造体とメソッドの宣言
type Masaru struct {
  name string
  age int
}
func (m Masaru) hello(){
  fmt.Printf("%s%dさい\n", m.name, m.age)
}

func (m Masaru) walk() {
  fmt.Println("walk...")
}
func (m Masaru) shout() {
  fmt.Println("shout....")
}

func main() {
  m := Masaru{"mmmm", 20}
  var h Human = m
  h.hello()
  h.walk()
}
