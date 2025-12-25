package main

import (
	"github.com/engmtcdrm/go-eggy"
	pp "github.com/engmtcdrm/go-prettyprint"

	"example.com/example/examples"
)

func main() {
	ex := eggy.NewExamplePrompt(examples.AllExamples).
		Title(pp.Yellow("Examples of Eggy"))
	ex.Show()
}
