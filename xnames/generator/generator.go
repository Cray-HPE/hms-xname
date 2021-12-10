package main

import (
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/Cray-HPE/hms-xname/xnametypes"
)

type XnameTypeNode struct {
	Parent   *XnameTypeNode
	Children []*XnameTypeNode

	Entry  xnametypes.HMSCompRecognitionEntry
	Fields []string

	FunctionParameter string
}

func main() {
	//
	// Build up XnameHierarchyNodes
	//
	typesToIgnore := map[xnametypes.HMSType]bool{}
	typesToIgnore[xnametypes.SMSBox] = true
	typesToIgnore[xnametypes.Partition] = true
	typesToIgnore[xnametypes.HMSTypeAll] = true
	typesToIgnore[xnametypes.HMSTypeAllComp] = true
	typesToIgnore[xnametypes.HMSTypeAllSvc] = true

	nodes := map[xnametypes.HMSType]*XnameTypeNode{}

	for _, entry := range xnametypes.GetHMSCompRecognitionTable() {
		if typesToIgnore[entry.Type] {
			continue
		}

		if _, exists := nodes[entry.Type]; exists {
			panic(fmt.Errorf("Error: entry type already exists: %v", entry))
		}

		nodes[entry.Type] = &XnameTypeNode{
			Entry: entry,
		}
	}

	//
	// Link nodes together
	//
	for _, node := range nodes {
		if node.Entry.ParentType == xnametypes.HMSTypeInvalid {
			continue
		}

		parentNode, parentExists := nodes[node.Entry.ParentType]
		if !parentExists {
			panic(fmt.Errorf("Error: parent type (%v) does not exist for type (%v) ", node.Entry.ParentType, node.Entry.Type))
		}

		// Update links
		node.Parent = parentNode
		parentNode.Children = append(parentNode.Children, node)
	}

	//
	// Sort children
	//
	for _, node := range nodes {
		sort.Slice(node.Children, func(i, j int) bool {
			return node.Children[i].Entry.Type < node.Children[j].Entry.Type
		})
	}

	// Determine fields
	for _, node := range nodes {
		node.Fields = GetFields(node)
		// node.ParentFields = GetFields(node.Parent)

		typeName := string(node.Entry.Type)
		node.FunctionParameter = strings.ToLower(string(typeName[0])) + typeName[1:]
	}

	//
	// The type tree
	//
	root := nodes[xnametypes.System]
	typeNames := GetTypeNames(root)
	// for _, t := range typeNames {
	// 	fmt.Println(t)
	// }

	xnameTypes := []*XnameTypeNode{}
	for _, typeName := range typeNames {
		xnameTypes = append(xnameTypes, nodes[typeName])
	}

	//
	// Template
	//
	TemplateFile("./generator/types.go.tpl", "./types.go", xnameTypes)
	TemplateFile("./generator/util.go.tpl", "./util.go", xnameTypes)
}

func Traverse(node *XnameTypeNode, level int) {
	fmt.Printf("%s- %v\n", strings.Repeat("  ", level), node.Entry.Type)
	for _, child := range node.Children {
		Traverse(child, level+1)
	}
}

func GetTypeNames(node *XnameTypeNode) []xnametypes.HMSType {
	types := []xnametypes.HMSType{node.Entry.Type}

	for _, child := range node.Children {
		types = append(types, GetTypeNames(child)...)
	}

	return types
}

func GetFields(node *XnameTypeNode) []string {
	if node == nil {
		return nil
	}

	if node.Entry.Type == xnametypes.System {
		return nil
	}

	return append(GetFields(node.Parent), string(node.Entry.Type))
}

func TemplateFile(sourceFilePath, destFilePath string, xnameTypes []*XnameTypeNode) {

	fmt.Println("Templating", sourceFilePath)
	t, err := template.ParseFiles(sourceFilePath)
	if err != nil {
		panic(err)
	}
	f, err := os.Create(destFilePath)
	if err != nil {
		panic(err)
	}
	if err := t.Execute(f, xnameTypes); err != nil {
		panic(err)
	}

	fmt.Println("Formatting", destFilePath)
	cmd := exec.Command("go", "fmt", destFilePath)
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
