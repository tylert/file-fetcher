File Fetcher
============


TODO
----

* stderr and stdout maybe???
* add in common aria2 config entries???


Running It
----------

::

    pushd aria2
    go run ${PRODUCT}/main.go | aria2c -i -
    popd

* https://github.com/aria2/aria2
* https://aria2.github.io
* https://wiki.archlinux.org/title/Aria2
* http://ariang.mayswind.net/
* https://github.com/ziahamza/webui-aria2


Next Items
----------

* Arch Linux ARM for odroid-xu3, odroid-c2???
* https://releases.hashicorp.com/index.html
* https://releases.hashicorp.com/index.json

::

    curl https://api.github.com/repos/hairyhenderson/gomplate/releases/latest | jq '.assets[].browser_download_url'
    curl https://api.github.com/repos/hairyhenderson/gomplate/releases/latest | jq '.tarball_url'
    curl https://api.github.com/repos/hairyhenderson/gomplate/releases/latest | jq '.tag_name'

    curl https://api.github.com/repos/jsiebens/hashi-up/releases/latest | jq '.assets[].browser_download_url'
    curl https://api.github.com/repos/jsiebens/hashi-up/releases/latest | jq '.tarball_url'
    curl https://api.github.com/repos/jsiebens/hashi-up/releases/latest | jq '.tag_name'


Other
-----

* https://github.com/siku2/arigo
* https://github.com/P3TERX/Aria2-Pro-Docker
* https://github.com/wahyd4/aria2-ariang-docker
* https://github.com/openthings/kubernetes-tools/blob/master/aria/aria2-service.yaml
* https://github.com/k8s-at-home/charts/tree/master/charts/stable/aria2
* https://github.com/cavaliercoder/grab
* https://github.com/melbahja/got
* https://github.com/akshaykhairmode/summon
* https://github.com/codingonHP/file_downloader
* https://github.com/huydx/hget
* https://github.com/hashicorp/go-getter
* https://stackoverflow.com/questions/11692860/how-can-i-efficiently-download-a-large-file-using-go
* https://progolang.com/how-to-download-files-in-go/
* https://github.com/nfx/go-htmltable  for scraping more stubborn table data???
* https://gist.github.com/salmoni/27aee5bb0d26536391aabe7f13a72494  more complex goquery example
