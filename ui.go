package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var titleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#d7dae1")).
	PaddingTop(1)

var itemStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#7FAD8A")).
	PaddingLeft(2)

func PrintTitle(s string) {
	fmt.Println(titleStyle.Render(s))
}

func Print(s string) {
	fmt.Println(itemStyle.Render(s))
}
