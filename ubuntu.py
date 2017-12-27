#!/usr/bin/env python3

# pip install requests
import requests

# pip install beautifulsoup4
from bs4 import BeautifulSoup

from fetcher import (fetch_ubuntu_release_files,
                     fetch_ubuntu_testing_files)


if __name__ == '__main__':
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

    r = requests.get('http://cdimage.ubuntu.com/ubuntu-server/daily/current')
    r.raise_for_status()
    s = BeautifulSoup(r.text, 'html.parser')

    for link in s.find_all('a'):
        if 'iso' in link.text and 'amd64' in link.text and 'zsync' not in link.text:
            iso = link.text
    fetch_ubuntu_testing_files(iso)
