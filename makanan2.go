package main

import (
  "fmt"
  "errors"
  "strings"
)

type menu []string

func (x menu) has(food string) bool {
	food = strings.ToLower(food)
	for _, i := range x {
		if i == food {
			return true
		}
	}
	return false
}

func validate(input string) (bool, error){
  if strings.TrimSpace(input) == ""{
    return false, errors.New("Mohon masukan pesanan anda")
  }
  return true, nil
}

func showData(x menu) {
	for _, food := range x {
		fmt.Println("Pesanan anda :", food)
	}
}

func showMenu() {
	fmt.Println("Toko Makanan Indonesia")
	fmt.Println("======================")
	fmt.Println("Tahu")
	fmt.Println("Tempe")
	fmt.Println("Sambel")
	fmt.Println("Nasi")
	fmt.Println("Lele")
	fmt.Println("Ayam")
	fmt.Println("======================")
}


func getInput(data menu, x menu) {
	var input, inp string
	for cond1 := true; cond1; cond1 = strings.ToLower(inp) == "y" {
		for cond2 := true; cond2; cond2 = !x.has(input) {
			fmt.Print("Masukan menu pesanan anda dalam huruf (eg: Tahu) : ")
			fmt.Scanln(&input)
      if valid, err := validate(input) ; valid {
        showData(data)
      } else {
        fmt.Println(err.Error())
      }
			if x.has(input) {
				data = append(data, strings.ToLower(input))
			}
		}
		fmt.Print("Lanjutkan memesan ? (Y/T) : ")
		fmt.Scanln(&inp)
	}
}
func main() {
	var data menu
	menu2 := menu{"tahu", "tempe", "sambel", "nasi", "lele", "ayam"}
	showMenu()
	getInput(data, menu2)
  fmt.Println("Terima kasih atas pesanan anda")
}
