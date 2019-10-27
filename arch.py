#!/usr/bin/env python

import requests
from bs4 import BeautifulSoup

from fetcher import fetch_file


def main():
    '''Main function.'''

    r = requests.get('https://muug.ca/mirror/archlinux/iso/latest')
    r.raise_for_status()
    s = BeautifulSoup(r.text, 'html.parser')

    for link in s.find_all('a'):
        if 'iso' in link.text and 'torrent' not in link.text and 'sig' not in link.text:
            iso = link.text

    fetch_file('https://muug.ca/mirror/archlinux/iso/latest/{}.sig'.format(iso),
               '{}.sig'.format(iso))
    fetch_file('https://muug.ca/mirror/archlinux/iso/latest/{}'.format(iso),
               '{}'.format(iso))

    for link in s.find_all('a'):
        if 'bootstrap' in link.text and 'sig' not in link.text:
            iso = link.text

    fetch_file('https://muug.ca/mirror/archlinux/iso/latest/{}.sig'.format(iso),
               '{}.sig'.format(iso))
    fetch_file('https://muug.ca/mirror/archlinux/iso/latest/{}'.format(iso),
               '{}'.format(iso))


if __name__ == '__main__':
    main()
