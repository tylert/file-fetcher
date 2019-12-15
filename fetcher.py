import logging
import hashlib
import os

import semver
import requests


def find_latest(versions):
    '''Scan all versions in list to find the latest one.'''

    latest = '0.0.0'

    for version in versions:
        if '+ent' not in version and '-alpha' not in version \
                and '-beta' not in version and '-rc' not in version \
                and '-oci' not in version:
            logging.warning('{} {}'.format(latest, version))
            latest = semver.max_ver(latest, version)
        else:
            logging.warning('SKIPPING {}'.format(version))

    return latest


def fetch_file(url, output):
    '''Fetch a given file from url and save it as output.'''

    # Find the size of the file if it exists already
    existing_size = 0
    try:
        existing_size = os.stat(output).st_size
    except FileNotFoundError:
        logging.warning('Not found {}'.format(output))

    with requests.get(url, stream=True) as r:
        r.raise_for_status()

        new_size = int(r.headers['content-length'])

        # Try to determine the content type, if possible
        new_type = 'unset'
        try:
            new_type = r.headers['content-type']
        except KeyError:
            new_type = 'unknown'

        logging.warning('{},{},{},{}'.format(output, existing_size, new_size, new_type))

        # Check if the existing file is already the expected size
        if new_type != 'text/plain':
            if existing_size == new_size:
                return

        with open(output, 'wb') as outfile:
            for chunk in r.iter_content(1024):
                outfile.write(chunk)


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
