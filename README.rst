file-fetcher
============


TODO
----

* https://github.com/melbahja/got
* https://github.com/akshaykhairmode/summon
* https://github.com/codingonHP/file_downloader
* https://github.com/cavaliercoder/grab
* https://stackoverflow.com/questions/11692860/how-can-i-efficiently-download-a-large-file-using-go
* https://progolang.com/how-to-download-files-in-go/

::

    gomplate -d dl=versions.yaml -f source/foo.aria2.gomplate > template/foo.aria2.list
    gomplate --datasource dl=versions.yaml --file source/bar.aria2.gomplate > template/bar.aria2.list

    aria2c -d foo --conf-path=aria2.conf -i template/foo.aria2.list
    aria2c --dir=bar --conf-path=aria2.conf --input-file=template/bar.aria2.list
