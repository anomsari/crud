package test

import(
  "fmt"
  "time"
  "testing"
  "anom/model"

)
func TestPerhitungan(t *testing.T){
//test penjumlahan
 t.Run("Test untuk Fungsi Penjumlahan", func(t *testing.T){
   var testPenjumlahan = []struct{
     a            int
     b            int
     hasilMaunya  int
   }{
     {a:1,  b:1,  hasilMaunya:2},
     {a:1,  b:3,  hasilMaunya:4},
     {a:4,  b:6,  hasilMaunya:10},
     {a:7,  b:11, hasilMaunya:18},
     {a:12, b:31, hasilMaunya:43},
     {a:-1, b:4,  hasilMaunya:3},
     {a:-8, b:1,  hasilMaunya:-7},
   }
   for _, input := range testPenjumlahan{
     hasilDapatnya := model.Penjumlahan(input.a, input.b)
     if hasilDapatnya != input.hasilMaunya{
       t.Fatalf("%v+%v = Dapatnya :%v, Maunya :%v\n",input.a, input.b, hasilDapatnya, input.hasilMaunya)
     }
     fmt.Printf("%v+%v = Dapatnya :%v, Maunya :%v\n",input.a, input.b, hasilDapatnya, input.hasilMaunya)
   }
 })
//test perkalian
  t.Run("Test untuk Fungsi Perkalian", func(t *testing.T){
    var testPerkalian = []struct{
      a            int
      b            int
      hasilMaunya  int
    }{
      {a:1,  b:1,  hasilMaunya:1},
      {a:1,  b:3,  hasilMaunya:3},
      {a:4,  b:6,  hasilMaunya:24},
      {a:7,  b:11, hasilMaunya:77},
      {a:12, b:31, hasilMaunya:372},
      {a:-1, b:4,  hasilMaunya:-4},
      {a:-8, b:1,  hasilMaunya:-8},
    }
    for _, input := range testPerkalian{
      hasilDapatnya := model.Perkalian(input.a, input.b)
      if hasilDapatnya != input.hasilMaunya{
        t.Fatalf("%v*%v = Dapatnya :%v, Maunya :%v\n",input.a, input.b, hasilDapatnya, input.hasilMaunya)
      }
      fmt.Printf("%v*%v = Dapatnya :%v, Maunya :%v\n",input.a, input.b, hasilDapatnya, input.hasilMaunya)
    }
  })

//test pemangkatan
 t.Run("Test untuk Fungsi Pemangkatan", func(t *testing.T){
   var testPemangkatan = []struct{
     a            int
     b            int
     hasilMaunya  int
   }{
     {a:1,  b:1,  hasilMaunya:1},
     {a:1,  b:3,  hasilMaunya:1},
     {a:4,  b:6,  hasilMaunya:4096},
     {a:7,  b:11, hasilMaunya:1977326743},
     {a:12, b:3,  hasilMaunya:1728},
     {a:-1, b:-4, hasilMaunya:1},
     {a:-8, b:1,  hasilMaunya:-8},
   }
   for _, input := range testPemangkatan{
     hasilDapatnya := model.Pemangkatan(input.a, input.b)
     if hasilDapatnya != input.hasilMaunya{
       t.Fatalf("%v^%v = Dapatnya :%v, Maunya :%v\n",input.a, input.b, hasilDapatnya, input.hasilMaunya)
     }
     fmt.Printf("%v^%v = Dapatnya :%v, Maunya :%v\n",input.a, input.b, hasilDapatnya, input.hasilMaunya)
   }
 })

 //test faktorial
  t.Run("Test untuk Fungsi faktorial", func(t *testing.T){
    var testFaktorial = []struct{
      a            int
      b            int
      hasilMaunya  int
    }{
      {a:5,  hasilMaunya:120},
      {a:6,  hasilMaunya:720},
      {a:9,  hasilMaunya:362880},
      {a:3,  hasilMaunya:6},
      {a:12, hasilMaunya:479001600},

    }
    for _, input := range testFaktorial{
      hasilDapatnya := model.Faktorial(input.a)
      if hasilDapatnya != input.hasilMaunya{
        t.Fatalf("%v = Dapatnya :%v, Maunya :%v\n",input.a, hasilDapatnya, input.hasilMaunya)
      }
      fmt.Printf("%v= Dapatnya :%v, Maunya :%v\n",input.a, hasilDapatnya, input.hasilMaunya)
    }
  })

// test max slice
t.Run("Test untuk Fungsi Max", func(t *testing.T){
  var testMax = []struct{
    a            []int
    hasilMaunya  int
  }{
    {a: []int{5,10,12,13,18,8},  hasilMaunya:18},
    {a: []int{5,90,32,16,98,80}, hasilMaunya:98},
    {a: []int{2,50,72,93,108,8}, hasilMaunya:108},
    {a: []int{1,10,11,13,5,8},   hasilMaunya:13},
  }
  for _, input := range testMax{
    hasilDapatnya := model.Max(input.a )
    if hasilDapatnya != input.hasilMaunya{
      t.Fatalf("%v = Dapatnya :%v, Maunya :%v\n",input.a,  hasilDapatnya, input.hasilMaunya)
    }
    fmt.Printf("%v = Dapatnya :%v, Maunya :%v\n",input.a, hasilDapatnya, input.hasilMaunya)
  }
})

t.Run("Test untuk Fungsi Min", func(t *testing.T){
  var testMin = []struct{
    a            []int
    hasilMaunya  int
  }{
    {a: []int{5,10,12,13,18,8},  hasilMaunya:5},
    {a: []int{5,90,32,16,98,80}, hasilMaunya:5},
    {a: []int{2,50,72,93,108,8}, hasilMaunya:2},
    {a: []int{1,10,11,13,5,8},   hasilMaunya:1},
  }
  for _, input := range testMin{
    hasilDapatnya := model.Min(input.a )
    if hasilDapatnya != input.hasilMaunya{
      t.Fatalf("%v = Dapatnya :%v, Maunya :%v\n",input.a,  hasilDapatnya, input.hasilMaunya)
    }
    fmt.Printf("%v = Dapatnya :%v, Maunya :%v\n",input.a, hasilDapatnya, input.hasilMaunya)
  }
})

// bilangan Ganjil
t.Run("Test untuk Fungsi Bilangan Ganjil", func(t *testing.T){
  var testGanjil = []struct{
    a            int
    hasilMaunya  []int
  }{
    {a:3, hasilMaunya: []int{1,3,5}},
    {a:5, hasilMaunya: []int{1,3,5,7,9}},
    {a:6, hasilMaunya: []int{1,3,5,7,9,11}},
    {a:9, hasilMaunya: []int{1,3,5,7,9,11,13,15,17}},
    {a:1, hasilMaunya: []int{1,}},
  }
  for _, input := range testGanjil{
    hasilDapatnya := model.BilGan(input.a )
    if len(hasilDapatnya) != len (input.hasilMaunya){
      t.Errorf("leng maunya :%v, leng dapatnya :%v\n", input.hasilMaunya, hasilDapatnya )
    }
    for index, value := range hasilDapatnya {
      if value != input.hasilMaunya[index]{
        t.Fatalf("Nilainya %v, dapatnya :%v, maunya %v\n", input.a, input.hasilMaunya[index], hasilDapatnya[index] )
      }
      fmt.Printf("Nilainya %v, dapatnya :%v, maunya %v\n", input.a, input.hasilMaunya[index], hasilDapatnya[index] )
    }

    // fmt.Printf("%v = Dapatnya :%v, Maunya :%v\n",input.a, hasilDapatnya, input.hasilMaunya)
  }
})
//
// bilangan Genap
t.Run("Test untuk Fungsi Bilangan Genap", func(t *testing.T){
  var testGenap = []struct{
    a            int
    hasilMaunya  []int
  }{
    {a:3, hasilMaunya: []int{0,2,4}},
    {a:5, hasilMaunya: []int{0,2,4,6,8}},
    {a:6, hasilMaunya: []int{0,2,4,6,8,10}},
    {a:9, hasilMaunya: []int{0,2,4,6,8,10,12,14,16}},

  }
  for _, input := range testGenap{
    hasilDapatnya := model.BilGenap(input.a)
    if len(hasilDapatnya) != len(input.hasilMaunya){
      t.Errorf("leng maunya :%v, leng dapatnya :%v\n", len(hasilDapatnya), len(input.hasilMaunya) )
    }
    for index, value := range hasilDapatnya {
      if value != input.hasilMaunya[index]{
        t.Fatalf("Nilainya %v, dapatnya :%v, maunya %v\n", input.a, input.hasilMaunya[index], value )
      }
      fmt.Printf("Nilainya %v, dapatnya :%v, maunya %v\n", input.a, input.hasilMaunya[index], value )
    }

    // fmt.Printf("%v = Dapatnya :%v, Maunya :%v\n",input.a, hasilDapatnya, input.hasilMaunya)
  }

})

//gabung ganjil Genap
t.Run("Test untuk Fungsi Bilangan Ganjil Genap", func(t *testing.T){
  var testGenap = []struct{
    a            int
    Ganjil      []int
    Genap       []int
  }{
    {a:3, Genap: []int{0,2,4}, Ganjil: []int{1,3,5}},
    {a:5, Genap: []int{0,2,4,6,8}, Ganjil: []int{1,3,5,7,9}},

  }
  for _, input := range testGenap{
    hasilDapatnya := model.BilGenap(input.a)
    if len(hasilDapatnya) != len(input.Genap){
      t.Errorf("leng maunya :%v, leng dapatnya :%v\n", len(hasilDapatnya), len(input.Genap) )
    }
    for index, value := range hasilDapatnya {
      if value != input.Genap[index]{
        t.Fatalf("Nilainya %v, dapatnya :%v, maunya %v\n", input.a, input.Genap[index], value )
      }
      fmt.Printf("Nilainya %v, dapatnya :%v, maunya %v\n", input.a, input.Genap[index], value )
    }

    // fmt.Printf("%v = Dapatnya :%v, Maunya :%v\n",input.a, hasilDapatnya, input.hasilMaunya)
  }

  for _, input := range testGenap{
    hasilDapatnya := model.BilGan(input.a )
    if len(hasilDapatnya) != len (input.Ganjil){
      t.Errorf("leng maunya :%v, leng dapatnya :%v\n", input.Ganjil, hasilDapatnya )
    }
    for index, value := range hasilDapatnya {
      if value != input.Ganjil[index]{
        t.Fatalf("Nilainya %v, dapatnya :%v, maunya %v\n", input.a, input.Ganjil[index], hasilDapatnya[index] )
      }
      fmt.Printf("Nilainya %v, dapatnya :%v, maunya %v\n", input.a, input.Ganjil[index], hasilDapatnya[index] )
    }

    // fmt.Printf("%v = Dapatnya :%v, Maunya :%v\n",input.a, hasilDapatnya, input.hasilMaunya)
  }

})
// tanggal
t.Run("Test untuk Konversi Tanggal", func(t *testing.T){
  var testKonversiTanggal = []struct{
    date            time.Time
    tanggalMaunya   string
  }{
    {date: time.Date(1996,3,28,0,0,0,0, time.UTC),tanggalMaunya : "28-3-1996"},
    {date: time.Date(1995,2,13,0,0,0,0, time.UTC),tanggalMaunya : "13-2-1995"},
    {date: time.Date(1994,6,16,0,0,0,0, time.UTC),tanggalMaunya : "16-6-1994"},
    {date: time.Date(1992,11,8,0,0,0,0, time.UTC),tanggalMaunya : "8-11-1992"},
    {date: time.Date(1999,12,18,0,0,0,0, time.UTC),tanggalMaunya : "18-12-1999"},
    {date: time.Date(1988,1,2,0,0,0,0, time.UTC),tanggalMaunya : "2-1-1988"},

  }
  for _, input := range testKonversiTanggal{
    tanggalDapatnya := model.TanggalLahir(input.date )
    if tanggalDapatnya != input.tanggalMaunya{
      t.Fatalf("time.Time-nya = %v, Dapat stringnya, =%v, Mau strinya =%v\n",input.date,  tanggalDapatnya, input.tanggalMaunya)
    }
      fmt.Printf("time.Time-nya = %v, Dapat stringnya, =%v, Mau strinya =%v\n",input.date,  tanggalDapatnya, input.tanggalMaunya)
    }
})

// print string
t.Run("Test untuk Fungsi Penjumlahan", func(t *testing.T){
  var testString = []struct{
    a            int
    b            string
    mauString    string
  }{
    {a:1,  b:"a",   mauString:"a"},
    {a:2,  b:"aa",  mauString:"aaaa"},
    {a:2,  b:"ba",  mauString:"baba"},
    {a:3,  b:"a",   mauString:"aaa"},
    {a:4,  b:"abc", mauString:"abcabcabcabc"},

  }
  for _, input := range testString{
    hasilStringnya := model.Stringnya(input.a, input.b)
    if hasilStringnya != input.mauString{
      t.Fatalf("Angkanya = %v, Textnya : %v, Dapatnya :%v, Maunya :%v\n",input.a, input.b, hasilStringnya, input.mauString)
    }
    fmt.Printf("Angkanya = %v, Textnya : %v, Dapatnya :%v, Maunya :%v\n",input.a, input.b, hasilStringnya, input.mauString)
  }
})

//split
t.Run("Test untuk Fungsi split", func(t *testing.T){
  var testSplit = []struct{
    a            []int
    splitnya     string
    mauString    string
  }{
    {a:[]int{3,4}, splitnya:"www.gundarma.ac.id/student/10111359",  mauString:"2011"},
    {a:[]int{3,4}, splitnya:"www.gundarma.ac.id/student/10113893",  mauString:"2013"},
    {a:[]int{3,4}, splitnya:"www.gundarma.ac.id/student/51414419",  mauString:"2014"},
  }
  for _, input := range testSplit{
    hasilSplitnya := model.StringSplit(input.a, input.splitnya)
    if hasilSplitnya != input.mauString{
      t.Fatalf("Split1 = %v, Textnya =%v, Dapatnya :%v, Maunya :%v\n",input.a,  input.splitnya, hasilSplitnya, input.mauString)
    }
    fmt.Printf("Split1 = %v, Textnya =%v, Dapatnya :%v, Maunya :%v\n",input.a, input.splitnya, hasilSplitnya, input.mauString)
  }
})


 }
