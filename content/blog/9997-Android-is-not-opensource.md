+++
Title = "[2021.04.22] Android is not open source."
Layout = "page"
Collection = "Blog"
CrumbParent = "blog"
Tags = [ "BTnet" ]
+++

Android is not open source, and can we, please, stop pretending that it is?

# [2021.04.22] Android is not open source

What come in your mind when you think _OpenSource_? I see Richard Stallman, DFSG, and hundreds of merge requests and issues, where people discuss what should happen with the project, what should be done, and what shouldn't. How does it look like in android?

On [source.android.com](https://source.android.com/) we can read this:

 > As an open source project, Android's goal is to avoid any central point of failure in which one industry player can restrict or control the innovations of any other player. To that end, Android is a full, production-quality operating system for consumer products, complete with customizable source code that can be ported to nearly any device and public documentation that is available to everyone.

Wow! That's great, but wait, on [source.android.com/setup/start/faqs](https://source.android.com/setup/start/faqs), I have read that:

 > some parts of the next version of Android including the core platform APIs are developed in a private branch.

and that source code is released _When it's ready_.
 > Releasing the source code is a fairly complex process. Some parts of Android are developed in the open, and that source code is always available. Other parts are developed first in a private tree, and that source code is released when the next platform version is ready.

Happens in private because what? So when [somebody will rely on your system](https://github.com/termux/termux-app/issues/1072), and find a bug, instead of situation getting discussed, they will get [Won't Fix (Intended behavior)](https://issuetracker.google.com/issues/128554619). But come on! At least they are sorry for that _We apologize for the undocumented change in Android Q. The documentation will be corrected in a future release._

Linux docs have nicely explained [why it is important, to develop in an open way](https://www.kernel.org/doc/html/latest/process/1.Intro.html). But sure, better stick with 4.9 kernel, some proprietary drivers, and never get mainline, so your device won't be fully usable.

But let's not argue if google's way of developing things is good or bad (it's bad). It doesn't matter, because AOSP is not whay you are using on your device, you are running proprietary OS, that passes Android Compatibility Test Suite, provided by your manufacturer.

Okay, okay, but you can build AOSP from source... Can you? Not from source, you need `/vendor/`, your `Android.mk` file includes files that are not in the repository, files that are shipped by your vendor, by the one who maintain it all.

At least kernel is open source, but instead of mainlining it, vendors choose to not care about it.

So to sum this up. Android is not open source. Parts of it's code are available, but it has nearly nothing to do with with `FOSS` that you know from any other project.