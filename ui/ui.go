package ui

import (
	"io"
	"maps"
	"slices"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func New(registrar *map[string]func(io.Writer)) *tea.Program {
	keys := slices.Collect(maps.Keys(*registrar))
	model := Model{cursor: 0, funcMap: *registrar, choices: keys}
	return tea.NewProgram(model)
}

type Model struct {
	funcMap map[string]func(writer io.Writer)
	choices []string
	cursor  int
	output  string
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", "space":
			key := m.choices[m.cursor]
			selectedFunc, found := m.funcMap[key]
			if !found {
				return m, nil
			}
			writeCollector := WriteCollector{}
			selectedFunc(&writeCollector)
			m.output = writeCollector.Output
		}
	}
	return m, nil
}

func (m Model) View() tea.View {
	var s strings.Builder
	s.WriteString("Choose Experiment to run\n\n")
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s.WriteString(cursor)
		s.WriteString(choice)
		s.WriteString("\n")
	}
	s.WriteString("\n\nPress q to quit.\n")
	s.WriteString("\n")
	s.WriteString(m.output)
	return tea.NewView(s.String())
}

type WriteCollector struct {
	Output string
}

func (c *WriteCollector) Write(p []byte) (n int, err error) {
	c.Output += strings.Clone(string(p))
	return len(p), nil
}
