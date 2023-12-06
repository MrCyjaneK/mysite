---
title: "P3pch4t Devlog #1 2023/11/29"
description: It's been a long way since the project began, let's summarize everything in this overly specific post
slug: p3pch4t-devlog-part-1
date: 2023-11-29 00:00:00+0000
# image: cover.jpg
categories:
    - Projects
tags:
    - P3pch4t
    - I2p
weight: 1
---

# [2023.11.29] P3pch4t Devlog #1

Okay, it's been a while. Let's start from the beginning, what is p3pch4t? It **was** a draft, made 9 months ago by me, mostly in a [single commit](https://github.com/MrCyjaneK/p3pch4t/commit/a3140a15e9e257e39786691ea9b4e4f48ffc8aac). That draft app is pretty much working, but it is far from perfect, so I've decided to not fix it but to create an entirely separate entity one more time and focus on making it from scratch. Why from scratch? I believe that it is much better to start from the ground up after learning many important things from the first project.

## Insights Gained

Thanks to a few people who helped me test the project in the early stages, or offered me their knowledge in the form of conversation, I was able to gain some important takeaways regarding the project direction, here are some of them.

### Stripping Away the Unnecessary

My initial approach was to create one app with all the features inside (as you can see in the last commits I've even started to implement a web browser as a part of the chat app). Well, guess what? That's not what users want, instead of making everything in one app I was told that it is better to make a simple project that is extensible through plugins.

That's what I've done in the new rewrite of the app. We got [WebXDC](https://webxdc.org/) (with the help of [@Im-Beast](https://im-beast.com/)) as the main feature-driving factor, this way we can keep all the core features simple, and allow users to use as many features as they want without forcing them to learn how to navigate thru the bloat.

### User-Centric Design

Next on the road was simplicity. I knew that users who aren't technical need a way to use the code, so obviously I've.. [put 12-step instructions on how to enable I2P tunnel on Android..](https://github.com/MrCyjaneK/p3pch4t/tree/a3140a15e9e257e39786691ea9b4e4f48ffc8aac/assets/images/help). Sadly, a significant number of testers (all 3 of them) found themselves bored and anxious at the sight of step 1 out of 12.

Conquering the challenge, I have decided to ship I2Pd together with the app, to accomplish this I've created a few repositories

 - [dart_i2p](https://git.mrcyjanek.net/p3pch4t/dart_i2p) - Home for the configuration, logic, and build scripts
 - [flutter_i2p](https://git.mrcyjanek.net/p3pch4t/flutter_i2p) - Flutter UI that allows to browse logs and edit configuration (well, this is a planned feature - currently configurable options are pretty limited).
 - [flutter_i2p_bins-prebuild](https://git.mrcyjanek.net/p3pch4t/flutter_i2p_bins-prebuild) - Simply an online repository of all precompile i2pd binaries.

### Bells and Whistles

Users want fanciness, and I'm going to deliver on it, chat background, and custom colors, all were present in the draft version and will be eventually ported to the new version (I have more important things to do right now than changing backgrounds)

### Bridging the Gap: P2P with User-Focus

After seeing how p2p networks work most people reacted in a way that I can sum up using this quote:

> If this is the solution, I want my problem back.

When using other chat apps I've noticed many wrong things, file sharing usually was limited to a fraction of a megabyte, network usage was insanely high, battery drain was unacceptably high and you couldn't reach a user (or leave the message on a server so it will reach the user when available) when one party was offline, then we have discovering new users, using multiple devices etc...

That's why p3prelay was born, and instantly buried alive - but in its time it will come back to life. P3prelay had one job - to get a signed and encrypted message and relay it to the given address, grab a response, and echo it back to the user when polled for it. As simple as it may sound it solves many issues while creating almost none. P3prelay solved:
 - Reaching users when one party is offline.
 - Battery drain (one user may choose to use a relay on clearnet, so no need to run i2pd).
 - Cross-network issues, with my plan to support i2p, tor, and clearnet (with the primary focus being on I2p) users who opt to use tor only should be able to reach i2p users.

But this also creates one tiny issue, which is the trust issue with the relay operator. Luckily relays are fully optional and do not require any specific code changes, relay is literally a user in the middle that relays messages.

Why is it a **tiny** issue? Because it is an opt-in feature that doesn't compromise security. Relays only forward encrypted events (and plaintext `introduce` event). And since there is nothing specific going on between a relay and a normal p2p user it doesn't compromise the security of that user, if the destination is over i2p relay will connect to it over i2p. It can be compared to how emails work, if you self-host your email, and somebody from gmail.com sends you an email does it make your email less self-hosted? In my opinion, it is bad to rely on a single relay so the source code will be fully open (currently there is a rather not-functional draft that just proved the idea to work), and p3pch4t will be able to use multiple relays at once, either by selecting them randomly or using multiple relays at once. This will hopefully solve all the problems, without causing too much trouble. As always - I'm open to suggestions.

## The rewrite

I've started (and finished) creating a working library in Dart and built a fully functional chat on top of it, but after dealing with a few database issues (that's a skill issue on my end), I've decided to just bind to the Golang library for now. Why Go? Matterbridge is written in it. Once I finish groups.. ah yes. groups. we don't do that here, everything is a one-to-one chat.

_Wait what?_

Every chat is one-to-one chat, this includes groups.

_Why?_

It greatly simplifies the design of the protocol, makes management of the group easy, it also takes care of the state of the group (assigning IDs, holding information about who is admin and who's banned). Fully peer-to-peer groups are on my roadmap, but speaking from the UX side I doubt that they will ever find many users. Imagine joining a group of ~87 people and sharing a 7MB video. Suddenly we need to route this video somehow between the mesh of the users to make sure that you won't waste 609MB of internet, we also need to ensure that the video was delivered to every single person, so we at least need to ping every of the members to ensure that the message was delivered to them (sure, we can also do some fancy networking thing and make ask 8 users to talk to 8 other random users and bring some signed proof to us that the message was received), but this is heavily increasing the complexity of the protocol while also lowering the UX.

If we use the principle of having a user which is a group server will greatly increase the user's experience by providing something reliable, and for small groups, it will not cause any issues and could easily be hosted on a low-end device and for bigger groups a dedicated device or cloud could be used.

## What's the state

 - WebXDC has not been tested in the latest releases but it [was fully working before recent major updates](https://x.com/mrcyjanex/status/1692504333307273635).
 - One-to-one chats (like one-to-one, the group thing is being implemented at the very moment)
 - Filesystem-like attachments system
 - We have an actual i2p connection together with public key handshake embedded and working on Windows, Linux, and Android out of the box (MacOS, iOS I'm coming for you too. SailfishOS... I'll port flutter to run p3pch4t on you also.)

## What's the plan?

Next on my roadmap is to work on some of the WebXDC apps, bring WebXDC support to the desktop and other exotic platforms (currently it is only on Android), and add support for Matterbridge, to get some testing done. In a few months from now, I plan to move at least parts of my communication to p3pch4t, either by bridges or by forcing people to use my chat app :p, and by porting the app to all platforms I can find.

Also, the next updates will occur more regularly (assuming that from the first commit to this day almost 10 months have passed).

--------

Project state as of 2023.11.29


Task:
1. Core functionality
    1. [x] Running and configuring embedded i2p router.
    2. [x] Client handshake (Ensuring proper key exchange, showing fingerprint to both users to verify that the contact was properly added).
    3. [ ] ~~Edit the flutter's PGP library to add missing functionality~~ This task no longer applies, we have moved (at least for now) this logic to Golang code
    4. [ ] Event exchange (Ensuring that each event is delivered, grouping events by destination) _Partially done, we do exchange events. I'll consider this step done when the relays arrive_
    5. [ ] ~~Combining all possible events into a Dart library (This includes creating the events and designing how they should look - so we can ensure an efficient and easy-to-implement way of exchanging them.)~~ This step also doesn't apply. But we do have a Golang version of this, still partially WIP.
    6. [ ] Documenting the changes
    7. [ ] Creating an automatized build process to ensure users can make their own versions
    8. [ ] Bot API with libraries for Golang and Python/JavaScript
    9. [ ] Creating a group server in Golang (using the above library)
2. User Interface
    1. [ ] Initial UI release - support all features of the backend library
    2. [ ] UI polishing – including better desktop support.
3. Beta Release
    1. [ ] Android
    2. [ ] Linux
    3. [ ] Windows
    4. [ ] MacOS
    5. [ ] iOS
    6. [ ] SailfishOS???
4. Stable Release
    1. [ ] Fix security issues and bugs
    2. [ ] Release the app on stores
5. WebXDC
    1. [ ] Notes
    2. [ ] Todo
    3. [ ] Calendar (with external sync) _Work started, by [@Im-Beast](https://im-beast.com/)._ 