package src

import "fmt"

//Maps tes switch case
func Maps(){
	// declare a variable, by default map will be nil
	var countryCapitalMap map[string]string/* define the map as nil map can not be assigned any value*/
   	/* create a map*/
   	countryCapitalMap = make(map[string]string)//map akan bisa disi data
   
	/* insert key-value pairs in the map*/
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"
	
	/* print map using keys*/
	for country := range countryCapitalMap {
		fmt.Println("Capital of",country,"is",countryCapitalMap[country])
	}
	
	delete(countryCapitalMap,"France");
	fmt.Println("Entry for France is deleted") 

	for country := range countryCapitalMap {
		fmt.Println("Capital of",country,"is",countryCapitalMap[country])
	}
}

