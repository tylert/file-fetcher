#!/usr/bin/env python

import os

import requests
from bs4 import BeautifulSoup


def find_debian_version(release):
    '''Fetch a complete Debian release.'''

    r = requests.get(f'https://cdimage.debian.org/cdimage/{release}/amd64/iso-cd')
    r.raise_for_status()
    s = BeautifulSoup(r.text, 'html.parser')

    for link in s.find_all('a'):
        if 'iso' in link.text and 'netinst' in link.text and 'mac' not in link.text and 'edu' not in link.text:
            iso = link.text

    print(f'{iso}')


def main():
    '''Main function.'''

    find_debian_version('release/current')
    find_debian_version('archive/latest-oldstable')


if __name__ == '__main__':
    main()
