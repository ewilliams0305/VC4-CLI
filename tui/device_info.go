package tui

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	vc "github.com/ewilliams0305/VC4-CLI/vc"
)

type DeviceTableModel struct {
	Table table.Model
	Help  HelpModel
	row   string
}

func NewDeviceInfo() DeviceTableModel {
	return DeviceTableModel{}
}

func (m DeviceTableModel) Init() tea.Cmd {
	return DeviceInfoCommand
}

func (m DeviceTableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd
	switch msg := msg.(type) {

	case vc.DeviceInfo:
		return NewDeviceTable(msg), nil

	case error:
		return NewDeviceErrorTable(msg), nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return InitialModel(), DeviceInfoCommand
		case "down":
			m.Table.SetCursor(m.Table.Cursor() + 1)
		case "up":
			m.Table.SetCursor(m.Table.Cursor() - 1)

		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.Table.SelectedRow()[1]),
			)
		}
	}
	m.row = m.Table.SelectedRow()[0] + ": " + m.Table.SelectedRow()[1]
	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}

func (m DeviceTableModel) View() string {
	s := DisplayLogo()
	s += BaseStyle.Render(m.Table.View()) + "\n\n"

	if len(m.row) > 0 {
		s += SelectText.Render(m.row)
	}

	s += m.Help.renderHelpInfo()

	return s
}

func DefaultStyles() table.Styles {
	return table.Styles{
		Selected: lipgloss.NewStyle().Padding(0, 1),
		Header:   lipgloss.NewStyle().Bold(true).Padding(0, 1),
		Cell:     lipgloss.NewStyle().Padding(0, 1),
	}
}

func HomeDeviceInfo(info vc.DeviceInfo) DeviceTableModel {
	columns := []table.Column{
		{Title: "SERVER INFOMATION", Width: 20},
		{Title: "", Width: 50},
	}

	rows := []table.Row{
		{"Hostname", info.Name},
		{"MAC Address", info.MACAddress},
		{"Build Date", info.BuildDate},
		{"App Version", info.ApplicationVersion},
		{"Firmware", info.Version},
		{"Mono Version", info.MonoVersion},
		{"Python Version", info.PythonVersion},
		{"Manufacturer", info.Manufacturer},
		{"Model", info.Model},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(false),
		table.WithStyles(DefaultStyles()),
		table.WithHeight(9),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(true)
	// s.Selected = s.Selected.
	// 	Foreground(lipgloss.Color("229")).
	// 	Background(lipgloss.Color("57")).
	// 	Bold(false)
	t.SetStyles(s)

	t.Blur()

	return DeviceTableModel{
		Table: t,
		Help:  NewHelpModel()}
}

func NewDeviceTable(info vc.DeviceInfo) DeviceTableModel {
	columns := []table.Column{
		{Title: "SERVER INFOMATION", Width: 20},
		{Title: "", Width: 50},
	}

	rows := []table.Row{
		{"Hostname", info.Name},
		{"MAC Address", info.MACAddress},
		{"Build Date", info.BuildDate},
		{"App Version", info.ApplicationVersion},
		{"Firmware", info.Version},
		{"Mono Version", info.MonoVersion},
		{"Python Version", info.PythonVersion},
		{"Manufacturer", info.Manufacturer},
		{"Model", info.Model},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(false),
		table.WithStyles(DefaultStyles()),
		table.WithHeight(9),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	t.Blur()

	return DeviceTableModel{
		Table: t,
		Help:  NewHelpModel()}
}

func NewDeviceErrorTable(msg vc.VirtualControlError) DeviceTableModel {
	columns := []table.Column{
		{Title: "SERVER ERROR", Width: 20},
		{Title: "", Width: 100},
	}

	rows := []table.Row{
		{"ERROR", msg.Error()},
		{"", ""},
		{"MESSAGE", "There was an error connecting to the VC4 service"},
		{"", "Please verify your IP address and token"},
		{"", "Please veriify the virtualcontrol service is enabled and running."},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("#FF0000")).
		Bold(true).Italic(true)
	t.SetStyles(s)

	return DeviceTableModel{Table: t, Help: NewHelpModel()}
}

func DeviceInfoCommand() tea.Msg {

	info, err := server.DeviceInfo()
	if err != nil {
		return err
	}
	return info
}
