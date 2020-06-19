package src

import "fmt"

// Profile adalah struct
 type Profile struct{//nama struct boleh huruf kecil semua
	id int
	nama string
	umur int
	hobi string
	tinggi float32
}
//Structure tes Structure
func Structure(){
	var  profile Profile//bisa lebih dari satu var/object
	profile.id = 1
	profile.nama = "juliar"
	profile.umur = 24
	profile.hobi = "tidur"
	profile.tinggi = 168
	fmt.Println(profile)
	printProf(profile)
}
func printProf( profile Profile ) {
	fmt.Printf( "profile id : %d\n", profile.id);
	fmt.Printf( "profile nama : %s\n", profile.nama);
	fmt.Printf( "profile hobi : %s\n", profile.hobi);
	fmt.Printf( "profile umur : %d\n", profile.umur);
 }