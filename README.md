[![](https://goreportcard.com/badge/linuxsuren/cgit)](https://goreportcard.com/report/linuxsuren/cgit)
[![](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/linuxsuren/cgit)
[![Contributors](https://img.shields.io/github/contributors/linuxsuren/cgit.svg)](https://github.com/linuxsuren/cgit/graphs/contributors)
[![GitHub release](https://img.shields.io/github/release/linuxsuren/cgit.svg?label=release)](https://github.com/linuxsuren/cgit/releases/latest)
![GitHub All Releases](https://img.shields.io/github/downloads/linuxsuren/cgit/total)

cgit is a wrapper of git.

# Get started

Clone a repo from GitHub without the whole URL: `cgit clone linuxsuren/cgit`

Provide an alias for git: `cgit alias set cm 'checkout master'`, then you can checkout branch to master via: `cgit cm`

List all alias commands: `cgit alias list`

## Mirror

`cgit` can set a mirror address for you if it's very slow with fetching data from GitHub.

Run this command `cgit mirror` in your local git repository directory, 
it'll change the fetch address to `github.com.cnpmjs.org`. Reversing it is very easy, 
just run command `cigt mirror --enable=false`.

# Install

You can install it via [brew](https://github.com/Homebrew/homebrew-core):

`brew install linuxsuren/linuxsuren/cgit`

Or, you can also install it via [hd](https://github.com/LinuxSuRen/http-downloader):

`hd install -t 8 linuxsuren/cgit`
