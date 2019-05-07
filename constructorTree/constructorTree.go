package constructorTree

import "reflect"
import "fmt"

type TreeBuilder interface {
	BuildTreeFor(reflect.Type) funcNode
}

func NewTreeBuilder(constructors []reflect.Type, constants map[string]reflect.Type) TreeBuilder {
	return treeBuilder{constructors, constants, funcNode{}}
}

type treeBuilder struct {
	constructors []reflect.Type
	constants map[string]reflect.Type
	tree funcNode
}

func (builder treeBuilder) BuildTreeFor(headConstructor reflect.Type) funcNode {
	return builder.newFuncNode(headConstructor)
}

type funcNode struct {
	Signature string
	InputConstructors [][]funcNode
}

func (builder treeBuilder) newFuncNode(f reflect.Type) funcNode {
	childs := [][]funcNode{}
	for i := 0; i < f.NumIn(); i++ {
		childs = append(childs, builder.searchForParameter(f.In(i)))
	}
	signature := f.String()
	return funcNode{signature, childs}
}

func (builder treeBuilder) searchForParameter(parameterType reflect.Type) []funcNode {
	parameter := []funcNode{}
	searchType := parameterType
	switch parameterType.Kind() {
	case reflect.Slice:
		searchType = parameterType.Elem()
	default:
		searchType = parameterType
	}
	parameter = append(parameter, builder.searchForParameterConstructors(searchType)...)
	parameter = append(parameter, builder.searchForParameterConstants(searchType)...)
	return parameter
}

func (builder treeBuilder) searchForParameterConstructors(parameterType reflect.Type) []funcNode {
	inChilds := []funcNode{}
	for _, function := range builder.constructors {
		switch function.Kind() {
		case reflect.Func:
			if (function.Out(0).Implements(parameterType)){
				inChilds = append(inChilds, builder.newFuncNode(function))
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

func (builder treeBuilder) searchForParameterConstants(parameterType reflect.Type) []funcNode {
	inChilds := []funcNode{}
	for constantName, constantType := range builder.constants {
		if (constantType.Implements(parameterType)) {
			inChilds = append(inChilds, funcNode{constantName, [][]funcNode{}})
		}
	}
	return inChilds
}


