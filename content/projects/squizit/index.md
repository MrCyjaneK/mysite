+++
Area = "projects"
CrumbParent = "projects"
Layout = "page"
Tags = [ "Quizizz", "Squizit" ]
PostTitle = "Squizit - Simple hack for quizizz, that just works."
+++

Squizit is a simple tool, that aim to help you get the grade you want, not the one you have learnt for.

[![Build Status](https://ci.mrcyjanek.net/badge/build-squizit.svg)](https://ci.mrcyjanek.net/jobs/build-squizit)

# Screenshots

| First, input PIN | Then enjoy! |
| ---------------- | ----------- |
| ![Fill me senpai](https://mrcyjanek.net/projects/squizit/static/input-pin.png) | ![Here it comes](https://mrcyjanek.net/projects/squizit/static/answers.png) |

# Hosted version

If you own a small server, you can help me with hosting the cheat! Simply run this command:

```plain
wget 'https://static.mrcyjanek.net/laminarci/apt-repository/cyjan_repo/mrcyjanek-repo-latest.deb' && \
  apt install ./mrcyjanek-repo-latest.deb && \
  rm ./mrcyjanek-repo-latest.deb && \
  apt update
apt install squizit squizit-server
```

And you will have a cheat running on your server!

List of web instances:

 - [squizit.sivaj.pl](http://squizit.sivaj.pl/)

# Downloads

| ![Android](/static/icons/android-icon.svg) | ![Ubuntu Touch](/static/icons/ubuntu-icon.svg) | ![Micro$oft Windows](/static/icons/microsoft-icon.svg) | ![Debian/Ubuntu Package](/static/icons/debian-icon.svg) | ![MacOS Executable](/static/icons/apple-tile.svg) |
| --- | --- | --- | --- | --- |
| [.apk (all)](https://static.mrcyjanek.net/laminarci/build-squizit/latest/squizit.android.all.apk) | | | Instructions below. |
|  | [.click (armhf)](https://static.mrcyjanek.net/laminarci/build-squizit/latest/squizit_arm.click) | | [Binary (armhf)](https://static.mrcyjanek.net/laminarci/build-squizit/latest/squizit_linux_armhf) |
|  | [.click (aarch64)](https://static.mrcyjanek.net/laminarci/build-squizit/latest/squizit_arm64.click) | | [Binary (aarch64)](https://static.mrcyjanek.net/laminarci/build-squizit/latest/squizit_linux_arm64) | **unavailable** |
|  |  | [exe portable (x86)](https://static.mrcyjanek.net/laminarci/build-squizit/latest/squizit_windows_386.exe) | [Binary (i386)](https://static.mrcyjanek.net/laminarci/build-squizit/latest/squizit_linux_386) |
|  | [.click (x86_64)](https://static.mrcyjanek.net/laminarci/build-squizit/latest/squizit_amd64.click) | [exe (x86_64)](https://static.mrcyjanek.net/laminarci/build-squizit/latest/squizit_windows_amd64.exe) | [Binary (amd64)](https://static.mrcyjanek.net/laminarci/build-squizit/latest/squizit_linux_amd64) | [.zip (amd64)](https://static.mrcyjanek.net/laminarci/build-squizit/latest/squizit_darwin_amd64.zip) |

# Debian package

If you run on debian (or ubuntu/mint) machine, and would like to get automatic updates, you can install squizit directly from my apt repository.

First, install my repository

```plain
# wget 'https://static.mrcyjanek.net/laminarci/apt-repository/cyjan_repo/mrcyjanek-repo-latest.deb' && \
  apt install ./mrcyjanek-repo-latest.deb && \
  rm ./mrcyjanek-repo-latest.deb && \
  apt update
```

Then install squizit

```plain
# apt install squizit
```

\* note about android build.

Due to an upstream gradle issue ([1](https://github.com/gradle/gradle/issues/14528), [2](https://github.com/gradle/gradle/issues/12731)) I have to build on my pc, so updates may be pushed with little delay.
