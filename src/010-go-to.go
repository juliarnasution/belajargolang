package src

import "fmt"

//Goto tes switch case
func Goto(){
	/* local variable definition */
	var a int = 10

	/* do loop execution */
	LOOP: for a < 20 {
		if a == 15 {
			/* skip the iteration */
			a = a + 1
			goto LOOP
		}
		fmt.Printf("value of a: %d\n", a)
		a++     
	}
}