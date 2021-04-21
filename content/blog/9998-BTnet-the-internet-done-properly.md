+++
Title = "[2021.04.21] BTnet - The internet done properly."
Layout = "page"
Collection = "Blog"
CrumbParent = "blog"
Tags = [ "BTnet" ]
+++

# [2021.04.21] BTnet - The internet done properly.

No matter who you are, you have probably seen the signs of the time, either by your president twitter account getting banned or by your favourite youtuber getting demonetised. Or have you ever seen a 404 page instead of the controversial article you have bookmarked to read later? BTnet will fix it.

# BTnet - Honest internet.

With BTnet your site is safe, from everything. From things getting changed or deleted, and from paying for the servers.

BTnet use commonly used protocol named BitTorrent, well... to be honest BTnet is just a proxy between your web browser and the BitTorrent network. Possibilities are endless, as long as no server-side code processing is needed, so blogs, video hosting, news, offline copies of websites will all work under BTnet.

## How does it work?

BTnet runs a http proxy on `127.0.0.1:8080`, when you go to a btnet address like this one: http://6626ae3c23b19bf4ba7d17c765be2c83935d51a3.btnet/, btnet looks for the hash (662...1a3) over DHT and BitTorrent trackers, then it download requested files, and serves it right to your browser.

## That's it?

Yes. For now. I have plans on implementing a simple federated database, that will act as the missing thing - the server to handle comments, likes and user actions in general. Also every file is signed, so it's easy to check if newer version of the same site got deployed by the same person, or if we have impostor.

Most of that is not (yet) implemented, but that will change soon, as this project is now my priority.

If you want, go to https://git.mrcyjanek.net/mrcyjanek/btnet, to check the source code and setup instructions.