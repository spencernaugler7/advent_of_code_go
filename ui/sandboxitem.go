package ui

type SandboxItem struct {
	value string
}

func (s SandboxItem) FilterValue() string { return s.value }

func (s SandboxItem) Title() string { return s.value }

func (s SandboxItem) Description() string { return s.value }
