sf-tui

`sf-tui` is a terminal user interface (TUI) for the Salesforce CLI (`sf`), built with Go and the [Bubble Tea](https://github.com/charmbracelet/bubbletea) framework. It aims to provide a more interactive and visual way to manage Salesforce environments and metadata directly from the command line.

## Features

### 1. Org Picker
Easily switch between your authenticated Salesforce orgs.
- Lists all authenticated orgs (Scratch Orgs, Sandboxes, Dev Hubs, etc.).
- Displays aliases, usernames, and instance URLs.
- Visual indicator for the current default org.
- Quick search and selection to update your global `target-org`.

### 2. Metadata Generator
Streamline the creation of new Salesforce metadata.

## Prerequisites

- [Go](https://go.dev/doc/install) (1.19 or later)
- [Salesforce CLI (sf)](https://developer.salesforce.com/tools/sfdxcli) installed and authenticated with at least one org.

## Installation

Clone the repository and build the binary:

```bash
git clone https://github.com/brtheo/sf-tui.git
cd sf-tui
go build -o sf-tui
```

## Keybindings

- **Arrows / J / K**: Navigate lists.
- **Enter**: Confirm selection / Submit input.
- **Esc / Ctrl+C**: Quit the application.


## Roadmap

- [ ] Add support for Apex Class and Trigger generation.
- [ ] Make output directory configurable via flags.
- [ ] Improve error handling and user feedback during CLI execution.
- [ ] Add more `sf` command integrations (e.g., deployments, log viewing).

## License

[MIT](LICENSE)
