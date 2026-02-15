# sf-tui

`sf-tui` is a terminal user interface (TUI) for the Salesforce CLI (`sf`), built with Go and the [Bubble Tea](https://github.com/charmbracelet/bubbletea) framework. It aims to provide a more interactive and visual way to manage Salesforce environments and metadata directly from the command line.

## Disclaimer

⚠️ This Readme is AI generated and may contain inaccuracies or errors ⚠️

## Features

### 1. Org Picker
Easily switch between your authenticated Salesforce orgs.
- Lists all authenticated orgs (Scratch Orgs, Sandboxes, Dev Hubs, etc.).
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

### 3. Metadata Retriever
Browse, search, and select metadata from your Salesforce org to generate deployment packages.
- **Two-step wizard interface**: First select a metadata type, then browse all available metadata of that type.
- **Interactive metadata type picker**: Easily navigate and select from all available metadata types.
- **Searchable metadata table**: Filter metadata by name or other properties in real-time.
- **Metadata selection**: Check/uncheck individual metadata items for inclusion in your deployment package.
- **Package.xml generation**: Automatically generate a valid `package.xml` file with your selected metadata.
- **Live filtering**: Search and filter metadata as you type.
- **Real-time status**: Animated spinner shows when metadata is being fetched from your org.

## Prerequisites

- [Go](https://go.dev/doc/install) (1.19 or later)
- [Salesforce CLI (sf)](https://developer.salesforce.com/tools/sfdxcli) installed and authenticated with at least one org.

## Installation

Clone the repository and build the binary:

```bash
git clone https://github.com/brtheo/sf-tui.git
cd sf-tui
go build -o sf-tui
