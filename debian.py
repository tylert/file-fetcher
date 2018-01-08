#!/usr/bin/env python3

import os

# pip install requests
import requests

# pip install beautifulsoup4
from bs4 import BeautifulSoup

from fetcher import fetch_file


def fetch_debian_archive_files(release, iso):
    '''Fetch a Debian archive release.'''

    fetch_file('https://cdimage.debian.org/cdimage/archive/{}/amd64/iso-cd/SHA512SUMS'.format(release),
               'SHA512SUMS-{}.txt'.format(os.path.splitext(iso)[0]))
    fetch_file('https://cdimage.debian.org/cdimage/archive/{}/amd64/iso-cd/SHA512SUMS.sign'.format(release),
               'SHA512SUMS-{}.sign'.format(os.path.splitext(iso)[0]))
    fetch_file('https://cdimage.debian.org/cdimage/archive/{}/amd64/iso-cd/{}'.format(release, iso),
               '{}'.format(iso))


def fetch_debian_release_files(iso):
    '''Fetch a Debian release.'''

    fetch_file('https://cdimage.debian.org/cdimage/release/current/amd64/iso-cd/SHA512SUMS',
               'SHA512SUMS-{}.txt'.format(os.path.splitext(iso)[0]))
    fetch_file('https://cdimage.debian.org/cdimage/release/current/amd64/iso-cd/SHA512SUMS.sign',
               'SHA512SUMS-{}.sign'.format(os.path.splitext(iso)[0]))
    fetch_file('https://cdimage.debian.org/cdimage/release/current/amd64/iso-cd/{}'.format(iso),
               '{}'.format(iso))


def fetch_debian_testing_files(iso):
    '''Fetch a Debian testing release.'''

    fetch_file('https://cdimage.debian.org/cdimage/weekly-builds/amd64/iso-cd/SHA512SUMS',
               'SHA512SUMS-{}.txt'.format(os.path.splitext(iso)[0]))
    fetch_file('https://cdimage.debian.org/cdimage/weekly-builds/amd64/iso-cd/SHA512SUMS.sign',
               'SHA512SUMS-{}.sign'.format(os.path.splitext(iso)[0]))
    fetch_file('https://cdimage.debian.org/cdimage/weekly-builds/amd64/iso-cd/{}'.format(iso),
               '{}'.format(iso))


if __name__ == '__main__':

    # wheezy (EoL is 2018-05-??)
    r = requests.get('https://cdimage.debian.org/cdimage/archive/7.11.0/amd64/iso-cd')
    r.raise_for_status()
    s = BeautifulSoup(r.text, 'html.parser')

    for link in s.find_all('a'):
        if 'iso' in link.text and 'netinst' in link.text and '-mac-' not in link.text:
            iso = link.text
    fetch_debian_archive_files('7.11.0', iso)

    # jessie (EoL is 2020-04-?? or 2020-05-??)
    r = requests.get('https://cdimage.debian.org/cdimage/archive/latest-oldstable/amd64/iso-cd')
    r.raise_for_status()
    s = BeautifulSoup(r.text, 'html.parser')

    for link in s.find_all('a'):
        if 'iso' in link.text and 'netinst' in link.text and '-mac-' not in link.text:
            iso = link.text
    fetch_debian_archive_files('latest-oldstable', iso)

    # stretch (EoL is 2022-06-??)
    r = requests.get('https://cdimage.debian.org/cdimage/release/current/amd64/iso-cd')
    r.raise_for_status()
    s = BeautifulSoup(r.text, 'html.parser')

    for link in s.find_all('a'):
        if 'iso' in link.text and 'netinst' in link.text and '-mac-' not in link.text:
            iso = link.text
    fetch_debian_release_files(iso)

    # buster (EoL is 20??-??-??)
    r = requests.get('https://cdimage.debian.org/cdimage/weekly-builds/amd64/iso-cd')
    r.raise_for_status()
    s = BeautifulSoup(r.text, 'html.parser')

    for link in s.find_all('a'):
        if 'iso' in link.text and 'netinst' in link.text and '-mac-' not in link.text:
            iso = link.text
    fetch_debian_testing_files(iso)
