+++
Area = "portfolio"
CrumbParent = "portfolio"
Layout = "page"
Tags = [ "TOR", "FOSS" ]
Title = "Simple Tor File Server - Minimal OnionShare replacement"
+++

# Simple Tor File Server

[![Build Status](https://ci.mrcyjanek.net/badge/build-tor_file_server.svg)](https://ci.mrcyjanek.net/jobs/build-tor_file_server)

Simple Tor File Server is OnionShare replacement, that do only one thing, serve things over http, over tor, it's less than a 100 lines of code.

Usage is as simple as this:

```bash
$ tor-file-server # To host current directory
$ tor-file-server -path=/home # To share your home dir
User: pSbn
Pass: kHBsjcmevSMahOdJ
Host: 7vtzlno7rsthisisnotvalidtorlinkn3kywj3sz7i3t2dqk4vlzqfgsfid.onion:9832/
```

# Installation

Simply run `go build` inside of this repo, or download prebuilt from [my ci](https://ci.mrcyjanek.net/jobs/build-tor_file_server)

On debian you can get it from my repo:

```bash
# wget 'https://static.mrcyjanek.net/laminarci/apt-repository/cyjan_repo/mrcyjanek-repo-latest.deb' && \
  apt install ./mrcyjanek-repo-latest.deb && \
  rm ./mrcyjanek-repo-latest.deb && \
  apt update
```
After that install tor-file-server
```bash
# apt install tor-file-server
```

