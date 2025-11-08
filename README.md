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

### Homebrew (macOS/Linux)

```bash
brew install artyultra/dockerydo/dockerydo
```

### Pre-built Binaries

Download the latest release for your platform from the [releases page](https://github.com/artyultra/dockerydo/releases).

Available for:

- Linux (amd64, arm64)
- macOS (amd64, arm64)
- Windows (amd64, arm64)

Extract the archive and move the binary to your PATH:

```bash
# Example for Linux/macOS
tar -xzf dockerydo_*.tar.gz
sudo mv dockerydo /usr/local/bin/
```

### Arch Linux (AUR)

```bash
yay -S dockerydo-bin
# or
paru -S dockerydo-bin
```

### Debian/Ubuntu

```bash
# Download the .deb file from releases
wget https://github.com/artyultra/dockerydo/releases/download/v<VERSION>/dockerydo_<VERSION>_linux_amd64.deb
sudo dpkg -i dockerydo_<VERSION>_linux_amd64.deb
```

### Red Hat/Fedora/CentOS

```bash
# Download the .rpm file from releases
wget https://github.com/artyultra/dockerydo/releases/download/v<VERSION>/dockerydo_<VERSION>_linux_amd64.rpm
sudo rpm -i dockerydo_<VERSION>_linux_amd64.rpm
```

### Alpine Linux

```bash
# Download the .apk file from releases
wget https://github.com/artyultra/dockerydo/releases/download/v<VERSION>/dockerydo_<VERSION>_linux_amd64.apk
sudo apk add --allow-untrusted dockerydo_<VERSION>_linux_amd64.apk
```

### From Source

```bash
git clone https://github.com/artyultra/dockerydo.git
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

This project is licensed under the MIT License.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Roadmap

- [ ] Additional container operations
- [ ] Docker Compose support
- [ ] Performance metrics and monitoring
- [ ] Export capabilities
- [ ] Configuration file support
