package src

import "fmt"

//SwitchCase tes switch case
func SwitchCase(){
	var number int = 10
	switch number {
	case 1:
		fmt.Println("angka 1")//tidak perlu break
	case 10:
		fmt.Println("angka 10")
	default :
		fmt.Println("bukan 1 atau 10")
	}

	fmt.Println("------------")
	//switch case kedua
	nomor := 11
	switch {
	case nomor%2==0:
		fmt.Println("angka genap")
	case nomor%2!=0:
		fmt.Println("angka ganjil")
	default :
		fmt.Println("bukan angka")
	}
}