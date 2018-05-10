//not in the mode to use test tables
package invasion

import (
	"testing"
	"fmt"
)

const test_file = "../testmap"

var world WorldMap


func TestCityParser(t *testing.T) {
	l := "Foo north=Bar west=Baz south=Qu-ux"
	city := ParseCity(l)

	if city.name != "Foo" {
		t.Error("For", city.name, "Got", city.name, "Expected", "Foo") 
	}
	
}




func TestCitiesParser(t *testing.T) {
	cities, e := ParseCities(test_file)

	if e != nil {
		t.Error("File Does Not Exist")
	}

	expected := City{"Foo", []Connection{
		Connection{direction:"north", city:"Bar"},
		Connection{city:"Baz", direction: "west"},
		Connection{direction:"south", city:"Qu-ux"},
	}}

	
	got := cities[0]

	if fmt.Sprintf("%v", expected) != fmt.Sprintf("%v", got) {
		t.Error( "Got", got, "Expected", expected)
	}

	
}




func TestDestroy(t *testing.T) {
	WM, _ := DecodeWorldMap(test_file)
	WM.DestroyCity("Foo")

	p := func(C Connection) bool {
		return C.city == "Foo"
	}

	city := WM.cities["Bar"]
	
	if city.Contains(p) {
		t.Error("Bar is still connected to foo")
		fmt.Printf("%v", city)
	}  

}






func TestMove(t *testing.T) {

	
	WM, _ := DecodeWorldMap(test_file)
	WM.InitAliens(1)
	
	alien := WM.aliens[1]

	go func() {
		for {
			msg := <- alien.ch
			fmt.Printf("%s \n", msg.status)
		} 
	}()

	
	oldCity := alien.location
	
	WM.Move(alien)
	
	
	alien1 := WM.aliens[1]
	
	if (alien1.moveCtr != 2) {
		t.Error("Alien Counter Was not incremented")
	}

	if (alien1.location == oldCity) {
		t.Error("Alien city hasn't changed") 
	}
	
} 

