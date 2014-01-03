package main

import (
  "fmt"
  "os"
  "syscall"
  "text/scanner"
)

func main() {
  var Stdin = os.NewFile(uintptr(syscall.Stdin), "/dev/stdin")
  var s scanner.Scanner
  s.Init(Stdin)
  s.Mode = scanner.ScanIdents | scanner.ScanFloats | scanner.ScanChars | scanner.ScanStrings | scanner.ScanRawStrings | scanner.ScanComments
  tok := s.Scan()
  for tok != scanner.EOF {
    switch tok {
    case scanner.Ident:
      fmt.Println("Ident")
    case scanner.Int:
      fmt.Println("Int")
    case scanner.Float:
      fmt.Println("Float")
    case scanner.Char:
      fmt.Println("Char")
    case scanner.String:
      fmt.Println("String")
    case scanner.RawString:
      fmt.Println("RawString")
    case scanner.Comment:
      fmt.Println("Comment")
    }
    tok = s.Scan()
  }
}
