#!/usr/bin/env python

import os

import requests
from bs4 import BeautifulSoup


def find_debian_version(release):
    '''Fetch a complete Debian release.'''

    r = requests.get('https://cdimage.debian.org/cdimage/{}/amd64/iso-cd'.format(release))
    r.raise_for_status()
    s = BeautifulSoup(r.text, 'html.parser')

    for link in s.find_all('a'):
        if 'iso' in link.text and 'netinst' in link.text and 'mac' not in link.text and 'edu' not in link.text:
            iso = link.text

    print('https://cdimage.debian.org/cdimage/{}/amd64/iso-cd/SHA512SUMS'.format(release),
          'SHA512SUMS-{}.txt'.format(os.path.splitext(iso)[0]))
    print('https://cdimage.debian.org/cdimage/{}/amd64/iso-cd/SHA512SUMS.sign'.format(release),
          'SHA512SUMS-{}.sign'.format(os.path.splitext(iso)[0]))
    print('https://cdimage.debian.org/cdimage/{}/amd64/iso-cd/{}'.format(release, iso),
          '{}'.format(iso))


def main():
    '''Main function.'''

    # https://en.wikipedia.org/wiki/Debian_version_history#Release_table

    # bullseye 11.x (EoL is 2026?-??-01)
    find_debian_version('weekly-builds')

    # buster 10.x (EoL is 2024-??-01)
    find_debian_version('release/current')

    # stretch 9.x (EoL is 2022-06-01)
    find_debian_version('archive/latest-oldstable')


if __name__ == '__main__':
    main()
