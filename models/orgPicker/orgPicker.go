package orgPicker
import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/list"
	"fmt"
	"os/exec"
	"time"
)
var ORG_LIST = []string{"org", "list", "--json"}
var SET_ORG = []string{"config", "set","target-org"}

type viewState string
type tickMsg time.Time
type setStateMsg string
const (
	IDLE viewState = "IDLE"
	DONE viewState = "DONE"
)

type Model struct {
	list list.Model
	state viewState
	currentOrgAlias string
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

func New() Model {
	raw, err := exec.Command("sf",ORG_LIST...).Output()
	if err != nil {
		fmt.Println(err)
	}
	orgs, err := UnmarshalOrgs(raw)
	if err != nil {
		fmt.Println(err)
	}

	orgDescriptors := fillOrgs(orgs.Result)
	currentOrgAlias := findDefaultOrgAlias(orgDescriptors)
	var orgItems = []list.Item{}
	for _, org := range orgDescriptors {
		orgItems = append(orgItems,
			orgItem {
				alias: org.Alias,
				isDefaultOrg: org.Alias == currentOrgAlias,
				username: org.Username,
				instanceUrl: org.InstanceURL,
			},
		)
	}
	orgItemList := list.New(orgItems, newOrgItemDelegate(), 0, 0)
	orgItemList.SetShowFilter(false)
	orgItemList.SetShowStatusBar(false)
	orgItemList.SetShowPagination(false)
	orgItemList.SetShowTitle(false)
	orgItemList.SetShowHelp(false)
	return Model{
		list: orgItemList,
		state: IDLE,
		currentOrgAlias: currentOrgAlias,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
				case tea.KeyEnter:
					m.currentOrgAlias = m.list.SelectedItem().(orgItem).alias
					toggleCheckboxes(m.list)
					return m, func() tea.Cmd {
						return func() tea.Msg {
							SET_ORG = append(SET_ORG, m.currentOrgAlias)
							err := exec.Command("sf", SET_ORG...).Run()
							if err != nil {
								fmt.Println("Error setting org:", err)
							}
							m.state = DONE
							return setStateMsg("DONE")
						}
					}()
				case tea.KeyCtrlC, tea.KeyEsc:
					return m, tea.Quit
			}
			case setStateMsg:
				return m, tea.Tick(time.Second, func(t time.Time) tea.Msg {
           return tickMsg(t)
         })
			case tickMsg:
        return m, tea.Quit
			case tea.WindowSizeMsg:
				m.list.SetSize(msg.Width - 1, msg.Height - 1)
	}
	var listCmd tea.Cmd
	m.list, listCmd = m.list.Update(msg)
	return m, listCmd
}
func (m Model) View() string {
	switch m.state {
	case IDLE:
		return m.list.View()
	case DONE:
		return fmt.Sprintf("Default ☁️ org is now : %s", m.currentOrgAlias)
	}
	return ""
}
