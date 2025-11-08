package main

import (
	"dockerydo/internal/app"
	"dockerydo/internal/types"
	"dockerydo/internal/ui"
	"flag"
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// changes

type bubbletea struct {
	model types.Model
}

func (b bubbletea) Init() tea.Cmd {
	return app.Init(b.model)
}

func (b bubbletea) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	updatedModel, cmd := app.Update(msg, b.model)
	b.model = updatedModel
	return b, cmd
}

func (b bubbletea) View() string {
	return ui.View(b.model)
}

func main() {
	versionFlag := flag.Bool("version", false, "Print version and exit")
	flag.BoolVar(versionFlag, "v", false, "Print version information (shorthand)")
	flag.Parse()

	// check for version flag
	if *versionFlag {
		fmt.Printf("dockerydo %s (commit: %s, built: %s)\n", version, commit, date)
		return
	}
	m := app.NewModel()
	p := tea.NewProgram(bubbletea{model: m})
	_, err := p.Run()

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
