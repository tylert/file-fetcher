# XXX TODO Convert print to logging
from __future__ import print_function

import os

# pip install semver
import semver

# pip install requests
import requests


def find_latest(versions):
    '''Scan all versions in list to find the latest one.'''

    latest = '0.0.0'

    for version in versions:
        print('{} {}'.format(latest, version))
        latest = semver.max_ver(latest, version)

    return latest


def fetch_file(url, output):
    '''Fetch a given file from url and save it as output.'''

    r = requests.get(url, stream=True)
    r.raise_for_status()

    print('Fetching {}'.format(output))

    with open(output, 'wb') as outfile:
        for block in r.iter_content(1024):
            outfile.write(block)


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
