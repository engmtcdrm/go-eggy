package eggy

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewExamplePrompt(t *testing.T) {
	t.Run("with examples", func(t *testing.T) {
		examples := []Example{
			{Name: "Test1", Fn: func() {}},
			{Name: "Test2", Fn: func() {}},
		}

		ep := NewExamplePrompt(examples)

		require.NotNil(t, ep, "expected ExamplePrompt to be non-nil")
		require.Len(t, ep.examples, 2, "expected 2 examples in the ExamplePrompt")
		require.Equal(t, "Test1", ep.examples[0].Name, "expected first example name to be 'Test1'")
		require.Equal(t, "Test2", ep.examples[1].Name, "expected second example name to be 'Test2'")
	})

	t.Run("with nil examples", func(t *testing.T) {
		ep := NewExamplePrompt(nil)

		require.NotNil(t, ep, "expected ExamplePrompt to be non-nil")
		require.NotNil(t, ep.examples, "expected examples to be initialized to empty slice")
		require.Len(t, ep.examples, 0, "expected 0 examples in the ExamplePrompt")
	})

	t.Run("with empty examples", func(t *testing.T) {
		examples := []Example{}

		ep := NewExamplePrompt(examples)

		require.NotNil(t, ep, "expected ExamplePrompt to be non-nil")
		require.Len(t, ep.examples, 0, "expected 0 examples in the ExamplePrompt")
	})
}

func TestExamplePrompt_Title(t *testing.T) {
	t.Run("set title", func(t *testing.T) {
		ep := NewExamplePrompt(nil)

		result := ep.Title("My Examples")

		require.Equal(t, ep, result, "expected Title to return the same ExamplePrompt instance")
		require.Equal(t, "My Examples", ep.title, "expected title to be 'My Examples'")
	})

	t.Run("set empty title", func(t *testing.T) {
		ep := NewExamplePrompt(nil)
		ep.Title("Initial Title")
		ep.Title("")

		require.Equal(t, "", ep.title, "expected title to be empty")
	})

	t.Run("method chaining", func(t *testing.T) {
		ep := NewExamplePrompt(nil).
			Title("Chained Title")

		require.Equal(t, "Chained Title", ep.title, "expected title to be 'Chained Title'")
	})
}

func TestExamplePrompt_Repeat(t *testing.T) {
	t.Run("set repeat true", func(t *testing.T) {
		ep := NewExamplePrompt(nil)

		result := ep.Repeat(true)

		require.Equal(t, ep, result, "expected Repeat to return the same ExamplePrompt instance")
		require.True(t, ep.repeat, "expected repeat to be true")
	})

	t.Run("set repeat false", func(t *testing.T) {
		ep := NewExamplePrompt(nil)
		ep.Repeat(true)
		ep.Repeat(false)

		require.False(t, ep.repeat, "expected repeat to be false")
	})

	t.Run("default repeat is false", func(t *testing.T) {
		ep := NewExamplePrompt(nil)

		require.False(t, ep.repeat, "expected default repeat to be false")
	})

	t.Run("method chaining", func(t *testing.T) {
		ep := NewExamplePrompt(nil).
			Repeat(true)

		require.True(t, ep.repeat, "expected repeat to be true after chaining")
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

		require.Equal(t, "Test Title", ep.title, "expected title to be 'Test Title'")
		require.True(t, ep.repeat, "expected repeat to be true")
		require.Len(t, ep.examples, 1, "expected 1 example in the ExamplePrompt")
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

		require.Equal(t, "Test Example", example.Name, "expected Name to be 'Test Example'")

		example.Fn()
		require.True(t, called, "expected Fn to be called")
	})

	t.Run("nil function", func(t *testing.T) {
		example := Example{
			Name: "Nil Fn Example",
			Fn:   nil,
		}

		require.Equal(t, "Nil Fn Example", example.Name, "expected Name to be 'Nil Fn Example'")
		require.Nil(t, example.Fn, "expected Fn to be nil")
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
	require.Equal(t, expected, output, "expected output to indicate no examples available")
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
	require.Equal(t, expected, output, "expected output to include title and indicate no examples available")
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
