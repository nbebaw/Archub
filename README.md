# Archub

A lightweight and fast command-line AUR (Arch User Repository) helper written in Go for managing AUR packages on Arch Linux systems.

![GitHub](https://img.shields.io/github/license/nbebaw/Archub) ![GitHub all releases](https://img.shields.io/github/downloads/nbebaw/Archub/total) ![GitHub repo size](https://img.shields.io/github/repo-size/nbebaw/Archub) ![GitHub issues](https://img.shields.io/github/issues/nbebaw/Archub) ![GitHub closed issues](https://img.shields.io/github/issues-closed/nbebaw/Archub) ![GitHub release (latest by date)](https://img.shields.io/github/v/release/nbebaw/Archub) ![GitHub last commit](https://img.shields.io/github/last-commit/nbebaw/Archub)

## Features

- 🔍 **Search** AUR packages with detailed results
- 📦 **Install** packages directly from AUR
- 🔄 **Update** individual packages or all AUR packages at once
- ℹ️ **Package Info** - View comprehensive package details before installing
- 📋 **List** all installed AUR packages
- 🗑️ **Delete** packages with dependency removal
- ✅ **Check** for available updates without installing
- 🧹 **Cleanup** temporary build directories
- 💻 **System Info** - Display CPU, memory, and host information
- 🎨 **Color-coded** output for better readability
- ⚡ **Fast** and lightweight, written in Go

## Dependencies

**Required:**
- `git` - For cloning AUR repositories
- `pacman` - Arch Linux package manager (system default)
- `makepkg` - For building packages (part of pacman)
- `curl` - For fetching package information

**Optional:**
- `sudo` or `doas` - For installing packages with elevated privileges

## Shell Auto-Completions

- ✅ **Zsh** - Full auto-completion support included

## Installation

### Quick Install (without auto-completions)

```bash
wget https://github.com/nbebaw/Archub/releases/download/v0.0.5/archub
chmod +x archub
sudo mv archub /usr/bin
```

### Full Install (with Zsh auto-completions)

```bash
mkdir archub_x86_64
cd archub_x86_64
wget https://github.com/nbebaw/Archub/releases/download/v0.0.5/archub_x86_64.tar.gz
tar -xvf archub_x86_64.tar.gz
chmod +x archub
sudo mv archub /usr/bin
sudo mv LICENSE /usr/share/licenses/archub
sudo mv auto_completions/zsh /usr/share/zsh/site-functions/_archub
```

### Build from Source

```bash
git clone https://github.com/nbebaw/Archub.git
cd Archub
go build -ldflags "-X main.Version=v0.0.6" -o archub main.go
sudo mv archub /usr/bin
```
## Usage

```bash
archub [options]
```

### Available Options

| Option | Alias | Description |
|--------|-------|-------------|
| `-s [package]` | `--search` | Search for a package in AUR |
| `-Si [package]` | `--pkginfo` | Show detailed package information |
| `-i [package]` | `--install` | Install a package from AUR |
| `-u [package]` | `--update` | Update a specific package |
| `-u --all` | `--update --all` | Update all AUR packages |
| `-d [package]` | `--delete` | Delete a package |
| `-l` | `--list` | List all installed AUR packages |
| `-c` | `--clean` | Clean up temporary build directories |
| `--check` | | Check for available updates |
| `--info` | | Display system information |
| `--help` | | Show help message |
| `--version` | | Show version information |

## Examples

### Search for a package

```bash
archub -s yay
# or
archub --search yay
```

### View package information

Display comprehensive package details including dependencies, votes, maintainer, and more:

```bash
archub -Si paru
# or
archub --pkginfo paru
```

**Example output:**
```
Repository      : aur
Name            : paru
Version         : 2.1.0-2
Description     : Feature packed AUR helper
Maintainer      : Morganamilo
Votes           : 1148
Popularity      : 35.17
Depends On      : git  pacman  libalpm.so>=14
Optional Deps   : bat
                  devtools
Installed       : No
```

### Install a package

```bash
archub -i package_name
# or
archub --install package_name
```

### Update packages

Update a specific package:
```bash
archub -u package_name
# or
archub --update package_name
```

Update all AUR packages:
```bash
archub -u --all
# or
archub --update --all
```

### List installed AUR packages

```bash
archub -l
# or
archub --list
```

### Delete a package

Removes the package along with its dependencies:
```bash
archub -d package_name
# or
archub --delete package_name
```

### Check for updates

Lists packages with available updates without installing them:
```bash
archub --check
```

### Clean up temporary files

Remove temporary build directories:
```bash
archub -c
# or
archub --clean
```

### View system information

```bash
archub --info
```
## How It Works

1. **Search**: Queries the AUR RPC v5 API for package information
2. **Install**: Clones the AUR git repository to `~/.config/archub_tmp`, then builds using `makepkg -si`
3. **Update**: Fetches the latest PKGBUILD version from AUR and compares with installed version
4. **Package Info**: Retrieves comprehensive metadata from AUR including dependencies, popularity, and maintainer information

## Configuration

Archub uses `~/.config/archub_tmp` as its temporary directory for cloning and building packages. This directory can be cleaned up using the `-c` flag.

## Comparison with Other AUR Helpers

| Feature | Archub | yay | paru |
|---------|--------|-----|------|
| Written in | Go | Go | Rust |
| Package Search | ✅ | ✅ | ✅ |
| Package Info | ✅ | ✅ | ✅ |
| Install Packages | ✅ | ✅ | ✅ |
| Update Packages | ✅ | ✅ | ✅ |
| Lightweight | ✅ | ✅ | ✅ |
| Simple Interface | ✅ | ❌ | ❌ |
| AUR Dependency Resolution | ❌ | ✅ | ✅ |
| PKGBUILD Review | ❌ | ✅ | ✅ |

Archub focuses on simplicity and speed, providing core AUR functionality without the complexity of larger helpers.

## Contributing

Contributions are welcome! Here's how you can help:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development

```bash
# Clone the repository
git clone https://github.com/nbebaw/Archub.git
cd Archub

# Build
go build -o archub main.go

# Run tests
go test ./...

# Run locally
./archub -s test-package
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with Go and love for the Arch Linux community
- Uses the [AUR RPC API](https://aur.archlinux.org/rpc/) for package information
- Inspired by popular AUR helpers like yay and paru
