package main

import (
	"fmt"

	"github.com/engmtcdrm/go-ansi"
	"github.com/engmtcdrm/go-pardon"
	pp "github.com/engmtcdrm/go-prettyprint"

	"example.com/examples"
)

func main() {
	pardon.SetDefaultIconFunc(func(icon string) string { return pp.Cyan(icon) })
	pardon.SetDefaultAnswerFunc(func(answer string) string { return pp.Yellow(answer) })
	pardon.SetDefaultCursorFunc(func(cursor string) string { return pp.Yellow(cursor) })
	pardon.SetDefaultSelectFunc(func(s string) string { return pp.Green(s) })

	showExamples()
	repeatPrompt()
}

// showExamples displays a list of examples and allows the user to select one to run.
func showExamples() {
	funcMap := map[string]func(){}
	names := make([]pardon.Option[string], 0, len(examples.AllExamples))

	// Populate map with available examples and their functions.
	for i, ex := range examples.AllExamples {
		funcMap[ex.Name] = ex.Fn
		names = append(names, pardon.NewOption(fmt.Sprintf("%d. %s", i+1, ex.Name), ex.Name))
	}

	var selectedName string

	selectPrompt := pardon.NewSelect(&selectedName).
		Title("Select an example:").
		Icon("").
		Options(names...).
		AnswerFunc(func(s string) string {
			return fmt.Sprintf("%s%s%s", ansi.Yellow, s, ansi.Reset)
		})

	if err := selectPrompt.Ask(); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println()

	// Run the selected example function.
	if fn, ok := funcMap[selectedName]; ok {
		fn()
	} else {
		fmt.Println("No function found for selection.")
	}

	fmt.Println()
}

// Keep prompting the user to run examples until they choose to exit.
func repeatPrompt() {
	cont := true
	for cont {
		fmt.Println()

		contPrompt := pardon.NewConfirm(&cont).
			Title("Do you want to run another example?")

		if err := contPrompt.Ask(); err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		if cont {
			fmt.Println()
			showExamples()
		}
	}
}
