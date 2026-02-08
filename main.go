package main

import (
	"os"

	"github.com/Stasyao37/worklyApp/models"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Инициализируем главную модель
	m := models.NewApp()

	// Запускаем приложение
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		println("Ошибка:", err)
		os.Exit(1)
	}
}
