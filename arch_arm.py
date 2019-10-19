#!/usr/bin/env python

import datetime

from fetcher import fetch_file


def main():
    '''Main function.'''

    date = datetime.datetime.today().strftime('%Y-%m-%d')

    # Odroid HC2/HC1/XU4/XU3
    fetch_file('http://os.archlinuxarm.org/os/ArchLinuxARM-odroid-xu3-latest.tar.gz.sig',
               'ArchLinuxARM-odroid-xu3-{}.tar.gz.sig'.format(date))
    fetch_file('http://os.archlinuxarm.org/os/ArchLinuxARM-odroid-xu3-latest.tar.gz.md5',
               'ArchLinuxARM-odroid-xu3-{}.tar.gz.md5'.format(date))
    fetch_file('http://os.archlinuxarm.org/os/ArchLinuxARM-odroid-xu3-latest.tar.gz',
               'ArchLinuxARM-odroid-xu3-{}.tar.gz'.format(date))

    # Raspberry Pi 4B
    fetch_file('http://os.archlinuxarm.org/os/ArchLinuxARM-rpi-4-latest.tar.gz.sig',
               'ArchLinuxARM-rpi-4-{}.tar.gz.sig'.format(date))
    fetch_file('http://os.archlinuxarm.org/os/ArchLinuxARM-rpi-4-latest.tar.gz.md5',
               'ArchLinuxARM-rpi-4-{}.tar.gz.md5'.format(date))
    fetch_file('http://os.archlinuxarm.org/os/ArchLinuxARM-rpi-4-latest.tar.gz',
               'ArchLinuxARM-rpi-4-{}.tar.gz'.format(date))

    # Raspberry Pi 3B
    fetch_file('http://os.archlinuxarm.org/os/ArchLinuxARM-rpi-3-latest.tar.gz.sig',
               'ArchLinuxARM-rpi-3-{}.tar.gz.sig'.format(date))
    fetch_file('http://os.archlinuxarm.org/os/ArchLinuxARM-rpi-3-latest.tar.gz.md5',
               'ArchLinuxARM-rpi-3-{}.tar.gz.md5'.format(date))
    fetch_file('http://os.archlinuxarm.org/os/ArchLinuxARM-rpi-3-latest.tar.gz',
               'ArchLinuxARM-rpi-3-{}.tar.gz'.format(date))


if __name__ == '__main__':
    main()
