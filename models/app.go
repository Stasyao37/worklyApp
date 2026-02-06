package models

import (
	tea "github.com/charmbracelet/bubbletea"
)

type AppModel struct {
	tabs      []TabModel
	activeTab int
	width     int
	height    int
}

type TabModel interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (TabModel, tea.Cmd)
	View() string
	Title() string
	Key() string
}

func newApp() *AppModel {
	return &AppModel{
		tabs: []TabModel{
			NewWeatherTab(),
			NewNotesTab(),
			NewTodoTab(),
			NewTrafficTab(),
		},
		activeTab: 0,
	}
}

func (m *AppModel) Init() tea.Cmd {
	var cmds []tea.Cmd
	for _, tab := range m.tabs {
		cmds = append(cmds, tab.Init())
	}
	return tea.Batch(cmds...)
}

func (m *AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	}
}
