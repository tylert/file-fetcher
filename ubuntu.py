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


def fetch_ubuntu_testing_files(iso):
    '''Fetch an Ubuntu testing release.'''

    fetch_file('http://cdimage.ubuntu.com/ubuntu-server/daily/current/SHA256SUMS',
               'SHA256SUMS-{}.txt'.format(os.path.splitext(iso)[0]))
    fetch_file('http://cdimage.ubuntu.com/ubuntu-server/daily/current/SHA256SUMS.gpg',
               'SHA256SUMS-{}.gpg'.format(os.path.splitext(iso)[0]))
    fetch_file('http://cdimage.ubuntu.com/ubuntu-server/daily/current/{}'.format(iso),
               '{}'.format(iso))


def main():
    '''Main function.'''

    # trusty 14.04.x (EoL is 2019-04-??)
    r = requests.get('http://releases.ubuntu.com/trusty')
    r.raise_for_status()
    s = BeautifulSoup(r.text, 'html.parser')

    for link in s.find_all('a'):
        if 'server' in link.text and 'iso' in link.text and 'amd64' in link.text and 'torrent' not in link.text and 'zsync' not in link.text:
            iso = link.text
    fetch_ubuntu_release_files('trusty', iso)

    for link in s.find_all('a'):
        if 'desktop' in link.text and 'iso' in link.text and 'amd64' in link.text and 'torrent' not in link.text and 'zsync' not in link.text:
            iso = link.text
    fetch_ubuntu_release_files('trusty', iso)

    # xenial 16.04.x (EoL is 2021-04-??)
    r = requests.get('http://releases.ubuntu.com/xenial')
    r.raise_for_status()
    s = BeautifulSoup(r.text, 'html.parser')

    for link in s.find_all('a'):
        if 'server' in link.text and 'iso' in link.text and 'amd64' in link.text and 'torrent' not in link.text and 'zsync' not in link.text:
            iso = link.text
    fetch_ubuntu_release_files('xenial', iso)

    for link in s.find_all('a'):
        if 'desktop' in link.text and 'iso' in link.text and 'amd64' in link.text and 'torrent' not in link.text and 'zsync' not in link.text:
            iso = link.text
    fetch_ubuntu_release_files('xenial', iso)

    # artful 17.10 (EoL is 2018-07-??)
    r = requests.get('http://releases.ubuntu.com/artful')
    r.raise_for_status()
    s = BeautifulSoup(r.text, 'html.parser')

    for link in s.find_all('a'):
        if 'server' in link.text and 'iso' in link.text and 'amd64' in link.text and 'torrent' not in link.text and 'zsync' not in link.text:
            iso = link.text
    fetch_ubuntu_release_files('artful', iso)

    for link in s.find_all('a'):
        if 'desktop' in link.text and 'iso' in link.text and 'amd64' in link.text and 'torrent' not in link.text and 'zsync' not in link.text:
            iso = link.text
    fetch_ubuntu_release_files('artful', iso)

    # bionic 18.04.x (EoL is 2023-04-??)
    r = requests.get('http://cdimage.ubuntu.com/ubuntu-server/daily/current')
    r.raise_for_status()
    s = BeautifulSoup(r.text, 'html.parser')

    for link in s.find_all('a'):
        if 'iso' in link.text and 'amd64' in link.text and 'zsync' not in link.text:
            iso = link.text
    fetch_ubuntu_testing_files(iso)


if __name__ == '__main__':
    main()
