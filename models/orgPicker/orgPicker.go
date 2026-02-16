package orgPicker

import (
	"fmt"
	"os/exec"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)
var ORG_LIST = []string{"org", "list", "--json"}
var SET_ORG = []string{"config", "set","target-org"}

type fetchOrgsMsg struct {
	orgs map[string]OrgDescriptor
	err  error
}

type viewState string
const (
	IDLE viewState = "IDLE"
	DONE viewState = "DONE"
)

type Model struct {
	list list.Model
	spinner spinner.Model
	state viewState
	currentOrgAlias string
	orgDescriptors    map[string]OrgDescriptor
	loading bool
}

type orgItem struct {
	alias, username, instanceUrl string
	isDefaultOrg bool
}
func (i orgItem) Title() string       { return i.alias }
func (i orgItem) Description() string { return i.username }
func (i orgItem) FilterValue() string { return i.alias }

func fillOrgs(orgResult Result) map[string]OrgDescriptor {
	orgs := make(map[string]OrgDescriptor)
	for _, org := range orgResult.Other { orgs[org.Alias] = org }
	for _, org := range orgResult.ScratchOrgs { orgs[org.Alias] = org }
	for _, org := range orgResult.NonScratchOrgs { orgs[org.Alias] = org }
	for _, org := range orgResult.DevHubs {	orgs[org.Alias] = org }
	for _, org := range orgResult.Sandboxes { orgs[org.Alias] = org }

	return orgs
}

func findDefaultOrgAlias(orgs map[string]OrgDescriptor) string {
	for _, org := range orgs {
		if org.IsDefaultUsername {
			return org.Alias
		}
	}
	return ""
}
func toggleCheckboxes(list list.Model) {
	index := list.Index()
	items := list.Items()
  for k, item := range items {
    if o, ok := item.(orgItem); ok {
      o.isDefaultOrg = k == index
      items[k] = o
    }
  }
  list.SetItems(items)
}

func (m *Model) handleFetchOrgsMsg(msg fetchOrgsMsg) {
	m.orgDescriptors = msg.orgs
	m.currentOrgAlias = findDefaultOrgAlias(msg.orgs)
	m.loading = false

	var orgItems = []list.Item{}
	for _, org := range msg.orgs {
		orgItems = append(orgItems,
			orgItem{
				alias:        org.Alias,
				isDefaultOrg: org.Alias == m.currentOrgAlias,
				username:     org.Username,
				instanceUrl:  org.InstanceURL,
			},
		)
	}
	m.loading = false
	m.list.SetItems(orgItems)
}

func New() Model {
	orgItemList := list.New([]list.Item{}, newOrgItemDelegate(), 0, 0)
	orgItemList.SetShowFilter(false)
	orgItemList.SetShowStatusBar(false)
	orgItemList.SetShowPagination(false)
	orgItemList.SetShowTitle(false)
	orgItemList.SetShowHelp(false)

	loader := spinner.New()
  loader.Spinner = spinner.MiniDot
  loader.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
  loader.Style.Align(lipgloss.Top)


	return Model{
		list: orgItemList,
		state: IDLE,
		spinner: loader,
		loading: true,
		orgDescriptors:  make(map[string]OrgDescriptor),
		currentOrgAlias: "",
	}
}

func fetchOrgsCmd() tea.Cmd {
	return func() tea.Msg {
		raw, err := exec.Command("sf", ORG_LIST...).Output()
		if err != nil {
			return fetchOrgsMsg{err: err}
		}

		orgs, err := UnmarshalOrgs(raw)
		if err != nil {
			return fetchOrgsMsg{err: err}
		}

		orgDescriptors := fillOrgs(orgs.Result)
		return fetchOrgsMsg{orgs: orgDescriptors, err: nil}
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, fetchOrgsCmd(),)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case fetchOrgsMsg:
			if msg.err != nil {
				fmt.Println("Error fetching orgs:", msg.err)
				return m, tea.Quit
			}
			m.handleFetchOrgsMsg(msg)
		case tea.KeyMsg:
			switch msg.Type {
				case tea.KeyEnter:
					m.currentOrgAlias = m.list.SelectedItem().(orgItem).alias
					toggleCheckboxes(m.list)
					m.state = DONE
					return m, func() tea.Cmd {
						return func() tea.Msg {
							SET_ORG = append(SET_ORG, m.currentOrgAlias)
							err := exec.Command("sf", SET_ORG...).Run()
							if err != nil {
								fmt.Println("Error setting org:", err)
							}
							return tea.Quit
						}
					}()
				case tea.KeyCtrlC, tea.KeyEsc:
					return m, tea.Quit
			}
			case tea.WindowSizeMsg:
				m.list.SetSize(msg.Width - 2, msg.Height - 2)
	}

	var listCmd, spinnerCmd tea.Cmd
	m.list, listCmd = m.list.Update(msg)
	m.spinner, spinnerCmd = m.spinner.Update(msg)
	return m, tea.Batch(listCmd, spinnerCmd)
}

func (m Model) View() string {
	switch m.state {
	case IDLE:
		if m.loading {
			return m.spinner.View()
		}
		return m.list.View()
	case DONE:
		return fmt.Sprintf("Default ☁️ org is now : %s", m.currentOrgAlias)
	}
	return ""
}
