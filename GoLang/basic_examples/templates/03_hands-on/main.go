package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name, Address, City, Zip, Region string
}

type hotels []hotel

//var Regions = newRegion()

func newRegion() *regions {
    return &regions{
        Southern:   "Southern",
        Central: "Central",
        Northern:  "Northern",
    }
}

type regions struct {
    Southern   string
    Central string
    Northern  string
}

const city string = "California"

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	cHotels := hotels {
		hotel{
			Name: "Bir Numara",
			Address: "Restoranın Yanı",
			City: city,
			Zip: "10000",
			Region: newRegion().Southern,
		},
		hotel{
			Name: "No İki",
			Address: "Postanenin Arkası",
			City: city,
			Zip: "20000",
			Region: newRegion().Central,
		},
		hotel{
			Name: "ÜÇ",
			Address: "Eski AVM'nin Yeri",
			City: city,
			Zip: "30000",
			Region: newRegion().Northern,
		},
	}


	err := tpl.Execute(os.Stdout, cHotels)
	if err != nil {
		log.Fatalln(err)
	}
}
