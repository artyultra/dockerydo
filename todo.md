# dockerydo - TODO List

## =4 High Priority - Core Functionality

### Container Operations

- [~] Implement start/stop/restart container actions
- [ ] Add container removal with confirmation prompt
- [ ] Add pause/unpause functionality
- [ ] Implement container creation/run interface
- [ ] Add bulk operations (multi-select containers)

### Log Viewing

- [ ] Implement container log viewer (real-time streaming)
- [ ] Add log follow mode (`docker logs -f`)
- [ ] Add log filtering/search within logs
- [ ] Support toggling timestamps in logs
- [ ] Add ability to specify number of log lines to show

### Better Port Display

- [x] Handle complex port mappings (multiple ports)
- [x] Display all port mappings for a container (not just first)
- [x] Add protocol display (TCP/UDP)
- [x] Show IP binding info (0.0.0.0 vs specific IPs)

## =á Medium Priority - Enhanced Features

### Search & Filtering

- [ ] Add fuzzy search/filter for container names
- [ ] Implement state-based filtering (running, stopped, all)
- [ ] Add label-based filtering
- [ ] Add image-based filtering
- [ ] Highlight search matches in table

### Resource Monitoring

- [ ] Show real-time CPU usage per container
- [ ] Show memory usage and limits
- [ ] Add network I/O stats
- [ ] Display disk usage per container
- [ ] Create a stats dashboard view

### Container Details Enhancements

- [ ] Show environment variables
- [ ] Display volume mounts with full paths
- [ ] Show network details (IPs, networks, aliases)
- [ ] Display health check status
- [ ] Show restart policy and count
- [ ] Add exec/shell into container functionality

### Docker Compose Integration

- [ ] Show docker-compose projects grouping
- [ ] Add compose service actions (up/down/restart service)
- [ ] Display compose file paths (clickable)
- [ ] Show service dependencies graph

## =â Nice to Have - Polish & UX

### UI/UX Improvements

- [ ] Add sortable columns (click header or keyboard shortcut)
- [ ] Implement multi-pane layout option
- [ ] Add container health status indicators (colors)
- [ ] Show uptime in human-readable format (days, hours)
- [ ] Add loading spinners for async operations
- [ ] Implement confirmation dialogs for destructive actions
- [ ] Add status bar with summary stats (total, running, stopped)

### Configuration

- [ ] Create config file support (~/.config/dockerydo/config.yaml)
- [ ] Add configurable refresh interval
- [ ] Allow custom keyboard shortcuts
- [ ] Support custom color themes
- [ ] Save column width preferences
- [ ] Remember last view mode (table/detailed)

### Theme & Appearance

- [ ] Add theme switcher (multiple Catppuccin variants)
- [ ] Support light/dark mode toggle
- [ ] Add custom color scheme support
- [ ] Make table columns configurable (show/hide)

## =5 Advanced Features

### Image Management

- [ ] Add image list view
- [ ] Show image details (layers, size, tags)
- [ ] Implement image removal
- [ ] Add image pull functionality
- [ ] Show dangling images

### Network Management

- [ ] List Docker networks
- [ ] Show containers per network
- [ ] Create/remove networks
- [ ] Display network driver details

### Volume Management

- [ ] List Docker volumes
- [ ] Show volume usage and containers using them
- [ ] Add volume removal with warnings
- [ ] Display volume driver and mount points

### Remote Docker Support

- [ ] Support Docker context switching
- [ ] Connect to remote Docker daemons
- [ ] Show current Docker context in status bar
- [ ] Support multiple contexts with quick switch

### Export & Reporting

- [ ] Export container list to CSV/JSON
- [ ] Generate container reports
- [ ] Copy container details to clipboard
- [ ] Add screenshot/export of current view

## =à Technical Improvements

### Code Quality

- [ ] Add unit tests for Docker parsing functions
- [ ] Add integration tests for UI components
- [ ] Improve error handling and display
- [ ] Add logging functionality (debug mode)
- [ ] Document public functions and types

### Performance

- [ ] Optimize container list refresh (only fetch changed data)
- [ ] Add caching for docker inspect results
- [ ] Implement lazy loading for detailed views
- [ ] Reduce memory footprint for large container lists

### Build & Distribution

- [ ] Add GitHub Actions CI/CD
- [ ] Create release builds for multiple platforms
- [ ] Add Homebrew formula
- [ ] Create AUR package (Arch Linux)
- [ ] Add installation instructions to README
- [ ] Create demo GIF/video for README

## =Ý Documentation

- [ ] Write comprehensive README with features list
- [ ] Add keyboard shortcuts reference
- [ ] Create user guide with screenshots
- [ ] Document configuration options
- [ ] Add contributing guidelines
- [ ] Create architecture documentation

## = Bug Fixes & Edge Cases

- [ ] Fix ViewPort height calculation margins
- [ ] Handle edge cases in port parsing
- [ ] Test with containers having no labels
- [ ] Handle very long container names gracefully
- [ ] Test with 100+ containers (performance)
- [ ] Handle docker daemon connection failures gracefully
- [ ] Add proper error messages when Docker is not running

## <¯ Quick Wins (Easy Implementations)

- [ ] Add "copy container ID" keyboard shortcut
- [ ] Show container count in footer
- [ ] Add "time since last refresh" indicator
- [ ] Implement "jump to top/bottom" keyboard shortcuts (g/G)
- [ ] Add help screen with keyboard shortcuts (press ?)
- [ ] Show Docker version in status bar
- [ ] Add color-coded container states (green=running, red=stopped)

---

**Priority Legend:**

- =4 High Priority: Core functionality users expect
- =á Medium Priority: Enhanced features that add significant value
- =â Nice to Have: Polish and improved user experience
- =5 Advanced Features: Power user and advanced functionality
