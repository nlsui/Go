package constructorTree

import "reflect"

func GetConstructorTree(headConstructor reflect.Type, constructors []reflect.Type) funcNode {
	return newFuncNode(headConstructor, constructors)
}

type funcNode struct {
	Signature string
	InputFunctions [][]funcNode
}

func newFuncNode(f reflect.Type, constructors []reflect.Type) funcNode {
	childs := searchForChilds(f, constructors)
	n := f.String()
	return funcNode{n, childs}
}

func searchForChilds(f reflect.Type, constructors []reflect.Type) [][]funcNode {
	childs := [][]funcNode{}
	for i := 0; i < f.NumIn(); i++ {
		fIn := f.In(i)
		inChilds := []funcNode{}
		for _, function := range constructors {
			if (fIn == function.Out(0)) {
				inChilds = append(inChilds, newFuncNode(function, constructors))
			}
		}
		childs = append(childs, inChilds)
	}
	return childs
}

