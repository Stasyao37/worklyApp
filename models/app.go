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
