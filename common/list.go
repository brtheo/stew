package common

import "github.com/charmbracelet/bubbles/list"

type ListItem struct {
	checked bool

}
func (i ListItem) Title() string       { return i.alias }
func (i ListItem) Description() string { return i.username }
func (i ListItem) FilterValue() string { return i.alias }

func toggleCheckboxes(list list.Model) {
	index := list.Index()
	items := list.Items()
  for k, item := range items {
    if o, ok := item.(list.Item); ok {
      o.checked = k == index
      items[k] = o
    }
  }
  list.SetItems(items)
}
