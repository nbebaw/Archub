# Archub
Archub is a command-line tool designed for managing packages in Arch Linux's AUR (Arch User Repository).

## Dependencies
- <b>git</b>

## Installation
```sh
wget https://github.com/nbebaw/archub/releases/download/v0.0.1/archub
chmod +x archub
sudo mv archub /usr/local/bin
```
## Usage

```bash
archub [options]
Usage:
 archub [options]

Options:
 -s, --search [options]          : Search a package
 -i, --install [options]         : Install a package
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
This project is licensed under the Apache License.