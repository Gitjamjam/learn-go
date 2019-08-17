package main

import (
  "fmt"
  "runtime"
  "time"
)

func main() {
  fmt.Print("Go runs on")

  // C や java の switch と似ていますが、
  // 選択されたcaseだけを実行して
  // それ以降のcaseは実行されない
  switch os:= runtime.GOOS; os {
  case "darwin":
      fmt.Printls("OS X.")
  case "linux":
      fmt.Println("Linux.")
  default
      fmt.Printf("%s.", os)
  }

  fmt.Println("When's Saturday?")
  today := time.Now().Weekday()

  // switch case は 上から下へcaseを評価
  // case の条件が一致すればそこで停止(自動でbreak)
  switch time.Satuday {
  case today + 0:
      fmt.Println("today.")
  case today + 1:
      fmt.Println("Tomorrow")
  case today + 2:
      fmt.Println("In two days")
  default:
      fmt.Println("Too far away")
  }

  t := time.Now()
  // 条件のないswitch は switch trueと書くことと同じ
  switch {
  case t.Hour() < 12:
    fmt.Println("Good morning")
  case t.Hour() < 17:
    fmt.Println("good afternoon")
  default:
    fmt.Println("good evening")

  }

}
