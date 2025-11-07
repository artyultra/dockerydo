# dockerydo

A lightweight Terminal UI (TUI) for managing Docker resources on Linux and macOS, built as a simple alternative to Docker Desktop for basic Docker operations.

## Features

- **Container Management**: View, start, stop, restart, and remove containers
- **Image Management**: List and manage Docker images
- **Volume Management**: View and manage Docker volumes
- **Network Management**: Inspect and manage Docker networks
- **Interactive UI**: Clean terminal interface built with Bubble Tea
- **Multiple View Modes**: Switch between table view and detailed view
- **Real-time Logs**: View container logs directly in the TUI
- **Theme Support**: Customizable color themes
- **Lightweight**: No heavy desktop application required

## Requirements

- Go 1.24.1 or higher
- Docker installed and running
- Linux or macOS

## Installation

### From Source

```bash
git clone https://github.com/yourusername/dockerydo.git
cd dockerydo
go build
```

The binary will be created as `dockerydo` in the current directory.

## Usage

Simply run the binary:

```bash
./dockerydo
```

### Navigation

- **Tab**: Switch between tabs (Containers, Images, Volumes, Networks)
- **Arrow Keys**: Navigate through lists
- **Enter**: Select item for detailed view
- **q**: Quit application

## Project Structure

```
dockerydo/
├── main.go                 # Application entry point
├── internal/
│   ├── app/               # Application logic and update handlers
│   ├── docker/            # Docker operations (containers, images, volumes, networks)
│   ├── types/             # Type definitions and models
│   ├── ui/                # UI rendering and layout
│   ├── theme/             # Theme and color definitions
│   └── json/              # JSON handling utilities
```

## Development

### Building

```bash
go build
```

### Running

```bash
go run main.go
```

## Technologies Used

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Style definitions for terminal output
- [Bubbles](https://github.com/charmbracelet/bubbles) - TUI components (tables, viewports)

## License

[Add your license here]

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Roadmap

- [ ] Additional container operations
- [ ] Docker Compose support
- [ ] Performance metrics and monitoring
- [ ] Export capabilities
- [ ] Configuration file support
