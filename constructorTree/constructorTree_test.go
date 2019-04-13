package constructorTree

import card "github.com/DecentralCardGame/cardobject"
import "fmt"
import "reflect"
import "encoding/json"
import "testing"


func TestConstructorTree(t *testing.T) {
	c := GetConstructorTree(reflect.TypeOf(card.NewAction), getConstructors())
	bytes, err := json.Marshal(c)
    if err != nil {
        fmt.Println("Can't serialize", c)
    }
    fmt.Println(string(bytes))
}

func getConstructors() []reflect.Type {
	functions := []reflect.Type{}
	functions = append(functions, reflect.TypeOf(card.NewAction))
	functions = append(functions, reflect.TypeOf(card.NewCost))
	return functions
}