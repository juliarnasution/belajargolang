package src

import "fmt"


/*SelectCase adalah select statement is similar to switch statement with difference that case statements refers to channel communications.
*/
func SelectCase(){
	var number,a chan int // key chan untuk membuat channel
	var nomor, b int
	select {
		case nomor =  <-number:
			fmt.Println("menerima ",nomor," dari <-number")
		case a <-b:
			fmt.Println("mengirim ",b," ke a")
		case b, ok := <-a:
			if ok {
				fmt.Println("menerima ", b, " from a")
			 } else {
				fmt.Println("a sudah tutup")
			 }
		default :
			fmt.Println("tidak ada komunikasi")
	}
}