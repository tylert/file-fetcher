#!/usr/bin/env python3

import re

# pip install requests
import requests

# pip install beautifulsoup4
from bs4 import BeautifulSoup

from fetcher import fetch_file


def main():
    '''Main function.'''

    r = requests.get('https://github.com/keepassxreboot/keepassxc/releases/latest')
    r.raise_for_status()
    s = BeautifulSoup(r.text, 'html.parser')

    for link in s.find_all('a'):
        if 'tar.xz' in link.text and 'sig' not in link.text and 'DIGEST' not in link.text:
            matcher = re.compile('\d+\.\d+\.\d+')
            version = matcher.findall(link.text.split('\n')[2])[0]

    fetch_file('https://github.com/keepassxreboot/keepassxc/releases/download/{}/keepassxc-{}-src.tar.xz.DIGEST'.format(version, version),
               'keepassxc-{}-src.tar.xz.DIGEST'.format(version))
    fetch_file('https://github.com/keepassxreboot/keepassxc/releases/download/{}/keepassxc-{}-src.tar.xz.sig'.format(version, version),
               'keepassxc-{}-src.tar.xz.sig'.format(version))
    fetch_file('https://github.com/keepassxreboot/keepassxc/releases/download/{}/keepassxc-{}-src.tar.xz'.format(version, version),
               'keepassxc-{}-src.tar.xz'.format(version))


if __name__ == '__main__':
    main()
