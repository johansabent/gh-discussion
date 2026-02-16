package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Section struct {
	Name       string
	CategoryID string
}

type Model struct {
	ready           bool
	Owner           string
	Name            string
	width           int
	contentViewport contentViewport
	previewViewport previewViewport
	sections        []Section
	data            []Discussion
	cursor          cursor
	help            help.Model
	keyMap          keyMap
}

type contentViewport struct {
	width  int
	height int
}

type previewViewport struct {
	width  int
	height int
}

type cursor struct {
	currDiscID    int
	currSectionID int
}

type initMsg struct {
	data     []Discussion
	sections []Section
}

func NewModel(data *[]Discussion) Model {
	return Model{
		data:   *data,
		keyMap: DefaultKeyMap(),
		cursor: cursor{
			currSectionID: 0,
			currDiscID:    0,
		},
		help: help.New(),
	}
}

func getTitleWidth(viewportWidth int) int {
	return viewportWidth - usedWidth
}

func (m Model) contentWidth() int {
	return m.contentViewport.width
}

func (m Model) contentHeight() int {
	return m.contentViewport.height
}

func (m Model) previewWidth() int {
	return 50
}

func (m Model) previewHeight() int {
	return m.previewViewport.height
}

func (m *Model) cursorUp() {
	m.cursor.currDiscID = max(m.cursor.currDiscID-1, 0)
}

func (m *Model) cursorDown() {
	newCursor := min(m.cursor.currDiscID+1, len(m.data)-1)
	newCursor = max(newCursor, 0)

	m.cursor.currDiscID = newCursor
}

func (m *Model) prevSection() {
	m.cursor.currSectionID = max(m.cursor.currSectionID-1, 0)
}

func (m *Model) nextSection() {
	newCursor := min(m.cursor.currSectionID+1, len(m.sections)-1)
	newCursor = max(newCursor, 0)

	m.cursor.currSectionID = newCursor
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keyMap.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keyMap.CursorUp):
			m.cursorUp()
		case key.Matches(msg, m.keyMap.CursorDown):
			m.cursorDown()
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.help.Width = msg.Width
		verticalMargins := headerHeight + footerHeight + pagerHeight

		if !m.ready {
			m.contentViewport = contentViewport{
				width:  m.width - 50,
				height: msg.Height - verticalMargins - 1,
			}
			m.previewViewport = previewViewport{
				width:  0,
				height: msg.Height - verticalMargins + 1,
			}
			m.ready = true
		} else {
			m.contentViewport.height = msg.Height - verticalMargins - 1
			m.previewViewport.height = msg.Height - verticalMargins + 1
		}

		return m, nil
	}

	cmds = append(cmds, m.updateSection(msg))
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m *Model) updateSection(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keyMap.CursorUp):
			m.cursorUp()

		case key.Matches(msg, m.keyMap.CursorDown):
			m.cursorDown()

		case key.Matches(msg, m.keyMap.PrevSection):
			m.prevSection()

		case key.Matches(msg, m.keyMap.NextSection):
			m.nextSection()

		}
	}

	return cmd
}

func (m Model) View() string {
	if !m.ready {
		return "Loading..."
	}

	s := strings.Builder{}
	
	// Simple list view of discussions
	s.WriteString("Discussions\n")
	s.WriteString("===========\n\n")
	
	for i, d := range m.data {
		prefix := "  "
		if i == m.cursor.currDiscID {
			prefix = "> "
		}
		s.WriteString(fmt.Sprintf("%s%s\n", prefix, d.Title))
	}
	
	s.WriteString("\n")
	s.WriteString("Use ↑/↓ to navigate, q to quit\n")

	return s.String()
}

func (m Model) tabsView() string {
	return ""
}

func (m Model) sectionView() string {
	s := strings.Builder{}
	for i, d := range m.data {
		prefix := "  "
		if i == m.cursor.currDiscID {
			prefix = "> "
		}
		s.WriteString(fmt.Sprintf("%s%s\n", prefix, d.Title))
	}
	return s.String()
}

func (m Model) helperView() string {
	return "Use ↑/↓ to navigate, q to quit"
}
