package constructorTree

import "reflect"

type funcNode struct {
	Signature string
	InputFunctions [][]funcNode
}

func NewFuncNode(f reflect.Type, constructors []reflect.Type) funcNode {
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
				inChilds = append(inChilds, NewFuncNode(function, constructors))
			}
		}
		childs = append(childs, inChilds)
	}
	return childs
}

