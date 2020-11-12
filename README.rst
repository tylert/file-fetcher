file-fetcher
============


TODO
----

* automate verification of checksums and signatures

::

    gomplate -d dl=versions.yaml -f source/foo.aria2.gomplate > template/foo.aria2.list
    gomplate --datasource dl=versions.yaml --file source/bar.aria2.gomplate > template/bar.aria2.list

    aria2c -d foo --conf-path=aria2.conf -i template/foo.aria2.list
    aria2c --dir=bar --conf-path=aria2.conf --input-file=template/bar.aria2.list
