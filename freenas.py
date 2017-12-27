#!/usr/bin/env python3

import os

# pip install requests
import requests

# pip install beautifulsoup4
from bs4 import BeautifulSoup

from fetcher import fetch_file


def fetch_freenas_release_files(iso):
    '''Fetch a FreeNAS release.'''

    # FreeNAS does not provide a convenient download link for the PDF version
    # of the users' guide.  They only provide an online HTML version at
    # https://doc.freenas.org/.

    # To download a giant HTML blob, visit
    # https://download.freenas.org/stable/.

    # To download ReStructuredText versions needed to generate a PDF, visit
    # https://github.com/freenas/freenas-docs/.  NOTE:  The Makefile doesn't
    # work with GNU Make on Linux without changes and the requisite build
    # environment is a non-trivial install.

    # An unofficial PDF and epub version is available at
    # http://freenas.2trux.com.

    fetch_file('http://freenas.2trux.com/FreeNAS.pdf',
               '{}.pdf'.format(os.path.splitext(iso)[0]))
    fetch_file('https://download.freenas.org/stable/x64/{}.sha256'.format(iso),
               '{}.sha256'.format(iso))
    fetch_file('https://download.freenas.org/stable/x64/{}'.format(iso),
               '{}'.format(iso))


if __name__ == '__main__':
    r = requests.get('https://download.freenas.org/stable/x64')
    r.raise_for_status()
    s = BeautifulSoup(r.text, 'html.parser')

    for link in s.find_all('a'):
        if 'iso' in link.text and 'sha256' not in link.text:
            iso = link.text
    fetch_freenas_release_files(iso)
