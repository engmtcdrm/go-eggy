package eggy

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestNewExamplePrompt(t *testing.T) {
	t.Run("with examples", func(t *testing.T) {
		examples := []Example{
			{Name: "Test1", Fn: func() {}},
			{Name: "Test2", Fn: func() {}},
		}

		ep := NewExamplePrompt(examples)

		if ep == nil {
			t.Fatal("expected ExamplePrompt to be non-nil")
		}

		if len(ep.examples) != 2 {
			t.Errorf("expected 2 examples, got %d", len(ep.examples))
		}

		if ep.examples[0].Name != "Test1" {
			t.Errorf("expected first example name to be 'Test1', got '%s'", ep.examples[0].Name)
		}

		if ep.examples[1].Name != "Test2" {
			t.Errorf("expected second example name to be 'Test2', got '%s'", ep.examples[1].Name)
		}
	})

	t.Run("with nil examples", func(t *testing.T) {
		ep := NewExamplePrompt(nil)

		if ep == nil {
			t.Fatal("expected ExamplePrompt to be non-nil")
		}

		if ep.examples == nil {
			t.Error("expected examples to be initialized to empty slice, got nil")
		}

		if len(ep.examples) != 0 {
			t.Errorf("expected 0 examples, got %d", len(ep.examples))
		}
	})

	t.Run("with empty examples", func(t *testing.T) {
		examples := []Example{}

		ep := NewExamplePrompt(examples)

		if ep == nil {
			t.Fatal("expected ExamplePrompt to be non-nil")
		}

		if len(ep.examples) != 0 {
			t.Errorf("expected 0 examples, got %d", len(ep.examples))
		}
	})
}

func TestExamplePrompt_Title(t *testing.T) {
	t.Run("set title", func(t *testing.T) {
		ep := NewExamplePrompt(nil)

		result := ep.Title("My Examples")

		if result != ep {
			t.Error("expected Title to return the same ExamplePrompt instance")
		}

		if ep.title != "My Examples" {
			t.Errorf("expected title to be 'My Examples', got '%s'", ep.title)
		}
	})

	t.Run("set empty title", func(t *testing.T) {
		ep := NewExamplePrompt(nil)
		ep.Title("Initial Title")

		ep.Title("")

		if ep.title != "" {
			t.Errorf("expected title to be empty, got '%s'", ep.title)
		}
	})

	t.Run("method chaining", func(t *testing.T) {
		ep := NewExamplePrompt(nil).
			Title("Chained Title")

		if ep.title != "Chained Title" {
			t.Errorf("expected title to be 'Chained Title', got '%s'", ep.title)
		}
	})
}

func TestExamplePrompt_Repeat(t *testing.T) {
	t.Run("set repeat true", func(t *testing.T) {
		ep := NewExamplePrompt(nil)

		result := ep.Repeat(true)

		if result != ep {
			t.Error("expected Repeat to return the same ExamplePrompt instance")
		}

		if !ep.repeat {
			t.Error("expected repeat to be true")
		}
	})

	t.Run("set repeat false", func(t *testing.T) {
		ep := NewExamplePrompt(nil)
		ep.Repeat(true)

		ep.Repeat(false)

		if ep.repeat {
			t.Error("expected repeat to be false")
		}
	})

	t.Run("default repeat is false", func(t *testing.T) {
		ep := NewExamplePrompt(nil)

		if ep.repeat {
			t.Error("expected default repeat to be false")
		}
	})

	t.Run("method chaining", func(t *testing.T) {
		ep := NewExamplePrompt(nil).
			Repeat(true)

		if !ep.repeat {
			t.Error("expected repeat to be true after chaining")
		}
	})
}

func TestExamplePrompt_MethodChaining(t *testing.T) {
	t.Run("chain all methods", func(t *testing.T) {
		examples := []Example{
			{Name: "Example1", Fn: func() {}},
		}

		ep := NewExamplePrompt(examples).
			Title("Test Title").
			Repeat(true)

		if ep.title != "Test Title" {
			t.Errorf("expected title to be 'Test Title', got '%s'", ep.title)
		}

		if !ep.repeat {
			t.Error("expected repeat to be true")
		}

		if len(ep.examples) != 1 {
			t.Errorf("expected 1 example, got %d", len(ep.examples))
		}
	})
}

func TestExample(t *testing.T) {
	t.Run("struct fields", func(t *testing.T) {
		called := false
		fn := func() { called = true }

		example := Example{
			Name: "Test Example",
			Fn:   fn,
		}

		if example.Name != "Test Example" {
			t.Errorf("expected Name to be 'Test Example', got '%s'", example.Name)
		}

		example.Fn()
		if !called {
			t.Error("expected Fn to be called")
		}
	})

	t.Run("nil function", func(t *testing.T) {
		example := Example{
			Name: "Nil Fn Example",
			Fn:   nil,
		}

		if example.Name != "Nil Fn Example" {
			t.Errorf("expected Name to be 'Nil Fn Example', got '%s'", example.Name)
		}

		if example.Fn != nil {
			t.Error("expected Fn to be nil")
		}
	})
}

func TestExamplePrompt_Show_NoExamples(t *testing.T) {
	// Capture stdout to verify the output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	ep := NewExamplePrompt(nil)
	ep.examples = nil // Force nil examples
	ep.Show()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	expected := "No examples available.\n"
	if output != expected {
		t.Errorf("expected output '%s', got '%s'", expected, output)
	}
}

func TestExamplePrompt_Show_WithTitle_NoExamples(t *testing.T) {
	// Capture stdout to verify the output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	ep := NewExamplePrompt(nil)
	ep.examples = nil // Force nil examples
	ep.Title("My Test Title")
	ep.Show()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	expected := "My Test Title\n\nNo examples available.\n"
	if output != expected {
		t.Errorf("expected output '%s', got '%s'", expected, output)
	}
}

func TestDefaultFuncs(t *testing.T) {
	t.Run("DefaultAnswerFunc is set", func(t *testing.T) {
		if DefaultAnswerFunc == nil {
			t.Error("expected DefaultAnswerFunc to be non-nil")
		}

		result := DefaultAnswerFunc("test")
		if result == "" {
			t.Error("expected DefaultAnswerFunc to return a non-empty string")
		}
	})

	t.Run("DefaultCursorFunc is set", func(t *testing.T) {
		if DefaultCursorFunc == nil {
			t.Error("expected DefaultCursorFunc to be non-nil")
		}

		result := DefaultCursorFunc("test")
		if result == "" {
			t.Error("expected DefaultCursorFunc to return a non-empty string")
		}
	})

	t.Run("DefaultSelectFunc is set", func(t *testing.T) {
		if DefaultSelectFunc == nil {
			t.Error("expected DefaultSelectFunc to be non-nil")
		}

		result := DefaultSelectFunc("test")
		if result == "" {
			t.Error("expected DefaultSelectFunc to return a non-empty string")
		}
	})

	t.Run("DefaultIconFunc is set", func(t *testing.T) {
		if DefaultIconFunc == nil {
			t.Error("expected DefaultIconFunc to be non-nil")
		}

		result := DefaultIconFunc("test")
		if result == "" {
			t.Error("expected DefaultIconFunc to return a non-empty string")
		}
	})
}
