# XXX TODO Convert print to logging
from __future__ import print_function

import hashlib
import os
import json

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


def hash_file(directory, filename, blocksize=2**20, hash_method='sha512'):
    '''Calculate the checksum of a file using the specified method.'''

    if hash_method == 'sha512':
        file_hash = hashlib.sha512()
    elif hash_method == 'sha256':
        file_hash = hashlib.sha256()
    elif hash_method == 'sha2':
        file_hash = hashlib.sha2()
    else:
        file_hash = hashlib.sha1()

    if os.path.isfile(os.path.join(directory, filename)) and \
            os.access(os.path.join(directory, filename), os.R_OK):

        with open(os.path.join(directory, filename), 'rb') as filehandle:
            while True:
                block = filehandle.read(blocksize)
                if not block:
                    break
                file_hash.update(block)

    return file_hash.hexdigest()
