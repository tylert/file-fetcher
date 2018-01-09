#!/usr/bin/env python3

import os

# pip install requests
import requests

# pip install beautifulsoup4
from bs4 import BeautifulSoup

from fetcher import fetch_file


def fetch_debian_release(release):
    '''Fetch a complete Debian release.'''

    r = requests.get('https://cdimage.debian.org/cdimage/{}/amd64/iso-cd'.format(release))
    r.raise_for_status()
    s = BeautifulSoup(r.text, 'html.parser')

    for link in s.find_all('a'):
        if 'iso' in link.text and 'netinst' in link.text and '-mac-' not in link.text:
            iso = link.text

    fetch_file('https://cdimage.debian.org/cdimage/{}/amd64/iso-cd/SHA512SUMS'.format(release),
               'SHA512SUMS-{}.txt'.format(os.path.splitext(iso)[0]))
    fetch_file('https://cdimage.debian.org/cdimage/{}/amd64/iso-cd/SHA512SUMS.sign'.format(release),
               'SHA512SUMS-{}.sign'.format(os.path.splitext(iso)[0]))
    fetch_file('https://cdimage.debian.org/cdimage/{}/amd64/iso-cd/{}'.format(release, iso),
               '{}'.format(iso))


def main():
    '''Main function.'''

    # wheezy 7.x (EoL is 2018-05-??)
    fetch_debian_release('archive/7.11.0')

    # jessie 8.x (EoL is 2020-04-?? or 2020-05-??)
    fetch_debian_release('archive/latest-oldstable')

    # stretch 9.x (EoL is 2022-06-??)
    fetch_debian_release('release/current')

    # buster 10.x (EoL is 2024-??-??)
    fetch_debian_release('weekly-builds')


if __name__ == '__main__':
    main()
