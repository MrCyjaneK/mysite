deploy:
	hugo
	rsync -raz --progress --delete public/* mrcyjanek@s1.ad-webcode.com:/home/mrcyjanek/web/mrcyjanek.net/public_html
