#!/usr/bin/env python3

# pip install requests
import requests

# pip install beautifulsoup4
from bs4 import BeautifulSoup

from fetcher import (fetch_debian_archive_files,
                     fetch_debian_release_files, fetch_debian_testing_files)


if __name__ == '__main__':
    r = requests.get('https://cdimage.debian.org/cdimage/archive/latest-oldstable/amd64/iso-cd')
    r.raise_for_status()
    s = BeautifulSoup(r.text, 'html.parser')

    for link in s.find_all('a'):
        if 'iso' in link.text and 'netinst' in link.text and '-mac-' not in link.text:
            iso = link.text
    fetch_debian_archive_files('latest-oldstable', iso)

    r = requests.get('https://cdimage.debian.org/cdimage/release/current/amd64/iso-cd')
    r.raise_for_status()
    s = BeautifulSoup(r.text, 'html.parser')

    for link in s.find_all('a'):
        if 'iso' in link.text and 'netinst' in link.text and '-mac-' not in link.text:
            iso = link.text
    fetch_debian_release_files(iso)

    r = requests.get('https://cdimage.debian.org/cdimage/weekly-builds/amd64/iso-cd')
    r.raise_for_status()
    s = BeautifulSoup(r.text, 'html.parser')

    for link in s.find_all('a'):
        if 'iso' in link.text and 'netinst' in link.text and '-mac-' not in link.text:
            iso = link.text
    fetch_debian_testing_files(iso)
