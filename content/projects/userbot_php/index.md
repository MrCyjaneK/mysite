+++
Area = "projects"
CrumbParent = "projects"
Layout = "page"
Tags = [ "PHP", "Telegram", "FOSS" ]
PostTitle = "userbot.php - Badly written Telegram Userbot using Madeline Proto"
+++

# userbot.php
Telegram userbot in php, Thank's for creating MadelineProto @danog

# How to get started?

```
# Install php
sudo apt install php7.3 php7.3-gd php7.3-mbstring php7.3-fileinfo php7.3-xml
# Clone repository
git clone https://github.com/MrCyjaneK/userbot.php && cd userbot.php
# Login
php utils/login.php
# Clean updates (optional)
#If you want to speed up loading process keep command below running for few minutes
timeout 300 php utils/cleanupdates.php
# Start userbot
php index.php
```

To update use: `git pull`
