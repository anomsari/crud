package model
import (
	"math"
	"fmt"

)
//balok struct
type Balok struct{
	Panjang int
	Lebar   int
	Tinggi  int
}

type Segitiga struct{
	Alas 		int
	Sisi		int
}


func (b *Balok)Keliling() int {
return 4*( b.Panjang + b.Lebar+ b.Tinggi)
}

func (b *Balok)Luas() int {
return 2* b.Panjang+ 2*b.Lebar+ 2* b.Tinggi
}

func (b *Balok)Volume() int {

	return (b.Panjang *b.Lebar *b.Tinggi)
}

func (s *Segitiga) Kelilings() int{
	fmt.Print(s.Sisi)
	return s.Sisi+s.Sisi+s.Alas

}


func (s *Segitiga) Tinggis() float64{

		tinggi 	:= float64(s.Sisi* s.Sisi) - (0.5* float64(s.Alas))* (0.5* float64(s.Alas))
		// t 			:= float64(tinggi)
		hasil 	:= math.Sqrt (tinggi)

	return hasil
}

func (s *Segitiga) Luass() float64{
		tinggi := Segitiga{Alas: s.Alas, Sisi: s.Sisi}

	return 0.5* float64(s.Alas)* tinggi.Tinggis()
}
