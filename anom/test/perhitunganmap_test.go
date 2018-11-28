
package test

import (
  "fmt"
  "testing"
  "anom/model"
)

//balok
func TestPerhitunganMap(t *testing.T){
  t.Run("Test untuk Fungsi Perhitungan Map Balok", func(t *testing.T){
    var testMap = []struct{
      key          string
      p            int
      l            int
      t            int
      hasilMaunya  interface{}
    }{
      {key: "Keliling", p:2,l:2,t:2, hasilMaunya:24},
      {key: "Luas", p:2,l:2,t:2, hasilMaunya:12},
      {key: "Volume", p:2,l:2,t:2, hasilMaunya:8},


    }
    for _, input := range testMap{
      hasilDapatnya, errorDapatnya := model.PerhitunganMapBalok(input.key, input.p, input.l, input.t)
      if hasilDapatnya != input.hasilMaunya || errorDapatnya != nil{
        t.Fatalf("Hasil Maunya= %v, Dapatnya :%v, Errornya :%v\n",input.hasilMaunya, hasilDapatnya, errorDapatnya)
      }
      fmt.Printf("Hasil Maunya= %v, Dapatnya :%v, Errornya :%v\n",input.hasilMaunya, hasilDapatnya, errorDapatnya)
    }
  })

// segitiga
  t.Run("Test untuk Fungsi Perhitungan Map Segitiga", func(t *testing.T){
    var testMapSegitiga = []struct{
      key          string
      s            int
      a            int
      hasilMaunya  interface{}
    }{
      {key: "Kelilings",s:1, a:10, hasilMaunya:21},
      {key: "Tinggis", s:6 ,a:5, hasilMaunya:4.0},//karena definisi awalnya float64 jadi harus pakai ((.0)) sebagai konversi ke int
      {key: "Luass", s:60, a:50, hasilMaunya:1200.0},

    }
    for _, input := range testMapSegitiga{
      hasilDapatnya, errorDapatnya := model.PerhitunganMapSegitiga(input.key, input.s, input.a)
      if hasilDapatnya != input.hasilMaunya || errorDapatnya != nil{
        t.Fatalf("Hasil Maunya= %v, Dapatnya :%v, Errornya :%v\n",input.hasilMaunya, hasilDapatnya, errorDapatnya)
      }
      fmt.Printf("Hasil Maunya= %v, Dapatnya :%v, Errornya :%v\n",input.hasilMaunya, hasilDapatnya, errorDapatnya)
    }
  })

  //PerhitunganKata
  t.Run("Test untuk Fungsi Perhitungan Map Kata", func(t *testing.T){
    var testMapKata = []struct{
      key          string
      hasilMaunya map[string]int
    }{
      {key:"Kelilings satu", hasilMaunya:map[string]int{"Kelilings":1 , "satu": 1}},
      {key:"jalan dua - duaan", hasilMaunya:map[string]int{"jalan":1 , "dua": 1, "-": 1, "duaan": 1}},
      {key:"pergi sendiri", hasilMaunya:map[string]int{"pergi":1 , "sendiri": 1}},
    }
    for _, input := range testMapKata{
      hasilDapatnya := model.PerhitunganKata(input.key)
      if len(hasilDapatnya) != len(input.hasilMaunya){
        t.Errorf("leng maunya :%v, leng dapatnya :%v\n", len(hasilDapatnya), len(input.hasilMaunya) )
      }
      for index, value := range hasilDapatnya {
        if value != input.hasilMaunya[index]{

        t.Fatalf("Hasil Maunya= %v, Dapatnya :%v\n",input.hasilMaunya, hasilDapatnya )
      }
      fmt.Printf("Hasil Maunya= %v, Dapatnya :%v\n",input.hasilMaunya, hasilDapatnya )
    }
  }
  })

}
