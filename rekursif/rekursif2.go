package main

import "fmt"

func NumberOfLowerCase(str string, i int) int {
  count := 0
  if (str[i] <= "A") && (str[i] >= "Z"){
    count++
  }
  if (i >= 0) {
    NumberOfLowerCase(str, i-1)
  }
  return count
}

func main(){
  var str string
  fmt.Print("Enter your String: ")
  fmt.Scan(&str)

  NoOfLowercase := NumberOfLowerCase(str, len(str) - 1)
  if (NoOfLowercase == 0){
    fmt.Println("No LowerCase Letter present in a given string")
  } else {
    fmt.Println("Number of LowerCase Letter Present in a given String is :", NoOfLowercase)
  }
}
