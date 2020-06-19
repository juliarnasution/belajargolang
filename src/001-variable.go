package src

import (
	"fmt"
	"math/rand"
)

var global = "juliar"

// documentation for variable https://www.tutorialspoint.com/go/go_data_types.htm

// Run adalah fungsi harus Capital agar bisa di export
func Run() {
	const nilaiConstant = 10//nilai ini tidak akan bole diganti
	var jumlah float32
	var local = "nasution"
	// static type
	var umur int16 = 24
	var alamat string = "mikuasi"
	sekolah  := "SMA Pakue" // shorthand ini hanya bisa di dalam fungsi
	peringkat :=2
	nilai := 9.9
	jumlah = float32(nilai) * float32(peringkat)// type casting

	fmt.Println("Nomor yang keluar : ", rand.Intn(10))
	fmt.Println("nama saya " + global +" "+local)
	fmt.Println("umur ",umur, " alamat "+alamat)
	fmt.Println("sekolah " + sekolah +" peringkat",peringkat," nilainya ", nilai)
	fmt.Println("perkalian ",jumlah)
}



