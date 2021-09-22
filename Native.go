package main

import (
  "fmt"
  "log"
  "sort"
  "strings"
  "strconv"
)

var jenis, judul, vote, nama string
var arr1 [] string
var arr [][] string
var idsing int

func addData (){
  fmt.Print("\nMasukkan Judul Lagu : ")
  fmt.Scan(&jenis)
  fmt.Print("Masukkan Jumlah Vote (10 - 100) :")
  fmt.Scan(&vote)

  arr1 = []string{jenis, vote}
  arr = append(arr, arr1)

  sort.Slice(arr[:], func(i, j int)bool {
    return arr[i][1] > arr[j][1]
  })
}

func deleteData(ID int){
  var before, after [][]string
  before = arr[:ID-1]
  log.Printf("Data Before : ", before)
  arr = append(after, before...)
}

func countVote(){

}

func main (){
  var pilih int
  var x = true

  for x {
    fmt.Pritnln("\n\t====Daftar Pilihan==== ")
    fmt.Println("1. Input Data Penyanyi baru")
    fmt.Println("2. Hapus data penyanyi berdasarkan ID")
    fmt.Println("3. Tampilkan seluruh data penyanyi beserta jumlah data yang tersimpan dalam list")
    fmt.Println("4. Menampilkan top 3 musik terfavorit")
    fmt.Println("5. Menampilkan seluruh jumlah vote")
    fmt.Println("6. Menampilkan seluruh penyanyi dengan awalan huruf A: ")
    fmt.Println("0. Exit")
    fmt.Println("Masukkan Pilihan")
    fmt.Scan(&pilih)

    switch pilih {
    case 1:
      addData()

    case 2:
      fmt.Print("\nMasukkan ID penyanyi yang ingin di hapus: ")
      fmt.Scan(&idsing)

      deleteData(idsing)
      sort.Slice(arr[:], func(i, j int) bool {
        return arr[i][1] > arr[j][1]
      })

    case 3:
      fmt.Println("ID_Penyanyi \tNama Penyanyi \t\tVote")
      for i := 0; i < len(arr); i++ {
        var fltvote, err = strconv.ParseFloat(arr[i][1], 64)
        if err == nil {
          fmt.Printf("%d \t\t%s \t\t\t%s \t\t%n", i+1, arr[i][0], arr[i][1])
        }
      }
      fmt.Printf(""\nJumlah \t = %d", len(arr))
    case 4:
      fmt.Println("TOP 3 PENYANYI")
      fmt.Println("Rank \tNama Penyanyi \t\tVote")
      for i := 0; i <= 2; i++ {
        fmt.Printf("%d \t%s \t\t%s\n", i+1, arr[i][0], arr[i][1])
      }
    case 5:
    }
  }
}
