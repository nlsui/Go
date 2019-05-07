package constructorTree

import card "github.com/DecentralCardGame/cardobject"
import "fmt"
import "reflect"
import "encoding/json"
import "testing"


func TestConstructorTree(t *testing.T) {
	treeBuilder := NewTreeBuilder(getConstructors(), getConstants())
	tree := treeBuilder.BuildTreeFor(reflect.TypeOf(card.NewAction))
	bytes, err := json.Marshal(tree)
    if err != nil {
        fmt.Println("Can't serialize", tree)
    }
    fmt.Println(string(bytes))
}

func getConstructors() []reflect.Type {
	functions := []reflect.Type{}
	functions = append(functions, reflect.TypeOf(card.NewAction))
	functions = append(functions, reflect.TypeOf(card.NewCost))
	functions = append(functions, reflect.TypeOf(card.NewEffect))
	functions = append(functions, reflect.TypeOf(card.NewZoneChange))
	return functions
}

func getConstants() map[string]reflect.Type {
	m := make(map[string]reflect.Type)
	m["MANA"] = reflect.TypeOf(card.MANA)
	return m
}