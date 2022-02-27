package main

import(
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))}

type hotel struct {
	Name, Address, City, Zip string
	Reg Region
}

type Region int

const (
	South Region = iota
	Central
	North
)

func (r Region) String() string {
	return [...]string{"South", "Central", "North"}[r]
}

var fm = template.FuncMap{
	"str": Region.String,
}

var hotels []hotel


func main() {

	hotels = []hotel{
		hotel{
			Name: "Hotel California",
			Address: "The Strand",
			City: "San Francisco",
			Zip: "94107",
			Reg: South,},
		hotel{
			Name: "Hotel California",
			Address: "The Strand",
			City: "San Francisco",
			Zip: "94107",
			Reg: North,},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", hotels)
	if err != nil {
		log.Fatalln(err)
	}
}