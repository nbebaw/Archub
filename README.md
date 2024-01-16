# Archub
Archub is a command-line tool designed for managing packages in Arch Linux's AUR (Arch User Repository).

![GitHub](https://img.shields.io/github/license/nbebaw/Archub) ![GitHub all releases](https://img.shields.io/github/downloads/nbebaw/Archub/total) ![GitHub repo size](https://img.shields.io/github/repo-size/nbebaw/Archub) ![GitHub issues](https://img.shields.io/github/issues/nbebaw/Archub) ![GitHub closed issues](https://img.shields.io/github/issues-closed/nbebaw/Archub) ![GitHub release (latest by date)](https://img.shields.io/github/v/release/nbebaw/Archub) ![GitHub last commit](https://img.shields.io/github/last-commit/nbebaw/Archub)

## Dependencies
- <b>git</b>

## Terminal Auto Completions support
- Zsh

## Installation without zsh auto completions
```sh
wget https://github.com/nbebaw/Archub/releases/download/v0.0.4/archub
chmod +x archub
sudo mv archub /usr/bin
```
## Installation with zsh auto completions
```sh
mkdir archub_x86_64
cd archub_x86_64
wget https://github.com/nbebaw/Archub/releases/download/v0.0.4/archub_x86_64.tar.gz
tar -xvf archub_x86_64.tar.gz
chmod +x archub
sudo mv archub /usr/bin
sudo mv LICENSE /usr/share/licenses/archub
sudo mv auto_completions/zsh /usr/share/zsh/site-functions/_archub
```
## Usage

```bash
archub [options]
Usage:
 archub [options]

Options:
 -s, --search [package]          : Search a package
 -i, --install [package]         : Install a package
 -c, --clean                     : Clean up
 -l, --list                      : List all aur installed Packages
 -u, --update [package]          : Update a Package
 -u --all, --update --all        : Update all Packages
 -d, --delete [package]          : Delete a Package
 --check                         : Check for updates
 --help                          : How to use Archub
 --version                       : Version of Archub
```

## Examples
Search a package in AUR (You can also use --search)
```sh
archub -s package_name
```
Install a package from AUR (You can also use --install)
```sh
archub -i package_name
```
Update a specific package (You can also use --update)
```sh
archub -u package_name
```
Update a all packages (You can also use --update)
```sh
archub -u --all
```
Delete a package (You can also use --delete)
```sh
archub -d package_name
```
Check for updates
```sh
archub --check
```
## Contributing
Feel free to contribute by opening issues or pull requests.

## License
This project is licensed under the MIT License.
