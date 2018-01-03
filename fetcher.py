# XXX TODO Convert print to logging
from __future__ import print_function

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

    with requests.get(url, stream=True) as r, \
            open(output, 'wb') as outfile:
        r.raise_for_status()

        print('Fetching {} size {}'.format(output, int(r.headers['content-length'])))

        for block in r.iter_content(1024):
            outfile.write(block)
