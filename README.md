# sf-tui

`sf-tui` is a terminal user interface (TUI) for the Salesforce CLI (`sf`), built with Go and the [Bubble Tea](https://github.com/charmbracelet/bubbletea) framework. It aims to provide a more interactive and visual way to manage Salesforce environments and metadata directly from the command line.

## Disclaimer

⚠️ This Readme is AI generated and may contain inaccuracies or errors ⚠️

## Features

### 1. Org Picker
Easily switch between your authenticated Salesforce orgs.
- Lists all authenticated orgs (Scratch Orgs, Sandboxes, Dev Hubs, Sandboxes, etc.).
- Displays aliases, usernames, and instance URLs.
- Visual indicator for the current default org.
- Quick selection to update your global `target-org` with immediate confirmation feedback.

### 2. Metadata Generator
Streamline the creation of new Salesforce metadata with a simple, interactive form.
- **Lightning Web Components (LWC)**: Generate new LWC components.
- **Apex Classes**: Create new Apex classes.
- **Apex Triggers**: Create new Apex triggers with SObject selection.
- Input validation to ensure metadata names are not empty.
- Real-time feedback on successful creation.

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

## Usage

### Org Picker (Default)
Launch the org picker to switch between authenticated orgs:

```bash
./sf-tui
# or explicitly:
./sf-tui org-picker
```

Navigate using arrow keys and press **Enter** to set the selected org as the default target-org.

### Metadata Generator
Generate new Salesforce metadata with specified type and output path:

```bash
./sf-tui gen --type LWC --output /path/to/force-app
./sf-tui gen --type ApexClass --output /path/to/force-app
./sf-tui gen --type ApexTrigger --output /path/to/force-app
```

**Supported Metadata Types:**
- `LWC`: Lightning Web Components
- `ApexClass`: Apex Classes
- `ApexTrigger`: Apex Triggers

The generator will prompt you for a name and create the metadata in the specified output directory.

## Keybindings

- **↑ / ↓ (Arrows) / J / K**: Navigate lists and options.
- **Enter**: Confirm selection or submit input.
- **Esc / Ctrl+C**: Quit the application.

## Project Structure

```
sf-tui/
├── main.go              # Entry point with CLI argument parsing
├── models/
│   ├── orgPicker/       # Org selection and switching logic
│   │   ├── orgPicker.go
│   │   ├── orgItemDelegate.go
│   │   └── orgs.go
│   └── genMetadata/     # Metadata generation logic
│       └── genMetadata.go
└── README.md
```

## Roadmap

- [ ] Support for additional metadata types (custom objects, fields, etc.).
- [ ] Batch metadata creation.
- [ ] Integration with org deployments and log viewing.
- [ ] Enhanced error handling with detailed user feedback.
- [ ] Configuration file support for default output paths.

## License

[MIT]
