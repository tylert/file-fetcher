import logging
from hashlib import sha512, sha256, sha1, md5
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


def hash_file(filename, directory=None, hash_method='sha512', chunksize=2**16):
    '''Calculate the hash of a file using the specified method.'''

    # Any requested hash method that is invalid or that we forgot to import in
    # this module, will cause an exception
    file_hash = globals()[hash_method]()

    if directory is None:
        full_path = filename
    else:
        full_path = os.path.join(directory, filename)

    if os.path.isfile(full_path) and os.access(full_path, os.R_OK):
        with open(full_path, 'rb') as filehandle:
            while True:
                chunk = filehandle.read(chunksize)
                if not chunk:
                    break
                file_hash.update(chunk)
        return file_hash.hexdigest()
    else:
        return ''
