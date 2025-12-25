package eggy

import (
	"fmt"

	"github.com/engmtcdrm/go-pardon"
)

// Example represents a single example with a name and an associated function to execute.
type Example struct {
	Name string // Name of the example to be displayed in the selection prompt.
	Fn   func() // Function to execute when the example is selected.
}

// ExamplePrompt manages a collection of examples and handles user interaction.
type ExamplePrompt struct {
	title    string
	examples []Example
	repeat   bool
}

// NewExamplePrompt creates a new ExamplePrompt with the provided examples.
func NewExamplePrompt(examples []Example) *ExamplePrompt {
	if examples == nil {
		examples = make([]Example, 0)
	}

	return &ExamplePrompt{examples: examples}
}

func (ep *ExamplePrompt) Title(title string) *ExamplePrompt {
	ep.title = title
	return ep
}

// Repeat sets whether to repeat the prompt after an example is executed.
func (ep *ExamplePrompt) Repeat(repeat bool) *ExamplePrompt {
	ep.repeat = repeat
	return ep
}

// Show displays the example prompt and handles user interaction.
func (ep *ExamplePrompt) Show() {
	if ep.title != "" {
		fmt.Println(ep.title)
		fmt.Println()
	}

	if ep.examples == nil {
		fmt.Println("No examples available.")
		return
	}

	cont := ep.showExamples()

	if ep.repeat && cont {
		ep.repeatPrompt()
	}
}

// showExamples displays a list of examples and allows the user to select one to run.
func (ep *ExamplePrompt) showExamples() bool {
	funcMap := map[string]func(){}
	names := make([]pardon.Option[string], 0, len(ep.examples))

	// Populate map with available examples and their functions.
	for i, ex := range ep.examples {
		funcMap[ex.Name] = ex.Fn
		names = append(names, pardon.NewOption(fmt.Sprintf("%d. %s", i+1, ex.Name), ex.Name))
	}

	var selectedName string

	selectPrompt := pardon.NewSelect(&selectedName).
		Title("Select an example:").
		Icon("").
		Options(names...).
		SelectFunc(DefaultSelectFunc).
		CursorFunc(DefaultCursorFunc).
		AnswerFunc(DefaultAnswerFunc)

	if err := selectPrompt.Ask(); err != nil {
		fmt.Printf("Error: %v\n", err)
		return false
	}

	fmt.Println()

	// Run the selected example function.
	if fn, ok := funcMap[selectedName]; ok {
		fn()
	} else {
		fmt.Println("No function found for selection.")
	}

	return true
}

// repeatPrompt continues to prompt the user to run more examples until they choose to exit.
func (ep *ExamplePrompt) repeatPrompt() {
	// Keep prompting the user to run examples until they choose to exit.
	cont := true
	for cont {
		fmt.Println()

		contPrompt := pardon.NewConfirm(&cont).
			Title("Do you want to run another example?").
			IconFunc(DefaultIconFunc)

		if err := contPrompt.Ask(); err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		if cont {
			fmt.Println()
			cont = ep.showExamples()
		}
	}
}
