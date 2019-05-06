package constructorTree

import "reflect"
import "fmt"

func GetConstructorTree(headConstructor reflect.Type, constructors []reflect.Type) funcNode {
	return newFuncNode(headConstructor, constructors)
}

type funcNode struct {
	Signature string
	InputConstructors [][]funcNode
}

func newFuncNode(f reflect.Type, constructors []reflect.Type) funcNode {
	childs := [][]funcNode{}
	for i := 0; i < f.NumIn(); i++ {
		childs = append(childs, searchForParameter(f.In(i), constructors))
	}
	signature := f.String()
	return funcNode{signature, childs}
}

func searchForParameter(parameterType reflect.Type, constructors []reflect.Type) []funcNode {
	parameter := []funcNode{}
	searchType := parameterType
	switch parameterType.Kind() {
	case reflect.Slice:
		searchType = parameterType.Elem()
	default:
		searchType = parameterType
	}
	parameter = append(parameter, searchForParameterConstructors(searchType, constructors)...)
	return parameter
}

func searchForParameterConstructors(parameterType reflect.Type, constructors []reflect.Type) []funcNode {
	inChilds := []funcNode{}
	for _, function := range constructors {
		switch function.Kind() {
		case reflect.Func:
			if (function.Out(0).Implements(parameterType)){
				inChilds = append(inChilds, newFuncNode(function, constructors))
			}
		case reflect.Int:
			if (function.Implements(parameterType)){
				inChilds = append(inChilds, funcNode{function.String(), [][]funcNode{}})
			}
		default:
			fmt.Println(function.Kind())
		}
	}
	return inChilds
}


