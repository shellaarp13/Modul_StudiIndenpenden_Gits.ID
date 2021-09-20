package main

import "fmt"

func CountDigits(num int) int {
  count := 0

  if num > 0 {
    count++
    CountDigits (num / 10)
  }
  return count
}

func main (){
  var number int
  //count := 0

  fmt.Print("Enter a positive integer number: ")
  fmt.Scanf("%d", &number)

  count := CountDigits(number)
  fmt.Println("Total digits in number is : ", number, count)
}
