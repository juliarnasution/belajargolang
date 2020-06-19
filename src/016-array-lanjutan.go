package src
import "fmt"

// ArrayLanjutan adalah pembahan mendalam array
func ArrayLanjutan (){
	var fruits = [4] string{"apple","grape","banana","melon"}
	fmt.Println("jumlah element array fruits")
	fmt.Println("jumlah element array \t\t",len(fruits))
	fmt.Println("kapasitas \t\t",cap(fruits));
	fmt.Println("isi element \t\t", fruits)

	var buah = [...] int{1,2,3,4}
	fmt.Println("jumlah element array buah")
	fmt.Println("jumlah element array \t\t",len(buah))
	fmt.Println("kapasitas \t\t",cap(buah));
	fmt.Println("isi element \t\t", buah)

	var sayur = [2][3] string{[3]string {"paku","bayam","kangkung"},[3]string {"nangka","lodeh","kelor"}}
	fmt.Println("jumlah element array sayur")
	fmt.Println("jumlah element array \t\t",len(sayur))
	fmt.Println("kapasitas \t\t",cap(sayur));
	fmt.Println("isi element \t\t", sayur)
}