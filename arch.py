#!/usr/bin/env python3

# pip install requests
import requests

# pip install beautifulsoup4
from bs4 import BeautifulSoup

from fetcher import fetch_file


def fetch_arch_release_files(iso):
    '''Fetch an Arch release.'''

    fetch_file('https://mirror.cedille.club/archlinux/iso/latest/{}.sig'.format(iso),
               '{}.sig'.format(iso))
    fetch_file('https://mirror.cedille.club/archlinux/iso/latest/{}'.format(iso),
               '{}'.format(iso))


if __name__ == '__main__':
    r = requests.get('https://mirror.cedille.club/archlinux/iso/latest')
    r.raise_for_status()
    s = BeautifulSoup(r.text, 'html.parser')

    for link in s.find_all('a'):
        if 'iso' in link.text and 'torrent' not in link.text and 'sig' not in link.text:
            iso = link.text
    fetch_arch_release_files(iso)

    for link in s.find_all('a'):
        if 'bootstrap' in link.text and 'sig' not in link.text:
            iso = link.text
    fetch_arch_release_files(iso)
