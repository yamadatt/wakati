package main

import (
  "testing"
)

func TestFileRead(t *testing.T) {
  result := FileRead("sample.txt")
  expext := "あああ"
  if result != expext {
    t.Error("\n実際： ", result, "\n理想： ", expext)
  }

  t.Log("終了")
}