package model

import (
  "errors"
  "strings"
)

var buatMap = make(map[string]interface{})
var buatMapSegitiga = make(map[string]interface{})

// map balok
func PerhitunganMapBalok(s string,panjang int,lebar int, tinggi int)(interface{},error){
  balok := Balok{Panjang : panjang, Lebar : lebar, Tinggi : tinggi}
  buatMap[s] = balok

  switch s {
  case "Keliling":
    buatMap["Keliling"] = balok.Keliling()
    return buatMap["Keliling"], nil
  case "Luas":
    buatMap["Luas"] = balok.Luas()
    return buatMap["Luas"], nil
  case "Volume":
    buatMap["Volume"] = balok.Volume()
    return buatMap["Volume"], nil
  default:
    return nil, errors.New("key tidak ada!")
    }
  }

//map segitiga
    func PerhitunganMapSegitiga(b string, alas int, sisi int)(interface{},error){
      segitiga := Segitiga{Alas: alas, Sisi: sisi}
      buatMapSegitiga[b] = segitiga

      switch b {
      case "Kelilings":
        buatMapSegitiga["Kelilings"] = segitiga.Kelilings()
        return buatMapSegitiga["Kelilings"], nil
      case "Luass":
        buatMapSegitiga["Luass"] = segitiga.Luass()
        return buatMapSegitiga["Luass"], nil
      case"Tinggis":
        buatMapSegitiga["Tinggis"] = segitiga.Tinggis()
        return buatMapSegitiga["Tinggis"], nil

      default:
        return nil, errors.New("key tidak ada!")
        }

}

//map kata
func PerhitunganKata(s string) map[string]int{
  Map := make(map[string]int)
  theSplit := strings.Fields(s)
  for _, nilai :=range theSplit{
    Map[nilai] += 1
  }
  return Map
}
