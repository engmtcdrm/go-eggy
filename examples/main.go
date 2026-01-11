package main

import (
	"github.com/engmtcdrm/go-eggy"
	pp "github.com/engmtcdrm/go-prettyprint"

	"example.com/example/internal"
)

func main() {
	ex := eggy.NewExamplePrompt(internal.AllExamples).
		Title(pp.Yellow("Examples of Eggy"))
	ex.Show()
}
