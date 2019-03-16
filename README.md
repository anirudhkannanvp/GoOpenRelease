# GoOpenRelease
GoOpenRelease is a tool written in GoLang,is a simple application that gives us the highest patch version of every release between a minimum version and the highest released version by reading the Github Releases and using Semver for Comparison

Setting up the development environment:-

The tool was tested using UBUNTU 16.04 and go version :- go1.11 linux/amd64

STEPS TO SET UP DEVELOPMENT ENVIRONMENT:-

1) UPGRADE THE TOOLS

- sudo apt-get update
- sudo apt-get -y upgrade

2) DOWNLOAD THE FILE IN /tmp 

- cd /tmp
- wget https://dl.google.com/go/go1.11.linux-amd64.tar.gz

3) UNZIP AND MOVE THE FILE TO usr/local

- sudo tar -xvf go1.11.linux-amd64.tar.gz
- sudo mv go /usr/local

4) SET THE GOROOT, GOPATH AND PATH

- export GOROOT=/usr/local/go
- export GOPATH=$HOME/go
- export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

5) UPDATE THE SESSION USING THE COMMAND GIVEN BELOW OR OPEN A NEW SESSION IN TERMINAL

- source ~/.profile


COMMENTS/ IMPROVEMENTS OF THE APP:-
1) This is my first Application in Go apart from a few other problems submitted in Competitive Coding before .
2) If you find any ERRORS, BETTER WAYS, etc , Please send me a Pull Request and I'll be Happy to accept it
