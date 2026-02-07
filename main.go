package main

import (
	"os"
	"strconv"

	orgPicker "github.com/brtheo/sf-tui/models/orgPicker"
	newMetadata "github.com/brtheo/sf-tui/models/newMetadata"
	tea "github.com/charmbracelet/bubbletea"
)

func getArg(index int, fallback string) string {
	if len(os.Args) > 1 {
		return os.Args[index]
	}
	return fallback
}
func getIntArg(index int, fallback int) int {
	arg, err := strconv.Atoi(getArg(index, strconv.Itoa(fallback)))
	if err != nil {panic(err)}
	return arg
}
func getModelArg(index int, fallback ModelType) ModelType {
	return ModelType(getIntArg(index, int(fallback)))
}

type ModelType int

const (
	NewMetadata ModelType = iota
	OrgPicker
)

func main() {
	var model tea.Model
	switch getModelArg(1, NewMetadata) {
		case NewMetadata:
			model = newMetadata.New("LWC")
		case OrgPicker:
			model = orgPicker.New()


	}
tea.NewProgram(model, tea.WithAltScreen()).Run()
}
