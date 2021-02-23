# autotimezone - Automated TimeZone Utility for Installers

## Requirements 
* GhostBSD or FreeBSD
* go 1.16 or later
* tzsetup
* Network connection
* Root privileges (i.e., privileges via doas)

## Setup
``` 
sudo pkg install go
git clone https://github.com/vimanuelt/autotimezone
cd autotimezone
go build autotimezone.go
```
## Run
```
./autotimezone
```
