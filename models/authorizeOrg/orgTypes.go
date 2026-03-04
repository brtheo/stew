package authorizeOrg

import (
	"github.com/charmbracelet/bubbles/list"
)
type OrgTypeItem struct {
	title string
	description string
	checked bool
}
func (i OrgTypeItem) Title() string       { return i.title }
func (i OrgTypeItem) Description() string { return i.description }
func (i OrgTypeItem) FilterValue() string { return i.title }

var OrgTypes = []list.Item{
	OrgTypeItem{title: "Production", description: "login.salesforce.com"},
	OrgTypeItem{title: "Sandbox", description: "test.salesforce.com"},
	OrgTypeItem{title: "Custom URL", description: "Enter a custom login URL"},
}
