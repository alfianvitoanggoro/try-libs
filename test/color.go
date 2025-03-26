package test

import "github.com/fatih/color"

type Color struct {
	message string
}

func NewDebug(message string) *Color {
	return &Color{
		message: message,
	}
}

func (d *Color) Success() {
	// A newline will be appended automatically
	color.Green("Prints %s in green.", d.message)
}

func (d *Color) Error() {
	// A newline will be appended automatically
	color.Red("Prints %s in red.", d.message)
}

func (c *Color) Color() {
	// Success
	c.Success()

	// Error
	c.Error()
}
