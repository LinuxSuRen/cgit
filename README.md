[![](https://goreportcard.com/badge/linuxsuren/cgit)](https://goreportcard.com/report/linuxsuren/cgit)
[![](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/linuxsuren/cgit)
[![Contributors](https://img.shields.io/github/contributors/linuxsuren/cgit.svg)](https://github.com/linuxsuren/cgit/graphs/contributors)
[![GitHub release](https://img.shields.io/github/release/linuxsuren/cgit.svg?label=release)](https://github.com/linuxsuren/cgit/releases/latest)
![GitHub All Releases](https://img.shields.io/github/downloads/linuxsuren/cgit/total)

cgit is a wrapper of git.

# Get started

Clone a repo from GitHub without the whole URL: `cgit clone linuxsuren/cgit`

Sometimes it's very slow when clone the code from GitHub. So cgit will clone it by [a GitHub proxy](http://github.com.cnpmjs.org/).

Provide an alias for git: `cgit alias set cm 'checkout master'`, then you can checkout branch to master via: `cgit cm`

List all alias commands: `cgit alias list`

# Install

```
brew install linuxsuren/linuxsuren/cgit
```
