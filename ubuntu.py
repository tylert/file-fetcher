#!/usr/bin/env python3

import os

# pip install requests
import requests

# pip install beautifulsoup4
from bs4 import BeautifulSoup

from fetcher import fetch_file


def fetch_ubuntu_release_files(release, iso):
    '''Fetch an Ubuntu release.'''

    fetch_file('http://releases.ubuntu.com/{}/SHA256SUMS'.format(release),
               'SHA256SUMS-{}.txt'.format(os.path.splitext(iso)[0]))
    fetch_file('http://releases.ubuntu.com/{}/SHA256SUMS.gpg'.format(release),
               'SHA256SUMS-{}.gpg'.format(os.path.splitext(iso)[0]))
    fetch_file('http://releases.ubuntu.com/{}/{}'.format(release, iso),
               '{}'.format(iso))


def main():
    '''Main function.'''

    # bionic 18.04.x (EoL is 2023-04-01)
    r = requests.get('http://cdimage.ubuntu.com/releases/bionic/release')
    r.raise_for_status()
    s = BeautifulSoup(r.text, 'html.parser')

    for link in s.find_all('a'):
        if 'server' in link.text and 'iso' in link.text and 'amd64' in link.text and 'torrent' not in link.text and 'zsync' not in link.text:
            iso = link.text
    fetch_file('http://cdimage.ubuntu.com/releases/bionic/release/SHA256SUMS',
               'SHA256SUMS-{}.txt'.format(os.path.splitext(iso)[0]))
    fetch_file('http://cdimage.ubuntu.com/releases/bionic/release/SHA256SUMS.gpg',
               'SHA256SUMS-{}.gpg'.format(os.path.splitext(iso)[0]))
    fetch_file('http://cdimage.ubuntu.com/releases/bionic/release/{}'.format(iso),
               '{}'.format(iso))

    r = requests.get('http://releases.ubuntu.com/bionic')
    r.raise_for_status()
    s = BeautifulSoup(r.text, 'html.parser')

    for link in s.find_all('a'):
        if 'server' in link.text and 'iso' in link.text and 'amd64' in link.text and 'torrent' not in link.text and 'zsync' not in link.text:
            iso = link.text
    fetch_ubuntu_release_files('bionic', iso)

    # xenial 16.04.x (EoL is 2021-04-01)
    r = requests.get('http://releases.ubuntu.com/xenial')
    r.raise_for_status()
    s = BeautifulSoup(r.text, 'html.parser')

    for link in s.find_all('a'):
        if 'server' in link.text and 'iso' in link.text and 'amd64' in link.text and 'torrent' not in link.text and 'zsync' not in link.text:
            iso = link.text
    fetch_ubuntu_release_files('xenial', iso)


if __name__ == '__main__':
    main()
