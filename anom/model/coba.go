package model
import (
  "fmt"
  "time"
  "strings"
)
//jumlah
func Penjumlahan(a,b int) int{
  return a + b
}
//kurang
func Pengurangan(a,b int) int{
  return a - b
}
//bagi
func Penbagian(a,b int) int{
  return a / b
}
//kali
func Perkalian(a,b int) int{
  return a * b
}
//pangkat
func Pemangkatan(a,b int) int{

  c:=1
  for i:= 1; i<=b; i++{
    c= a*c

  }
  return c
}

//factorial
func Faktorial(a int) int{

  i:=1
  for j:=1; j<=a; j++ {
        i= j*i
    }
  // for i:= 1; i<=b; i++{
  //   c= i*b
  return i
  }


// max slice
func Max (m[]int)int {
  if len(m) == 0 {
    return 0
  }
  maksimum := m [0]
  for _, nilai := range m {
    if nilai > maksimum {
        maksimum = nilai
    }
  }
  return maksimum
}

// min slice
func Min (m[]int)int {
  if len(m) == 0 {
    return 0
  }
  minimum := m [0]
  for _, nilai := range m {
    if nilai < minimum {
        minimum = nilai
    }
  }
  return minimum
}

// //bil ganjil
func BilGanjil (n int) (hasilGenap, hasilGanjil[]int){
  // var hasilGenap, hasilGanjil[]int
  for i:= 0; i<= n; i++{
    hasil := (i*2)
    hasilGanjil = append(hasilGanjil, hasil+1)
    hasilGenap = append(hasilGenap, hasil)
  }
  return
}

//Bilangan Genap
func BilGenap (n int) []int{
  var hasil []int
  for i:= 0; i< n; i++{
    genap := (i*2)
    hasil = append(hasil,genap)
  }
  return hasil
}

//Bilangan ganjil
func BilGan (n int) []int{
  var hasil []int
  for i:= 0; i< n; i++{
    genap := (i*2)
    hasil = append(hasil,genap+1)
  }
  return hasil
}

func TanggalLahir(waktu time.Time) string{

		tahun, month, hari := waktu.Date()

	 	bulan := int (month)

	  brithday := fmt.Sprintf("%v-%v-%v",hari, bulan, tahun)

	return brithday
}

func Stringnya (a int, b string ) string{
  var hasil string
    for i:= 0; i< a; i++{
      hasil += b
}
return hasil
}

func StringSplit(a []int,stringnya string) string {

  splitnya := strings.Split(stringnya, "/")
  keyMasuk := strings.Split(splitnya[2],"")

  return "20"+keyMasuk[3] + keyMasuk[4]
}
