package src

import (
	"fmt"
	// "strconv"
)

// MultipleReturn adalah contoh fungsi dengan multiple return
func MultipleReturn(){
	age := 24 
	name:= "juliar" 
	pekerjaan := "Programmer"
	namaDanPekerjaan,umur := getBiography(age,name,pekerjaan)
	fmt.Println(namaDanPekerjaan)
	fmt.Println("umur saya =", umur)
	address, number := alamatSaya(10,"trans sulawesi")
	fmt.Println("jalan "+ address)
	fmt.Println("nomor =", number)

}

/*	fungsi untuk return biography dengan return dua string
	return pertama dan kedua dipisah dengan tanda koma (,)
	*/
func getBiography(age int, name string, pekerjaan string) (string, int){
	// myAge := strconv.Itoa(age) // format int ke string
	myAge := age+2
	return "nama saya = "+ name + " pekerjaan = "+pekerjaan,//return string pertama
		// "umurnya adalah " + myAge // return string kedua
		myAge//return kedua berupa integer
		
}
// function dengan nama kembalian spesifik
func alamatSaya (nomor int, jalan string)(jalanKerumah string, nomorRumah int){
	jalanKerumah = "jalannya "+ jalan
	nomorRumah = nomor * 10
	return jalanKerumah,nomorRumah // kembalian string dan int
	/* jika kembalian hanya satu nama variable kemablian tidak harus ditulis 
	atau hanya dengan menuliskan "return" saja*/
			
}