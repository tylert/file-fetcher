#!/usr/bin/env python3

import datetime

from fetcher import fetch_file


def main():
    '''Main function.'''

    date = datetime.datetime.today().strftime('%Y-%m-%d')

    fetch_file('http://os.archlinuxarm.org/os/ArchLinuxARM-rpi-3-latest.tar.gz.sig',
               'ArchLinuxARM-rpi-3-{}.tar.gz.sig'.format(date))
    fetch_file('http://os.archlinuxarm.org/os/ArchLinuxARM-rpi-3-latest.tar.gz.md5',
               'ArchLinuxARM-rpi-3-{}.tar.gz.md5'.format(date))
    fetch_file('http://os.archlinuxarm.org/os/ArchLinuxARM-rpi-3-latest.tar.gz',
               'ArchLinuxARM-rpi-3-{}.tar.gz'.format(date))


if __name__ == '__main__':
    main()
