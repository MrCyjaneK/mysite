---
title: Linux subsystem for Windows
description: Since when does Linux have less compatibility than other inferior OSes? Why is there no LSW somewhere?
slug: linux-subsystem-for-windows
date: 2023-12-20 00:00:00+0000
# image: cover.jpg
categories:
    - Linux
tags:
    - General
    - Linux
    - Docker
weight: 1       # You can add weight to some posts to override the default sorting (date descending)
---

# LSW - running Windows, the Linux way in docker (on Linux).

Linux (Windows) subsystem for Windows (for Linux)

> Also, am I the only person who thinks that a better name (for the competing product, WSL) would be _Linux subsystem for Windows_?

P.S. There is a TL;DR at the end of the article that allows you to skip 95% of the tasks here. If you came here intending to deploy your own LSW docker image.

# Motivation

I need a way to deploy software, and while it is a rather quick, easy, and enjoyable task that can be automated with a single Makefile, build.sh, and a Dockerfile on Linux it doesn't work like this on the more popular operating systems. Microsoft has recognized this issue, and instead of making life easier for people who port to Windows, they have decided to.. ship Linux. As we can see many projects (docker included) dropped support (to some extent) for Windows, and simply told end users to install WSL2. 

This clearly doesn't solve the issue on my end, I want the opposite. I want **Windows Containers** on Linux.

So without unnecessary bloat in the post let's get started.

## Getting started

This documents how the current image was created. Ideally, all of this will be automated using an unattended XML file, but a day only has 24 hours, and I'm not feeling like investing my time into Windows more than I need. That being said, I'll most likely do that at some point, and automate everything from the ground up with a single Dockerfile, that will do all the tasks. This is however a bit harder. Consider this LSW v1 and the proper thing will be a v2.

> **Important note:** I do not personalize the installation.
> It should be as default as it can be + Windows, I do not run
> any extra software, don't change wallpaper, screen resolution
> This will create a mess. Instead do only what is required
> p.s. Debloating is personalizing
> p.p.s. So is activating the copy, I don't support piracy <!-- in this project --> - 
> Windows should be activated, but this is the responsibility of
> the image user.

I've used a Windows 10 LTSC 2021 iso image, so I'll name my **disk** image accordingly.

### Create a disk image

64G is the recommended size by me, it will allow to embedding of relatively large images (NOTE: this is the upper limit of the partition, it can be resized on demand, 
this is only a good default value according to me, no solid math behind it. If you have a legitimate reason to use a larger partition please open an issue [GitHub](https://github.com/mrcyjanek/LSW) or [Gitea](https://git.mrcyjanek.net/mrcyjanek/LSW) - I'll be very happy to help with it).

```bash
$ qemu-img create \
    -f qcow2 \
    windows10iot_enterprise_ltsc2021.img 64G
```

### Boot the installation 

> TIP: do that offline to skip the m$ account login

> TIP: Note the `virtio-win-0.1.240.iso`. It is required to use `if=virtio`
> option to the `-drive` argument.
> This is not a hard requirement but we should catch performance wherever
> we can. During installation click `Load driver` when you won't see any 
> destination disk when trying to install.

```bash
$ qemu-system-x86_64 \
    -enable-kvm \
    -drive file=isos/virtio-win-0.1.240.iso,media=cdrom \
    -drive file=windows10iot_enterprise_ltsc2021.img,if=virtio \
    --drive file=isos/windows_10_iot_enterprise_ltsc_2021_257ad90f.iso,media=cdrom
    -net user,hostfwd=tcp::2222-:22 -net nic,model=rtl8139 \
    -m 8G
```

- Create a user named 'user' with an empty password
- Make sure that this 'user' is a system administrator.

### After installation power off the machine and boot it without the isos

It is to ensure that we can boot properly without the .iso in.

```bash
$ qemu-system-x86_64 \
    -enable-kvm \
    -drive file=windows10iot_enterprise_ltsc2021.img,if=virtio \
    -net user,hostfwd=tcp::2222-:22 \
    -net nic,model=rtl8139 \
    -m 8G
```

### Apply Windows update

Yes, we need to. Core image should be the most Windows thing that just works, keep in mind that this will be emulated, so if we skip some tasks then everybody who will use this image will need to do them multiple times. 

### Disable Windows fast startup

- Navigate to the Control Panel and select 'Power Options.'
- Choose 'Choose what the power buttons do.'
- Click on 'Change settings that are currently unavailable.'
- Under 'Shutdown settings,' uncheck the 'Turn on fast startup' box.

Let it run for a couple of hours, so it will do all the things that it needs to... I've seen CPU spikes (or rather heard the fan spin) randomly for the first couple of hours, despite nothing being done, it was just sitting (connected to the internet) on the desktop. Since we want this solution to be at least somewhat fast, it is better to do all the CPU-intensive tasks before shipping the image.

### Enable SSH Server on port 22

- Go to settings -> Apps -> Apps & features -> Optional features
- Add a feature
- Search for `Windows OpenSSH server`
- Click install
- Open `services` app
- Search for `OpenSSH SSH Server`
- Right click -> properties
- Startup type: set to `Automatic` (do **NOT** set it as ~~Automatic (delayed)~~ - this will add a timeout that doesn't wait for any features, it just delays by some time).
- Then click Apply and Ok
- Right click -> start
- Open Notepad as administrator
- Open file `%programdata%\ssh\sshd_config` (you can type `%programdata%`) in the top bar in explorer.exe
- Make sure that in the left bottom corner, you have selected All Files instead of Text Documents (*.txt)
- If the directory is empty go back to step 7.3 or try to reboot it. It's windows.
- Edit sshd_config to contain the following lines: Keep in mind that given lines may be already present make sure to find them first and change no to yes (and uncomment by removing #)
```plain
PasswordAuthentication yes
PermitEmptyPasswords yes
```
- Open `regedit.msc`
- In the tree on the left navigate to `Computer\HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Control\Lsa\`
- On the right side look for: `LimitBlankPasswordUse` and set it to `0`
- Restart sshd (`cmd.exe` as admin -> `net stop sshd`; `net start sshd`)
- **ON HOST** check if the ssh is working (`ssh user@127.0.0.1 -p2222`)
- Congratulations, it should be working now. This image is almost ready ~~to ship~~ to use personally.

### Chore tasks

- Open `explorer.exe`
- Right-click on Disk C: -> properties
- `Disk Cleanup`
- `Clean up system files` (this cleaned 1.3 GB in my case so from 21.2 used we went to.. 16.4.. which isn't that great, but is better than predicted.)
- Then uncheck `Allow files on this drive to have contents indexed in addition to file properties`. This takes [~10% extra space](https://support.microsoft.com/en-us/windows/search-indexing-in-windows-10-faq-da061c83-af6b-095c-0f7a-4dfecda4d15a) (not to mention IO and CPU time). (In my case 16.4 -> 16.3)
- Apply, click ignore all, and let it run for.. some time. At least until it finishes.
- Reboot the host, let it idle for some time, and power it off one last time outside of docker.

### Preparing docker image

There is a helper script that generates a Dockerfile for the given disk image.

```bash
$ ./docker_prep.sh \
    windows10iot_enterprise_ltsc2021.img \
    10iot_enterprise_ltsc2021
```

> NOTE: Don't have `docker_prep.sh`? Grab it from [GitHub](https://github.com/mrcyjanek/LSW) or [Gitea](https://git.mrcyjanek.net/mrcyjanek/LSW).

It will create a directory named `10iot_enterprise_ltsc2021`,

Then we can just `cd 10iot_enterprise_ltsc2021` and... you probably hoped that we could just build it. No, we need to allow docker to run in insecure mode (this is needed to access `/dev/kvm` and to speed up the disk optimization process). [reference](https://github.com/moby/buildkit/blob/master/frontend/dockerfile/docs/reference.md#run---securityinsecure)

While this isn't necessarily a **must** reducing the build time by >90% was quite important for me (800 seconds vs over 10000 without the option), especially because I've rebuilt it like.. many times.

```bash
# docker buildx create --driver-opt image=moby/buildkit:master  \
    --use --name insecure-builder \
    --buildkitd-flags '--allow-insecure-entitlement security.insecure'
# docker buildx use insecure-builder
```

And build our docker image, using the given builder

```bash
# docker buildx build \
    --allow security.insecure . \
    -f Dockerfile \
    -t windows:10iot_enterprise_ltsc2021 \
    --load
```

On my end, a resulting image just got produced and was only `8.93GB` in size

> NOTE: You may see different sizes across the article, and different checksum as well.
> This is because I didn't get this right first try...

```bash
$ docker images | grep 10iot_enterprise_ltsc2021
windows                10iot_enterprise_ltsc2021           480a40c821107   2 hours ago         8.93GB
```

Keep in mind that the Windows disk size may be still huge, and you could feel the urge to _"delete that thing that nobody uses"_ and could be safely removed, I get it, it is really tempting to debloat this thing - but don't do it. Not here at least. This falls into the "customization" part that shouldn't be done at all.

And that's it from the preparation side. The image is now fully usable. From now on we shouldn't apply manual patches, we can do everything in the form of `Dockerfiles` (that includes debloating the Windows image to reduce its size). Don't believe me? See for yourself

```bash
$ uname -s -r
Linux 6.1.68
$ docker run --device=/dev/kvm --rm -it windows:10iot_enterprise_ltsc2021 systeminfo
Host Name:                 DESKTOP-GLPCL5M
OS Name:                   Microsoft Windows 10 IoT Enterprise LTSC
OS Version:                10.0.19044 N/A Build 19044
OS Manufacturer:           Microsoft Corporation
OS Configuration:          Standalone Workstation
OS Build Type:             Multiprocessor Free
{...}
```

> NOTE: In interactive mode, the image runs `cmd.exe`
> But when arguments are given it is executed by `powershell --command "$@"`

So now we have a ready Windows image, what should we do next? Create flavors of the images so they will be useful. Despite Windows being already 8.93GB (compressed) in size, there is simply nothing that can be used in a `docker` way. Not even a C compiler, nothing.

## Flavors (customization)

My main goal is to compile Flutter apps, so let's start with it. Let's create a `Dockerfile` that will grab and install all of the flutter things.

>NOTE: many things are still missing, for example, things such as ENV, ARG, etc.. need extra work,
> I'll try to increase the compatibility so it will feel as native as it can, but this is a task
> for the future.

We can import our images easily

```dockerfile
FROM windows:10iot_enterprise_ltsc2021
```

Then let's put our required files in the image. Everything that will be put in the `/cdrom` directory will be accessible in a read-only CD (disk `D:\`). For large amounts of data, it may be desired to put everything directly in `/cdrom.iso`.

```dockerfile
ADD https://storage.googleapis.com/flutter_infra_release/releases/stable/windows/flutter_windows_3.16.4-stable.zip /cdrom/flutter.zip
```
Now that we have the `flutter.zip` in let's install it.

```dockerfile
RUN Expand-Archive D:\\flutter.zip -DestinationPath C:\\ \
    && setx /M PATH "$($env:path);C:\\flutter\\bin"
```

And that would be it. Except for the fact that we need `Microsoft Visual Studio for Desktop C++ Development` or something like this. As far as I know, there is no easy way of installing it using headless windows, so we will use `choco` instead to install all the required software.

I truly hate how verbose powershell commands are btw.

```dockerfile
FROM windows:10iot_enterprise_ltsc2021

RUN Set-ExecutionPolicy Bypass -Scope Process -Force \
    ; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072 \
    ; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))

# <reboot takes place somewhere in here>

RUN choco install -y visualstudio2022community \
    && choco install -y visualstudio2022-workload-nativedesktop \
    && choco install -y git \
    && choco install -y flutter
```

To do the most Windows thing that can be done (reboot after every little modification is made to the system), we.. don't need to do anything specific. When we `RUN` something we do the full cycle:

1. Boot up windows
2. Wait for the SSH connection to be available
3. Execute interactive shell / run the given command
4. Gently shutdown the machine (`shutdown -s -t 1`)
5. Wait for it to shut down, and do a quick cleanup.

While this is a disadvantage in most cases (it would be nice to enter straight into running VM), we are really close to how Windows works, we also ensure that the environment is relatively clean (i.e. no bugs occurring because of huge uptime), and works "the windows way".

So to do the recommended `reboot` after installing `choco`latey we can simply add another `RUN` command, instead of `&&`ing more commands.

To _quickly_ verify that this is correct we can build the flavor image:

```bash
$ docker build . -f Dockerfile -t windows:10iot_enterprise_ltsc2021-flutter
```

And I've sadly noticed that it is far (very far) from quickly. It turns out that `docker buildx` can't use images not from the registry. So we can't do the magic `insecure` trick here. My current solution is to go to sleep and wake up when it finishes executing. This is acceptable, but if somebody could host a registry of the images that would be cool.

There is also a little.. obstacle:

```bash
$ docker images | grep 10iot_enterprise_ltsc2021
windows                10iot_enterprise_ltsc2021-flutter   875e1df23c2f   25 minutes ago   17GB
windows                10iot_enterprise_ltsc2021           6a3a12dafb6d   48 minutes ago   9.77GB
```

Most likely compression is gone after writing to the disk so we are back to our normal disk size. This may be fine (even desired, I guess that it may take some time to decompress it).

But we most likely do not want to ship 17GB to the registry - especially when we can ship 9.77GB as an alternative, and some (native, not emulated) CPU usage doesn't hurt us that much (I think???).

This step is already done once - when we have prepared the initial image. So we may redo it one more time. There is actually a really easy way to achieve it.

```dockerfile
RUN '--disk-optimize' # note the quotes, otherwise docker will try to interact with this flag.
```

This won't start the Windows machine, instead, it will fill the disk with 0s (to wipe empty space) and compress it. This is a time-consuming task, you may want to not do it at all, or do it only when moving the images around.

> I want you, my dear reader to stop here for 10 seconds, and then admire the results.

So what are the results of --disk-optimize?

> :)

```bash
$ docker images | grep 10iot_enterprise_ltsc2021
windows                10iot_enterprise_ltsc2021   0d41658d4c94   43 years ago   28.3GB
```

So we have successfully gone from **9.77GB** to just **28.3GB** which has made us save **-289.66%** of disk space. Let's ignore this for now and learn about layers.

### Layers

Docker images work just like CD drives. That didn't help? No problem. Let's imagine the following scenario:

- You have put 5 songs on a CD disk.
- Then you have removed one that you didn't like
- And put 5 more songs.

So how many songs are on the CD? The answer is 10. The one song that you have removed is not actually removed, there is an extra metadata written to the filesystem 
that tells your computer that it got removed, but the 0s and 1s are still out there - this is caused by the limitation of CD disks that only were single-write disks.

> I, the post author, am aware of ReWritable CDs. But they don't fit well into my definition so let's pretend that they didn't exist.

The same issue occurs on our docker container, we have reduced the 18.53GB image down to 9.77GB, but since they were not made in the same layer both of them are still in the layers. To quickly verify that we are right we can use the following Dockerfile:

```dockerfile
FROM debian:stable

RUN dd if=/dev/urandom of=/blob.bin bs=1000000 count=128
RUN rm /blob.bin
RUN dd if=/dev/urandom of=/blob.bin bs=1000000 count=128
RUN dd if=/dev/urandom of=/blob.bin bs=1000000 count=128
```
To my best knowledge, we should weigh around whatever debian:stable weights + 128MB.

```bash
$ docker build . -f Dockerfile -t layers:latest
$ docker images | grep -e layers -e debian
layers              latest   6f9bdec3a5c7   2 minutes ago   501MB
debian              stable   7404e946945f   41 hours ago    117MB
```

But is this the case? Well, we are clearly far off the 117MB + 128MB, in fact, we have exceeded the limit 3 times (+117MB from debian:stable).

```bash
$ docker history layers:latest       
IMAGE          CREATED         CREATED BY                                      SIZE      COMMENT
6f9bdec3a5c7   2 seconds ago   RUN /bin/sh -c dd if=/dev/urandom of=/blob.b…   128MB     buildkit.dockerfile.v0
<missing>      4 seconds ago   RUN /bin/sh -c dd if=/dev/urandom of=/blob.b…   128MB     buildkit.dockerfile.v0
<missing>      5 seconds ago   RUN /bin/sh -c rm /blob.bin # buildkit          0B        buildkit.dockerfile.v0
<missing>      5 seconds ago   RUN /bin/sh -c dd if=/dev/urandom of=/blob.b…   128MB     buildkit.dockerfile.v0
<missing>      41 hours ago    /bin/sh -c #(nop)  CMD ["bash"]                 0B        
<missing>      41 hours ago    /bin/sh -c #(nop) ADD file:974e64d5811025cce…   117MB     
```

As we can see `SIZE` is 0B when we are removing the blob (which means no change in file size.), while each subsequent build has 128M of size used.

Back to the analogy - so what can we do to keep only 9 songs on the disk?

We could of course spend more time thinking about what music we want to pick and then carefully put it on the CD, but there is no joy in that. So instead let's buy a new CD disk and burn the image again.

Let's do this with our test Dockerfile:

```dockerfile
FROM debian:stable as stage00

RUN dd if=/dev/urandom of=/blob.bin bs=1000000 count=128
RUN rm /blob.bin
RUN dd if=/dev/urandom of=/blob.bin bs=1000000 count=128
RUN dd if=/dev/urandom of=/blob.bin bs=1000000 count=128

FROM scratch
COPY --from=stage00 / /
```

What are the results here?

```bash
$ docker images | grep -e layers -e debian
layers              latest    6b0c8f32270e   8 seconds ago    244MB
debian              stable    7404e946945f   41 hours ago     117MB
$ docker history layers:latest
IMAGE          CREATED         CREATED BY            SIZE      COMMENT
6b0c8f32270e   2 minutes ago   COPY / / # buildkit   244MB     buildkit.dockerfile.v0
```

Alright, so since this is proven to work by pure science, let's implement it directly in our Dockerfile.

Let's add to the ~~first~~ second line of our Dockerfile name since this will no longer be the image, it will be a builder.

```dockerfile
FROM debian:stable as stage00
```

And then, later in the Dockerfile let's just copy the final version onto its own layer

```dockerfile
FROM scratch
COPY --from=stage00 / /

ENV QEMU_KVM_FLAGS=""
ENV QEMU_KVM_FLAGS=" -enable-kvm -cpu host"
ENV QEMU_RAM_LIMIT=4G
ENV FORWARD_PORTS="8080"

ENTRYPOINT [ "/entrypoint.sh" ]
SHELL [ "/entrypoint.sh" ]
```

What are the results? I'd say that they are great, we are back at the "expected" size. We can probably trim .5 GB if we compile dependencies that we need instead of using Debian, and provide a debloated flavor of the image, but this is enough windows for me at least for now.

```bash
windows             10iot_enterprise_ltsc2021   701ae56325de   43 years ago   9.94GB
```

## TLDR

Everything went as I expected, except for the docker side of things, I've read all over the place how limited is docker, that a layer can be max 10GB in size, etc. But I've tried anyway and I didn't need to make any changes to the logic (such as splitting the disk image).

For a usage instruction go to my [GitHub](https://github.com/mrcyjanek/LSW) or [Gitea](https://git.mrcyjanek.net/mrcyjanek/LSW).

### Features

- Windows containers on Linux
- KVM support
- no-KVM support
- COPY files into the image (`/cdrom`)
- COPY files outside of the image (`C:\outputs` is copied to `/outputs`)
- Provide a CLI that can be used in `docker run` and `docker build`

## Future plans

I expect to make this project more standalone (i.e. provide a single Dockerfile that will do everything), so no manual setup will be required at all. And I also plan 
to figure something out for MacOS... They don't offer docker containers even for their own OS and are super hostile against using their OS on non-official machines. 
Not to mention the fact that they will most likely abandon x64 architecture altogether soon.

> p.s. I've almost frozen finishing off this post, I really do not enjoy traveling by Koleje Śląskie in winter...
> Idling temperatures are around ~45°C on my laptop.

```bash
$ sensors
coretemp-isa-0000
Adapter: ISA adapter
Package id 0:  +31.0°C 
```