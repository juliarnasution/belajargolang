package src
import (
	"fmt"
)

// Fungsi tutorial function
func Fungsi(){
	var hasil int
	var pesan string
	jamKerja := 12
	hasil = hitungBonus(jamKerja)
	fmt.Println("jumlah bonusnya ", hasil)
	pesan = message(hasil)
	fmt.Println("pesannya ", pesan)
}

func hitungBonus (jamKerja int) int {
	if(jamKerja >=12){
		return jamKerja * 100
	}else if(jamKerja>4){
		return jamKerja * 50
	}

	return jamKerja * 10
}

func message(gaji int) string{
	var hasil string
		switch {
			case gaji>100: hasil =  "Rajin sekali"
				break
			case gaji>50 : hasil = "Lumayan"
				break
			default : hasil = "malas"
		}
	return hasil
}