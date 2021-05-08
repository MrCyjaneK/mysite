+++
Area = "projects"
CrumbParent = "projects"
Layout = "page"
Tags = [ "ByteStreamer", "Torrents", "FOSS" ]
PostTitle = "ByteStreamer - Stream torrents from pte.nu directly on your device."
+++

# ByteStreamer

[![Build Status](https://ci.mrcyjanek.net/badge/build-bytestreamer.svg)](https://ci.mrcyjanek.net/jobs/build-bytestreamer)

Stream torrents from PolishTracker!

ByteStreamer is a simple app that aim to replace BitStreamer on platforms that it is not available - that is Android.

<!-- TBD: Screenshots. -->

This app, instead of looking for executables of various players, downloads .m3u8 file (on desktop), and that file contain stream link, which plays automatically (tested with vlc).

On android it simply 

# Downloads

| ![Android](/static/icons/android-icon.svg) | ![Micro$oft Windows](/static/icons/microsoft-icon.svg) | ![Debian/Ubuntu Package](/static/icons/debian-icon.svg) |
| --- | --- | --- |
| [apk](https://static.mrcyjanek.net/laminarci/build-bytestreamer/latest/android.all.apk) | [x86_64](https://static.mrcyjanek.net/laminarci/build-bytestreamer/latest/bytestreamer_windows_amd64.exe) | [amd64](https://static.mrcyjanek.net/laminarci/build-bytestreamer/latest/bytestreamer_linux_amd64)
| | [x86.exe](https://static.mrcyjanek.net/laminarci/build-bytestreamer/latest/bytestreamer_windows_386.exe) | [i386](https://static.mrcyjanek.net/laminarci/build-bytestreamer/latest/bytestreamer_linux_386) |
| | | [aarch64](https://static.mrcyjanek.net/laminarci/build-bytestreamer/latest/bytestreamer_linux_arm64) |
| | | [arm](https://static.mrcyjanek.net/laminarci/build-bytestreamer/latest/bytestreamer_linux_arm) |

## Debian Repository

If you wish to receive updates automatically, you can use my repository.

```plain
# wget 'https://static.mrcyjanek.net/laminarci/apt-repository/cyjan_repo/mrcyjanek-repo-latest.deb' && \
  apt install ./mrcyjanek-repo-latest.deb && \
  rm ./mrcyjanek-repo-latest.deb && \
  apt update
```

```plain
# apt install bytestreamer
```

# Tech Details.

All apps (windows, linux, android) share the same codebase. And are packaged using [Goprod](https://mrcyjanek.net/projects/goprod/).

<!-- Yup, I made it for you c: -->