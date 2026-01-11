// Package internal provides a collection of examples for this project
package internal

import (
	"fmt"

	"github.com/engmtcdrm/go-eggy"
)

var AllExamples = []eggy.Example{
	{Name: "No Repeat Example", Fn: NoRepeatExample},
	{Name: "Continue to Prompt Example", Fn: RepeatPromptExample},
}

var promptExamples = []eggy.Example{
	{Name: "Hello world!", Fn: HelloWorldExample},
	{Name: "Hi Mom!", Fn: HiMomExample},
}

func HelloWorldExample() {
	fmt.Println("Hello world!")
}

func HiMomExample() {
	fmt.Println("Hi Mom!")
}

func NoRepeatExample() {
	ex := eggy.NewExamplePrompt(promptExamples)
	ex.Show()
}

func RepeatPromptExample() {
	ex := eggy.NewExamplePrompt(promptExamples).
		Repeat(true)
	ex.Show()
}
