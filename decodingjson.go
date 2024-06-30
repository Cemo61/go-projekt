import (
	"fmt"
	"encoding/json"
	)

type Name struct {
	Vorname  string
	Nachname string
}
type Anschrift struct {
	Zeile1  string
	Adresse string
	PLZ     string
}
type Kunde struct {
	Name      Name
	Email     string
	Anschrift Anschrift
}

func main() {
	yasin := Kunde{
		Name: Name{Vorname: "Yasin",
			Nachname: "Yazgan",
		},
		Email: "yasinyazgan@beispiel.de",
		Anschrift: Anschrift{
			Zeile1:  "Bundeskanzleramt",
			Adresse: "Willy-Brandt-Straße 1",
			PLZ:     "123456 Berlin",
		},
	}
	//yasinJSON, fehl := json.Marshal(yasin)
	yasinJSON, fehl := json.MarshalIndent(yasin, "", "")
	if fehl == nil {
		fmt.Println(string(yasinJSON))
	} else {
		fmt.Println(fehl)
	}
}
