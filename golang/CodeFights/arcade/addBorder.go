package arcade

import "strings"

// Given a rectangular matrix of characters, add a border of asterisks(*) to it.

func addBorder(p []string) []string {
	var result []string

	ast := "*"

	// get the length of the first item and add 2 who mean asterisks at front and end
	l := len(p[0]) + 2

	// we need to have the first element to be the list of asterisks
	result = append(result, strings.Repeat(ast, l))

	for _, h := range p {
		result = append(result, ast+h+ast)
	}

	// we need to have the last element to be the list of asterisks
	result = append(result, strings.Repeat(ast, l))

	return result
}
