package src

import (
	"fmt"
)

// OperasiMateMatika  untuk operasi matematika
func OperasiMateMatika()  {
	var hasil, nilai1, nilai2 int
	fmt.Println("masukkan angka pertama")
	fmt.Scanln(&nilai1)
	fmt.Println("masukkan angka kedua")
	fmt.Scanln(&nilai2)
	hasil = pembagian(nilai1, nilai2)
	fmt.Println("hasilnya adalah ",hasil)
	hasil++
	// hasil--
	fmt.Println("hasilnya ++ adalah ",hasil)
}

func perkalian (nilai1 int, nilai2 int) int{
	return nilai1 * nilai2
}

func pembagian (nilai1 int, nilai2 int) int{
	return nilai1 / nilai2
}

func pengurangan (nilai1 int, nilai2 int) int{
	return nilai1 - nilai2
}

func modulus (nilai1 int) int{
	return nilai1%2
}