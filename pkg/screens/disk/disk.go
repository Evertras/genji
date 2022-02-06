package disk

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dustin/go-humanize"
)

const (
	padding  = 2
	maxWidth = 80
)

type Model struct {
	diskStats *diskStats
	progress  progress.Model
}

func New() Model {
	m := Model{
		diskStats: nil,
		progress:  progress.New(progress.WithGradient("#FF0000", "#0000FF")),
	}

	m.progress.ShowPercentage = true

	return m
}

func (m Model) Init() tea.Cmd {
	return func() tea.Msg {
		stats, err := diskStatsWd()

		if err != nil {
			return errMsg{err}
		}

		return stats
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case *diskStats:
		m.diskStats = msg

	case tea.KeyMsg:
		switch msg.String() {
		// Quit
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
	}

	return m, nil
}

func (m Model) View() string {
	var s string
	pad := strings.Repeat(" ", padding)

	if m.diskStats != nil {
		s = pad + m.progress.ViewAs(m.diskStats.percentFree) + "\n"
		s += pad + fmt.Sprintf("Disk stats: %s / %s free\n", humanize.IBytes(m.diskStats.free), humanize.IBytes(m.diskStats.total))
	} else {
		s = pad + "Reading disk...\n"
	}

	s += "\n"

	return s
}
