file-fetcher
============


TODO
----

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
