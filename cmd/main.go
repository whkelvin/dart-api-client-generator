package main

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/whkelvin/dart-client-generator/internal/generator"
	"os"
)

func main() {
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromFile("./assets/openapi.json")
	if err != nil {
		fmt.Print(err)
		return

	}

	t := doc.Tags
	for i := 0; i < len(t); i++ {
		fmt.Println(t[i].Name)
	}

	tmplContent, err := os.ReadFile("./assets/templates/model.dart")
	if err != nil {
		fmt.Println(err.Error())
	}

	res, err := generator.GenerateModel(generator.ModelGeneratorOptions{
		ClassName: "Test",
		FileName:  "test",
	}, string(tmplContent))
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(res)
}
