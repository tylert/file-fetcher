#!/usr/bin/env python3

from fetcher import fetch_file


if __name__ == '__main__':
    fetch_file('https://os.archlinuxarm.org/os/ArchLinuxARM-rpi-3-latest.tar.gz.sig',
               'ArchLinuxARM-rpi-3-latest.tar.gz.sig')
    fetch_file('https://os.archlinuxarm.org/os/ArchLinuxARM-rpi-3-latest.tar.gz',
               'ArchLinuxARM-rpi-3-latest.tar.gz')
