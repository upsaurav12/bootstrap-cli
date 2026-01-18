package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/common-nighthawk/go-figure"
)

var asciiStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color(primaryBlue)).
	Bold(true)

const (
	stepName = iota
	stepType
	stepRouter
	stepPort
	stepDB
	stepConfirm
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(softBlue))

	labelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(softBlue))

	hintStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(mutedGray))

	boxStyle = lipgloss.NewStyle().
			Padding(3, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(borderGray).
			Width(60)
)

func renderStep(m wizardModel, title, label, body, hint string) string {
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyle.Render(title),
		"",
		labelStyle.Render(label),
		"",
		body,
		"",
		hintStyle.Render(hint),
	)

	box := boxStyle.
		Width(m.width - 4).
		Height(m.height - 2).
		Render(content)

	return box
}

type wizardModel struct {
	step  int
	input ProjectInput

	text textinput.Model
	list list.Model
	quit bool

	width  int
	height int
}

type item string

func (i item) Title() string       { return string(i) }
func (i item) Description() string { return "" }
func (i item) FilterValue() string { return string(i) }

func initialWizardModel() wizardModel {
	return wizardModel{
		step: stepName,
		text: newTextInput(""),
	}
}

const (
	primaryBlue = lipgloss.Color("33")  // bright blue
	softBlue    = lipgloss.Color("39")  // lighter blue
	mutedGray   = lipgloss.Color("241") // hints
	borderGray  = lipgloss.Color("238") // borders
)

func newTextInput(placeholder string) textinput.Model {
	ti := textinput.New()
	ti.Prompt = "› "
	ti.Placeholder = placeholder
	ti.SetValue("") // ← critical
	ti.Focus()
	return ti
}

func (m wizardModel) Init() tea.Cmd {
	return nil
}

func renderHeader() string {
	fig := figure.NewFigure("Bootstrap  CLI", "slant", true)

	ascii := strings.Trim(fig.String(), "\n")

	return asciiStyle.Render(ascii)
}

func (m wizardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "esc":
			m.quit = true
			return m, tea.Quit

		case "enter":
			switch m.step {

			case stepName:
				m.input.Name = m.text.Value()
				m.step = stepType
				m.list = newList("Project type", []string{"rest"})
				return m, nil

			case stepType:
				m.input.Type = m.list.SelectedItem().(item).Title()
				m.step = stepRouter
				m.list = newList("Router", []string{"gin", "chi", "echo"})
				return m, nil

			case stepRouter:
				m.input.Router = m.list.SelectedItem().(item).Title()
				m.step = stepPort
				m.text = newTextInput("")
				return m, nil

			case stepPort:
				port := m.text.Value()
				if port == "" {
					port = "8080"
				}
				if _, err := strconv.Atoi(port); err != nil {
					return m, nil
				}
				m.input.Port = port
				m.step = stepDB
				m.list = newList("Database", []string{"postgres", "mysql", "mongo"})
				return m, nil

			case stepDB:
				m.input.DB = m.list.SelectedItem().(item).Title()
				m.step = stepConfirm
				return m, nil

			case stepConfirm:
				return m, tea.Quit
			}
		}
	}

	var cmd tea.Cmd
	if m.step == stepName || m.step == stepPort {
		m.text, cmd = m.text.Update(msg)
		return m, cmd
	}

	if m.step == stepType || m.step == stepRouter || m.step == stepDB {
		m.list, cmd = m.list.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m wizardModel) View() string {
	if m.quit {
		return ""
	}

	switch m.step {

	case stepName:
		return renderStep(
			m,
			renderHeader()+"\nCreate New Project",
			"Project name",
			m.text.View(),
			"Enter to continue • Esc to quit",
		)

	case stepType:
		return renderStep(
			m,
			"Project Type",
			"Select project type",
			m.list.View(),
			"↑↓ navigate • Enter select • Esc quit",
		)

	case stepRouter:
		return renderStep(
			m,
			"Router",
			"Select router",
			m.list.View(),
			"↑↓ navigate • Enter select • Esc quit",
		)

	case stepPort:
		return renderStep(
			m,
			"Application Port",
			"Port (default: 8080)",
			m.text.View(),
			"Enter to continue • Esc quit",
		)

	case stepDB:
		return renderStep(
			m,
			"Database",
			"Select database",
			m.list.View(),
			"↑↓ navigate • Enter select • Esc quit",
		)

	case stepConfirm:
		summary := fmt.Sprintf(
			"Project:  %s\nType:     %s\nRouter:   %s\nPort:     %s\nDatabase: %s",
			m.input.Name,
			m.input.Type,
			m.input.Router,
			m.input.Port,
			m.input.DB,
		)

		return renderStep(
			m,
			"Confirm Configuration",
			"Review your selections",
			summary,
			"Enter to generate • Esc to cancel",
		)
	}

	return ""
}

func newList(title string, values []string) list.Model {
	items := make([]list.Item, len(values))
	for i, v := range values {
		items[i] = item(v)
	}

	l := list.New(items, list.NewDefaultDelegate(), 20, 10)
	l.Title = title
	return l
}
