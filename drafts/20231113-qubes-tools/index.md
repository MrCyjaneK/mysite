---
title: Qubes-Tools app for dom0 (part 1)
description: Creating Qubes-Tools app, so terminal won't be a must-use app when daily-driving qubes.
slug: creating-qubes-tools-app
date: 2023-11-13 00:00:00+0000
# image: cover.jpg
categories:
    - Projects
tags:
    - OS
    - Security
    - Qubes
weight: 1
---

# [2023.11.13] Qubes-Tools app for dom0 (part 1)

I'm a Qubes user since 4 years (with a little break in between). My configuration is usually one "native" os, and Qubes (currently NixOS and Qubes 4.2). There are some tasks that simply cannot be achieved on QubesOS - for example watching movies when on battery.. most I'm able to achieve is around an hour and few minutes under "average" load, while at the same time doing the same things I can achieve up to 4 hours of battery, this is somewhat to be expected.

I'm personally very happy with how well QubesOS is behaving for me. It is very straightforward to do almost everything, especially considering the fact that I am software developer. I can easily spin up windows vm, install some specific linux distro, or open files / install apps in DispVMs.

The thing is. I consider myself a poweruser of my machine, when I'm told to add a parameter to cmdline I know that

 - I am supposed to be using grub (instead of systemd-boot or other solutions)
 - I know which file to edit
 - I need to know where that cmdline is
 - I need to know what command do I need to run to apply the update

That 'add a parameter to cmdline' sounds like adding salt to already done meal that's on already set table, with your dream wife sitting at the other side of the table... While instead you need to learn how to cook, buy the ingredients, assemble an Ikea table and once that is done you need to find your dream wife.

At least that's what I've been told by my friend when I don't him to do some simple cmdline change.

# A solution

What I've done (long time ago) was to create a set of shell scripts that achieved some simple tasks for me, and provided an interactive session. This was fine, and worked well for me - but this is not really a solution that you would like to present to some computer newbie.

We need something that is easy to use, something that handles all the edge-cases that may occur. Initially I've wanted to create an app in flutter but.. bringing prebuilds compiled by google... meh, that doesn't sound like a wise idea. Not to mention the fact that it is unrealistic to expect flutter to work over tor, let alone use it offline.

So I've decided to go with something new - that I've never tried before.

# The solution

Using [fyne toolkit](https://fyne.io/) - made to work in Go, without the need to grab any external dependencies, except for few packages from debian repo, and the [qubes-tools](https://git.mrcyjanek.net/mrcyjanek/qubes-tools) repository. I've made a simple app. It doesn't contain many features now but it does contain some features that I personally need. If you want any extra features please [contact me](/contact) me.

Features currently in app:

 - qvm-create-windows-qube Gui - simple interface that covers all qvm-create-windows-qubes use-cases, starting from the instalation, going thru downloading the iso, [fixing issue #75](https://github.com/ElliotKillick/qvm-create-windows-qube/issues/75) and creating the windows qube.
 - backup utility - create backup of qubes in a incremental way, much faster than built-in tool.
 - update - tool can recompile itself using latest source in dispvm