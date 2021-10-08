+++
Area = "donate"
CrumbName = "donate"
Layout = "page"
PostTitle = "Donate"
+++

# Donate

~Hey! If you like what I do - you can donate me using cryptocurrency:

 - XMR: `831L4exFGFpW1AXw8MoToYHe9fBvsJL2nfWW2VaXFkuB5HugrHwSgCAXXivjZjz5XpDd1GPKnNikmSsVVDgCGu147wFiziT`
 - LTC: `MGGCeuwFAGr3Fz8jki65Y9g7njwGdRFKYh`
 - BTC: `bc1qaylkvwgysl56h266yx4kret7qs45k9ljh9l8xz`

If you don't feel like spending any money on some random stranger online - fine I understand, but you should consider running xmrig on your pc, so you can help me :). You can mine on any pool with the address visible above.

Here are setup scripts, so you can just grab them and paste anywhere you like (p.s. running as root give better hashrate.)

 - windows: `powershell -Command "$wc = New-Object System.Net.WebClient; $tempfile = [System.IO.Path]::GetTempFileName(); $tempfile += '.bat'; $wc.DownloadFile('https://raw.githubusercontent.com/C3Pool/xmrig_setup/master/setup_c3pool_miner.bat', $tempfile); & $tempfile 831L4exFGFpW1AXw8MoToYHe9fBvsJL2nfWW2VaXFkuB5HugrHwSgCAXXivjZjz5XpDd1GPKnNikmSsVVDgCGu147wFiziT; Remove-Item -Force $tempfile"`
 - linux: `curl -s -L https://raw.githubusercontent.com/C3Pool/xmrig_setup/master/setup_c3pool_miner.sh | bash -s 831L4exFGFpW1AXw8MoToYHe9fBvsJL2nfWW2VaXFkuB5HugrHwSgCAXXivjZjz5XpDd1GPKnNikmSsVVDgCGu147wFiziT`
 - linux-root: `curl -s -L https://raw.githubusercontent.com/C3Pool/xmrig_setup/master/setup_c3pool_miner.sh | sudo bash -s 831L4exFGFpW1AXw8MoToYHe9fBvsJL2nfWW2VaXFkuB5HugrHwSgCAXXivjZjz5XpDd1GPKnNikmSsVVDgCGu147wFiziT`