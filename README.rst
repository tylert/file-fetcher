file-fetcher
============


::

    go run ${PRODUCT}/main.go | aria2c -i -

* Arch Linux ARM for rpi-aarch64, odroid-xu3, odroid-c2
* https://endeavouros.com/latest-release  the github.com one
* https://releases.hashicorp.com/index.html
* https://releases.hashicorp.com/index.json
* https://pikvm.org
* https://libreelec.tv/downloads/raspberry/
* https://www.pistar.uk/downloads/

::

    curl https://api.github.com/repos/keepassxreboot/keepassxc/releases/latest | jq '.assets[].browser_download_url'
    curl https://api.github.com/repos/keepassxreboot/keepassxc/releases/latest | jq '.tarball_url'
    curl https://api.github.com/repos/keepassxreboot/keepassxc/releases/latest | jq '.tag_name'

    curl https://api.github.com/repos/hairyhenderson/gomplate/releases/latest | jq '.assets[].browser_download_url'
    curl https://api.github.com/repos/hairyhenderson/gomplate/releases/latest | jq '.tarball_url'
    curl https://api.github.com/repos/hairyhenderson/gomplate/releases/latest | jq '.tag_name'

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

* https://github.com/aria2/aria2
* https://aria2.github.io
* https://wiki.archlinux.org/title/Aria2
* https://github.com/ziahamza/webui-aria2


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
* https://github.com/wahyd4/aria2-ariang-docker
* https://github.com/openthings/kubernetes-tools/blob/master/aria/aria2-service.yaml
