package src

import "fmt"

//Pointer tes switch case
func Pointer(){
	var nama string = "juliar"

	var namaku *string = &nama //* ponter , & reference
	fmt.Println("alamat var namaku",namaku)// alamat value di memory
	fmt.Println("alamat var nama",&nama)// alamat value di memory
	fmt.Println("alamat var nama ke 2",namaku)// alamat value di memory
	fmt.Println("value var namaku = "+ *namaku)//tampil value
	*namaku ="nasution"
	fmt.Println("value var nama setelah ganti = "+ nama)//tampil value
}