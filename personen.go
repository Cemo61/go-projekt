package main

import (
	"fmt"
	"encoding/json"
	)

type Personen struct {
	Vorname  string
	Nachname string
}

func main() {
	var clubmitglieder []Personen

	jsonString := `[
					{
					"vorname":"Cemil",
					"nachname":"Aydin"
					},

					{
					"vorname": "Yasin",
					"nachname": "Yazgan"
					},

					{
					"vorname": "Duha",
					"nachname": "Aydin"
					}
					]`

	json.Unmarshal([]byte(jsonString), &clubmitglieder)

	for _, person := range clubmitglieder {
		fmt.Println(person.Vorname, person.Nachname)
		//fmt.Println(person.Nachname)
	}
}

