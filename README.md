[![](https://goreportcard.com/badge/linuxsuren/cgit)](https://goreportcard.com/report/linuxsuren/cgit)
[![](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/linuxsuren/cgit)
[![Contributors](https://img.shields.io/github/contributors/linuxsuren/cgit.svg)](https://github.com/linuxsuren/cgit/graphs/contributors)
[![GitHub release](https://img.shields.io/github/release/linuxsuren/cgit.svg?label=release)](https://github.com/linuxsuren/cgit/releases/latest)
![GitHub All Releases](https://img.shields.io/github/downloads/linuxsuren/cgit/total)

cgit is a wrapper of git.

# Features

* Clone a repo from GitHub without the whole URL
* GitHub proxy transparent support
* Git command alias support 

# Install

```
brew install linuxsuren/linuxsuren/cgit
```

cgit is fully compatible with git. So you make an alias for it. Add the following line into you shell profile:

`alias git='cgit'`

For bash users, you edit it via: `vim ~/.bashrc`

For zsh users, you can edit via: `vim ~/.zshrc`

# Get started

## Clone 

`cgit clone linuxsuren/cgit`

## GitHub Proxy

Sometimes it's very slow when clone the code from GitHub. So cgit will clone it by [a GitHub proxy](http://github.com.cnpmjs.org/).

## Alias

Add a command alias: `cgit alias set cm 'checkout master'`

Use an alias: `cgit cm`

List all alias commands: `cgit alias list`
