package constructorTree

import "reflect"

func GetConstructorTree(headConstructor reflect.Type, constructors []reflect.Type) funcNode {
	return newFuncNode(headConstructor, constructors)
}

type funcNode struct {
	Signature string
	InputConstructors [][]funcNode
}

func newFuncNode(f reflect.Type, constructors []reflect.Type) funcNode {
	childs := searchForInputConstructors(f, constructors)
	signature := f.String()
	return funcNode{signature, childs}
}

func searchForInputConstructors(f reflect.Type, constructors []reflect.Type) [][]funcNode {
	inputConstructors := [][]funcNode{}
	for i := 0; i < f.NumIn(); i++ {
		fIn := f.In(i)
		inChilds := []funcNode{}
		for _, function := range constructors {
			if (canConstruct(function, fIn)){
				inChilds = append(inChilds, newFuncNode(function, constructors))
			}
		}
		inputConstructors = append(inputConstructors, inChilds)
	}
	return inputConstructors
}

func canConstruct(function reflect.Type, t reflect.Type) bool {
		constructs := false
		switch t.Kind() {
		case reflect.Interface:
			constructs = function.Out(0).Implements(t)
		case reflect.Slice:
			constructs = function.Out(0).Implements(t.Elem())
		default:
			constructs = false
		}
		return constructs
}

