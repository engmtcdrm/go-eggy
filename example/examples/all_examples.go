// Package examples provides a collection of examples for this project
package examples

type Example struct {
	Name string
	Fn   func()
}

var AllExamples = []Example{}
