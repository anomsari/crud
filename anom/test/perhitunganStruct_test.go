
package test

import (
  "fmt"
  "testing"
  "anom/model"
)

var dataBalok = []struct {
  inputBalok    model.Balok
  mauKeliling   int
  mauVolume     int
  mauLuas       int
}{
  {inputBalok: model.Balok{Panjang:2, Lebar:3,Tinggi:4}, mauKeliling:36,mauVolume:24,mauLuas:18},
  {inputBalok: model.Balok{Panjang:2, Lebar:2,Tinggi:2}, mauKeliling:24,mauVolume:8,mauLuas:12},
  {inputBalok: model.Balok{Panjang:5, Lebar:4,Tinggi:3}, mauKeliling:48,mauVolume:60,mauLuas:24},
  }


  var dataSegitiga = []struct {
    inputSegitiga    model.Segitiga
    mauKelilingg      int
    // mauTinggii        float64
    mauLuass          float64
  }{
    {inputSegitiga: model.Segitiga{Alas:12, Sisi:10,}, mauKelilingg:32,mauLuass:48},
    {inputSegitiga: model.Segitiga{Alas:5, Sisi:4,}, mauKelilingg:13,mauLuass:7.806247497997997},
    }

func TestPerhitunganStruct(t *testing.T){
t.Run("Test untuk Fungsi Keliling Balok", func(t *testing.T){// Keliling
  for _, item := range dataBalok{
    hasilKelilingnya := item.inputBalok.Keliling()
    if hasilKelilingnya != item.mauKeliling{
      t.Fatalf("panjang = %v, lebar :%v, tinggi: %v, Dapatnya :%v, Maunya :%v\n", item.inputBalok.Panjang, item.inputBalok.Lebar, item.inputBalok.Tinggi, hasilKelilingnya, item.mauKeliling)
    }
    fmt.Printf("panjang = %v, lebar :%v, tinggi: %v, Dapatnya :%v, Maunya :%v\n", item.inputBalok.Panjang, item.inputBalok.Lebar, item.inputBalok.Tinggi, hasilKelilingnya, item.mauKeliling)
  }
})

t.Run("Test untuk Fungsi Volume Balok", func(t *testing.T){//volume
  for _, item := range dataBalok{
    hasilVolumenya := item.inputBalok.Volume()
    if hasilVolumenya != item.mauVolume{
      t.Fatalf("panjang = %v, lebar :%v, tinggi: %v, Dapatnya :%v, Maunya :%v\n", item.inputBalok.Panjang, item.inputBalok.Lebar, item.inputBalok.Tinggi, hasilVolumenya, item.mauVolume)
    }
    fmt.Printf("panjang = %v, lebar :%v, tinggi: %v, Dapatnya :%v, Maunya :%v\n", item.inputBalok.Panjang, item.inputBalok.Lebar, item.inputBalok.Tinggi, hasilVolumenya, item.mauVolume)
  }
})
t.Run("Test untuk Fungsi Luas Balok", func(t *testing.T){//luas
  for _, item := range dataBalok{
    hasilLuasnya := item.inputBalok.Luas()
    if hasilLuasnya != item.mauLuas{
      t.Fatalf("panjang = %v, lebar :%v, tinggi: %v, Dapatnya :%v, Maunya :%v\n", item.inputBalok.Panjang, item.inputBalok.Lebar, item.inputBalok.Tinggi, hasilLuasnya, item.mauLuas)
    }
    fmt.Printf("panjang = %v, lebar :%v, tinggi: %v, Dapatnya :%v, Maunya :%v\n", item.inputBalok.Panjang, item.inputBalok.Lebar, item.inputBalok.Tinggi, hasilLuasnya, item.mauLuas)
  }
})

//segitiga
t.Run("Test untuk Fungsi Luas Segitiga", func(t *testing.T){//luas
  for _, item := range dataSegitiga{
    hasilLuasnya := item.inputSegitiga.Luass()
    if hasilLuasnya != item.mauLuass{
      t.Fatalf("Alas = %v, Sisi: %v, Dapatnya :%v, Maunya :%v\n", item.inputSegitiga.Alas, item.inputSegitiga.Sisi, hasilLuasnya, item.mauLuass)
    }
    fmt.Printf("Alas = %v, Sisi: %v, Dapatnya :%v, Maunya :%v\n", item.inputSegitiga.Alas, item.inputSegitiga.Sisi, hasilLuasnya, item.mauLuass)
  }
})

t.Run("Test untuk Fungsi Luas Segitiga", func(t *testing.T){//Keliling
  for _, item := range dataSegitiga{
    hasilLuasnya := item.inputSegitiga.Kelilings()
    if hasilLuasnya != item.mauKelilingg{
      t.Fatalf("Alas = %v, Sisi: %v, Dapatnya :%v, Maunya :%v\n", item.inputSegitiga.Alas, item.inputSegitiga.Sisi, hasilLuasnya, item.mauKelilingg)
    }
    fmt.Printf("Alas = %v, Sisi: %v, Dapatnya :%v, Maunya :%v\n", item.inputSegitiga.Alas, item.inputSegitiga.Sisi, hasilLuasnya, item.mauKelilingg)
  }
})
}
