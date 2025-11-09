# dockerydo

A lightweight Terminal UI (TUI) for managing Docker resources on Linux and macOS, built as a simple alternative to Docker Desktop for basic Docker operations.

## Motivation

Docker Desktop has become increasingly resource-heavy and comes with licensing restrictions for enterprise use. dockerydo was created to provide a lightweight, open-source alternative for developers who need basic Docker management capabilities without the overhead of a full desktop application. Built with Go and the Bubble Tea framework, it offers a fast, terminal-native experience that integrates seamlessly into command-line workflows.

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

## Quick Start

The fastest way to get started with dockerydo:

### macOS/Linux (Homebrew)

```bash
brew install artyultra/dockerydo/dockerydo
dockerydo
```

### Pre-built Binary

```bash
# Download the latest release for your platform
wget https://github.com/artyultra/dockerydo/releases/latest/download/dockerydo_$(uname -s)_$(uname -m).tar.gz

# Extract and install
tar -xzf dockerydo_*.tar.gz
sudo mv dockerydo /usr/local/bin/

# Run it
dockerydo
```

### From Source

```bash
git clone https://github.com/artyultra/dockerydo.git
cd dockerydo
go build
./dockerydo
```

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
dockerydo
```

Or if running from the source directory:

```bash
./dockerydo
```

### Navigation

- **Tab**: Switch between tabs (Containers, Images, Volumes, Networks)
- **Arrow Keys (↑/↓)**: Navigate through lists
- **Enter**: Select item for detailed view
- **Esc**: Return to list view from detail view
- **q**: Quit application

### Common Tasks

#### Managing Containers

1. Navigate to the Containers tab (default view)
2. Use arrow keys to select a container
3. Press Enter to view details and available actions
4. Common actions include:
   - Start/Stop containers
   - Restart containers
   - View logs
   - Remove containers

#### Managing Images

1. Press Tab to switch to the Images tab
2. Select an image to view details
3. Available actions:
   - Remove images
   - View image layers and metadata

#### Managing Volumes

1. Press Tab to navigate to the Volumes tab
2. Select a volume to inspect
3. View volume details and mount points

#### Managing Networks

1. Press Tab to navigate to the Networks tab
2. Select a network to view configuration
3. Inspect connected containers and network settings

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

Contributions are welcome! We appreciate your help in making dockerydo better.

### How to Contribute

1. **Fork the Repository**
   ```bash
   git clone https://github.com/artyultra/dockerydo.git
   cd dockerydo
   ```

2. **Create a Feature Branch**
   ```bash
   git checkout -b feature/your-feature-name
   ```

3. **Make Your Changes**
   - Follow Go best practices and conventions
   - Ensure your code is properly formatted (`go fmt`)
   - Add tests if applicable
   - Update documentation as needed

4. **Test Your Changes**
   ```bash
   go build
   ./dockerydo
   ```

5. **Commit Your Changes**
   ```bash
   git add .
   git commit -m "Add: description of your changes"
   ```

6. **Push and Create a Pull Request**
   ```bash
   git push origin feature/your-feature-name
   ```
   Then open a Pull Request on GitHub with a clear description of your changes.

### Development Guidelines

- Keep code modular and organized following the existing project structure
- Use meaningful variable and function names
- Comment complex logic
- Ensure the UI remains responsive and user-friendly
- Test with different Docker configurations and states

### Reporting Issues

If you find a bug or have a feature request:

1. Check if the issue already exists in the [Issues](https://github.com/artyultra/dockerydo/issues) section
2. If not, create a new issue with:
   - Clear description of the problem or feature
   - Steps to reproduce (for bugs)
   - Expected vs actual behavior
   - System information (OS, Docker version, Go version)

### Code of Conduct

Please be respectful and constructive in all interactions. We're all here to build something useful together.

## Roadmap

- [ ] Additional container operations
- [ ] Docker Compose support
- [ ] Performance metrics and monitoring
- [ ] Export capabilities
- [ ] Configuration file support
