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


def fetch_hashicorp_files(tool, version):
    '''Fetch a given tool version from Hashicorp.'''

    fetch_file('https://releases.hashicorp.com/{}/{}/{}_{}_SHA256SUMS'.format(tool, version, tool, version),
               '{}_{}_SHA256SUMS.txt'.format(tool, version))
    fetch_file('https://releases.hashicorp.com/{}/{}/{}_{}_SHA256SUMS.sig'.format(tool, version, tool, version),
               '{}_{}_SHA256SUMS.sig'.format(tool, version))

    if tool == 'vagrant':
        fetch_file('https://releases.hashicorp.com/{}/{}/{}_{}_x86_64.deb'.format(tool, version, tool, version),
                   '{}_{}_x86_64.deb'.format(tool, version))
    else:
        fetch_file('https://releases.hashicorp.com/{}/{}/{}_{}_linux_amd64.zip'.format(tool, version, tool, version),
                   '{}_{}_linux_amd64.zip'.format(tool, version))


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
