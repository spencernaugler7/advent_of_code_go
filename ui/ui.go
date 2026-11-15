package ui

import (
	"io"
	"maps"
	"slices"
	"strings"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var modelStyle = lipgloss.NewStyle().
	Align(lipgloss.Center).
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("69"))

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type Model struct {
	funcMap map[string]func(writer io.Writer)
	choices list.Model
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
		case "enter":
			i, ok := m.choices.SelectedItem().(SandboxItem)
			if ok {
				choice := i.value
				writer := WriteCollector{}
				functoRun := m.funcMap[choice]
				functoRun(&writer)
				m.output = writer.Output
			}
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.choices.SetSize(msg.Width-h, msg.Height-v)
	}
	var cmd tea.Cmd
	m.choices, cmd = m.choices.Update(msg)
	return m, cmd
}

func (m Model) View() tea.View {
	var s strings.Builder

	s.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, modelStyle.Render(m.choices.View()), modelStyle.Render(m.output)))
	v := tea.NewView(s.String())
	v.AltScreen = true
	return v
}

func New(registrar *map[string]func(io.Writer)) *tea.Program {
	keys := slices.Collect(maps.Keys(*registrar))

	var sandboxitems []list.Item
	for _, val := range keys {
		item := SandboxItem{value: val}
		sandboxitems = append(sandboxitems, item)
	}

	listModel := list.New(sandboxitems, list.NewDefaultDelegate(), 0, 0)
	listModel.Title = "Run go experiments"
	model := Model{
		funcMap: *registrar,
		choices: listModel,
	}
	return tea.NewProgram(model)
}
