# eggy 🥚

**E**xempli **G**ratia **G**o **Y**ielder

A lightweight Go framework for creating interactive example runners for your packages. Eggy makes it easy to showcase how your code works through selectable, executable examples.

## What is Eggy?

Eggy is designed to help package authors create interactive demos and examples for their libraries. Instead of writing separate example files or expecting users to read through documentation, eggy provides an interactive menu-driven interface where users can select and run examples on the fly.

## Installation

```bash
go get github.com/engmtcdrm/go-eggy
```

## Quick Start

```go
package main

import (
    "fmt"

    "github.com/engmtcdrm/go-eggy"
)

func main() {
    examples := []eggy.Example{
        {Name: "Hello World Example", Fn: helloWorldExample},
        {Name: "Hi Mom Example", Fn: hiMomExample},
    }

    eggy.NewExamplePrompt(examples).
        Title("My Package Examples").
        Show()
}

func helloWorldExample() {
    fmt.Println("Hello World!")
}

func hiMomExample() {
    fmt.Println("Hi Mom!")
}
```

## Features

### Basic Example Prompt

Create a simple example selector:

```go
examples := []eggy.Example{
    {Name: "Example One", Fn: exampleOne},
    {Name: "Example Two", Fn: exampleTwo},
}

eggy.NewExamplePrompt(examples).Show()
```

### Adding a Title

Display a title above the example menu:

```go
eggy.NewExamplePrompt(examples).
    Title("My Awesome Examples").
    Show()
```

### Repeat Mode

Allow users to run multiple examples in sequence without restarting:

```go
eggy.NewExamplePrompt(examples).
    Repeat(true).
    Show()
```

When `Repeat(true)` is set, after running an example, the user will be prompted to run another example until they choose to exit.

### Nested Examples

You can create nested example menus for organizing complex demonstrations:

```go
var topLevelExamples = []eggy.Example{
    {Name: "Basic Examples", Fn: showBasicExamples},
    {Name: "Advanced Examples", Fn: showAdvancedExamples},
}

var basicExamples = []eggy.Example{
    {Name: "Hello World", Fn: helloWorld},
    {Name: "Simple Math", Fn: simpleMath},
}

func showBasicExamples() {
    eggy.NewExamplePrompt(basicExamples).
        Repeat(true).
        Show()
}
```

## API Reference

### Types

#### `Example`

Represents a single example with a name and function to execute.

```go
type Example struct {
    Name string // Display name in the selection menu
    Fn   func() // Function to execute when selected
}
```

#### `ExamplePrompt`

Manages a collection of examples and handles user interaction.

### Functions

#### `NewExamplePrompt(examples []Example) *ExamplePrompt`

Creates a new ExamplePrompt with the provided examples.

#### `(*ExamplePrompt) Title(title string) *ExamplePrompt`

Sets a title to display above the example menu. Returns the ExamplePrompt for method chaining.

#### `(*ExamplePrompt) Repeat(repeat bool) *ExamplePrompt`

Sets whether to prompt the user to run additional examples after each execution. Returns the ExamplePrompt for method chaining.

#### `(*ExamplePrompt) Show()`

Displays the example prompt and handles user interaction.

## Customization

Eggy provides default styling functions that you can override:

```go
// Style the selected answer
eggy.DefaultAnswerFunc = func(answer string) string {
    return myCustomStyle(answer)
}

// Style the cursor
eggy.DefaultCursorFunc = func(cursor string) string {
    return myCustomStyle(cursor)
}

// Style the selected option
eggy.DefaultSelectFunc = func(s string) string {
    return myCustomStyle(s)
}

// Style icons
eggy.DefaultIconFunc = func(icon string) string {
    return myCustomStyle(icon)
}
```

## Example Project Structure

For a complete working example of eggy in action, see the [example](./example) directory. This is a great way to see how eggy works and get ideas for how to structure your own example runners.

```
example/
├── main.go              # Entry point
├── examples/
│   └── all_examples.go  # Example definitions
└── go.mod
```

To run the examples:

```bash
cd example
go run .
```

## License

See [LICENSE](./LICENSE) for details.
