---
title: "The hidden cost of simplicity - P3pch4t Devlog #3"
description: What is the real cost of simple api powered by p2p network? Let's investigate that and create a solution
slug: p3pch4t-devlog-part-3
date: 2023-12-13 00:00:00+0000
# image: cover.jpg
categories:
    - Projects
tags:
    - P3pch4t
    - I2p
weight: 1
---

# The hidden cost of simplicity - P3pch4t Devlog #3 [2023.12.13]

So one of the main goals of p3p is to be as simple as it is possible to create custom clients, bots, extensions, and plugins. For that reason, p3pch4t supports WebXDC and has a library for many languages (currently Golang and Dart, with Python and JavaScript being on the roadmap). But, is it 100% cost-free to go this way...? As you may have guessed from the headline - no, this is not the case.

## Simplicity is not... simple

Let's [skip the configuration](https://git.mrcyjanek.net/p3pch4t/chucknorris/src/commit/6801fec7dbd86393153312499cee11ff887a7e6c/main.go#L22-L38) and go straight to developing the bot. We will base ourselves on the Chuck Norris bot.

So let's create the entire bot.

```go
func main() {
    botPi = core.OpenPrivateInfo(path.Join(os.Getenv("HOME"), ".config", ".p3pchucknorris"), "Chuck Norris", "")
    botPi.Endpoint = core.Endpoint(os.Getenv("PRIVATEINFO_ROOT_ENDPOINT"))
    if !botPi.IsAccountReady() {
        botPi.Create("Chuck Norris", "chuchnorris-jokes@mrcyjanek.net", 4096)
    }
    botPi.MessageCallback = append(botPi.MessageCallback, botMsgHandler)

    <-make(chan bool)
}

var welcomeMessage = "Use !joke to get a joke"

func botMsgHandler(pi *core.PrivateInfoS, ui *core.UserInfo, evt *core.Event, msg *core.Message) {
    if msg.Body != "!joke" {
        pi.SendMessage(ui, core.MessageTypeText, welcomeMessage)
        return
    }
    pi.SendMessage(ui, core.MessageTypeText, getJoke())
}

func getJoke() {
    return "Placeholder Joke #1337"
}
```

This code is pretty self-explanatory

1. First we open some sessions in our .config directory
2. Then if the account isn't ready - we create it
3. Put some callback on the message
4. If the message is anything but `!joke` - we send a help message
5. If it is `!joke` - we send the joke

Pretty simple and straightforward, huh? Well, that is until you realize that there is plenty of stuff going on in the background that you may not want.

- Everything is stored, all messages, introductions - everything. This may be desired for some, but for others, this will simply waste space.
- All features all mandatory. If somebody sends a 1GB movie to you, p3pgo (under current defaults) will try to download it which.. may not be the desired solution for some.

For that reason, I've decided to create `p3pmini.go` which will only do the bare minimum of what's needed to make sure that only features that are required will be used. While simplicity will be hurt, at least a bit - we will have a much easier-to-use library for lower-level projects, and for those who prefer less bloat.

Speaking of which - the `p3pgroup` Golang server is done, but it works so bad due to the reasons mentioned above that I need to write `p3pmini.go` first, and then rewrite p3pgo in it to be at least somewhat satisfied with it. Currently, it is usable but not as optimal as I'd like it to be

--------

Project state as of 2023.12.13


Task:
1. Core functionality
    1. [x] Running and configuring embedded i2p router.
    2. [x] Client handshake (Ensuring proper key exchange, showing fingerprint to both users to verify that the contact was properly added).
    3. [ ] ~~Edit the flutter's PGP library to add missing functionality~~ This task no longer applies, we have moved (at least for now) this logic to Golang code
    4. [ ] Event exchange (Ensuring that each event is delivered, grouping events by destination) _Partially done, we do exchange events. I'll consider this step done when the relays arrive_
    5. [ ] ~~Combining all possible events into a Dart library (This includes creating the events and designing how they should look - so we can ensure an efficient and easy-to-implement way of exchanging them.)~~ This step also doesn't apply. But we do have a Golang version of this, still partially WIP.
    6. [ ] Documenting the changes
    7. [ ] Creating an automatized build process to ensure users can make their own versions
    8. [ ] Bot API with libraries for Golang and Python/JavaScript **We are putting it to test right now!**
    9. [ ] Creating a group server in Golang (using the above library) **S00n!**
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