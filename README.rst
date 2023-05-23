file-fetcher
============


::

    go mod init bla
    go mod tidy

    go run nncp.go
    go run debian.go


::

    curl https://api.github.com/repos/keepassxreboot/keepassxc/releases/latest | jq '.assets[].browser_download_url'
    curl https://api.github.com/repos/keepassxreboot/keepassxc/releases/latest | jq '.tarball_url'
    curl https://api.github.com/repos/keepassxreboot/keepassxc/releases/latest | jq '.tag_name'

    curl https://api.github.com/repos/ventoy/Ventoy/releases/latest | jq '.assets[].browser_download_url'
    curl https://api.github.com/repos/ventoy/Ventoy/releases/latest | jq '.tarball_url'
    curl https://api.github.com/repos/ventoy/Ventoy/releases/latest | jq '.tag_name'

    curl https://api.github.com/repos/Ultimaker/Cura/releases/latest | jq '.assets[].browser_download_url'
    curl https://api.github.com/repos/Ultimaker/Cura/releases/latest | jq '.tarball_url'
    curl https://api.github.com/repos/Ultimaker/Cura/releases/latest | jq '.tagname'

    curl https://api.github.com/repos/gqrx-sdr/gqrx/releases/latest | jq '.assets[].browser_download_url'
    curl https://api.github.com/repos/gqrx-sdr/gqrx/releases/latest | jq '.tarball_url'
    curl https://api.github.com/repos/gqrx-sdr/gqrx/releases/latest | jq '.tag_name'

    curl https://api.github.com/repos/hairyhenderson/gomplate/releases/latest | jq '.assets[].browser_download_url'
    curl https://api.github.com/repos/hairyhenderson/gomplate/releases/latest | jq '.tarball_url'
    curl https://api.github.com/repos/hairyhenderson/gomplate/releases/latest | jq '.tag_name'

    curl https://api.github.com/repos/alexellis/k3sup/releases/latest | jq '.assets[].browser_download_url'
    curl https://api.github.com/repos/alexellis/k3sup/releases/latest | jq '.tarball_url'
    curl https://api.github.com/repos/alexellis/k3sup/releases/latest | jq '.tag_name'

    curl https://api.github.com/repos/jsiebens/hashi-up/releases/latest | jq '.assets[].browser_download_url'
    curl https://api.github.com/repos/jsiebens/hashi-up/releases/latest | jq '.tarball_url'
    curl https://api.github.com/repos/jsiebens/hashi-up/releases/latest | jq '.tag_name'

    curl https://api.github.com/repos/dalefarnsworth-dmr/codeplug/tags | jq '.[].tarball_url' | head -1
    curl https://api.github.com/repos/dalefarnsworth-dmr/debug/tags | jq '.[].tarball_url' | head -1
    curl https://api.github.com/repos/dalefarnsworth-dmr/dfu/tags | jq '.[].tarball_url' | head -1
    curl https://api.github.com/repos/dalefarnsworth-dmr/dmrRadio/tags | jq '.[].tarball_url' | head -1
    curl https://api.github.com/repos/dalefarnsworth-dmr/editcp/tags | jq '.[].tarball_url' | head -1
    curl https://api.github.com/repos/dalefarnsworth-dmr/docCodeplug/tags | jq '.[].tarball_url' | head -1
    curl https://api.github.com/repos/dalefarnsworth-dmr/docker/tags | jq '.[].tarball_url' | head -1
    curl https://api.github.com/repos/dalefarnsworth-dmr/docs/tags | jq '.[].tarball_url' | head -1
    curl https://api.github.com/repos/dalefarnsworth-dmr/editcp/tags | jq '.[].tarball_url' | head -1
    curl https://api.github.com/repos/dalefarnsworth-dmr/genCodeplugInfo/tags | jq '.[].tarball_url' | head -1
    curl https://api.github.com/repos/dalefarnsworth-dmr/genFileData/tags | jq '.[].tarball_url' | head -1
    curl https://api.github.com/repos/dalefarnsworth-dmr/stdfu/tags | jq '.[].tarball_url' | head -1
    curl https://api.github.com/repos/dalefarnsworth-dmr/ui/tags | jq '.[].tarball_url' | head -1
    curl https://api.github.com/repos/dalefarnsworth-dmr/userdb/tags | jq '.[].tarball_url' | head -1
    # https://www.farnsworth.org/dale/codeplug/dmrRadio/downloads/
    # https://www.farnsworth.org/dale/codeplug/editcp/downloads/
    # https://www.farnsworth.org/dale/codeplug/editcp/


TODO
----

* https://github.com/siku2/arigo
* https://github.com/cavaliercoder/grab
* https://github.com/melbahja/got
* https://github.com/akshaykhairmode/summon
* https://github.com/codingonHP/file_downloader
* https://github.com/huydx/hget
* https://github.com/hashicorp/go-getter
* https://stackoverflow.com/questions/11692860/how-can-i-efficiently-download-a-large-file-using-go
* https://progolang.com/how-to-download-files-in-go/
* https://likegeeks.com/downloading-files-using-python/
* http://ariang.mayswind.net/
* https://github.com/P3TERX/Aria2-Pro-Docker
* https://github.com/k8s-at-home/charts/tree/master/charts/stable/aria2
* https://aria2.github.io/manual/en/html/aria2c.html#aria2.addUri
* https://github.com/wahyd4/aria2-ariang-docker
* https://github.com/openthings/kubernetes-tools/blob/master/aria/aria2-service.yaml

::

    gomplate -d dl=versions.yaml -f source/foo.aria2.gomplate > template/foo.aria2.list
    gomplate --datasource dl=versions.yaml --file source/bar.aria2.gomplate > template/bar.aria2.list

    aria2c -d foo --conf-path=aria2.conf -i template/foo.aria2.list
    aria2c --dir=bar --conf-path=aria2.conf --input-file=template/bar.aria2.list

* https://endeavouros.com/latest-release  (the github.com one)
* https://mirror.xenyth.net/archlinux/iso/latest  (x86_64, rpi-aarch64, odroid-xu3, odroid-c2)
* https://releases.hashicorp.com/index.html
* https://releases.hashicorp.com/index.json
* https://pikvm.org
* https://libreelec.tv/downloads/raspberry/
* https://www.pistar.uk/downloads/
* http://releases.ubuntu.com/
* http://cdimage.ubuntu.com/
* https://ubuntu.com/download/raspberry-pi
