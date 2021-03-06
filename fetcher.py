from hashlib import sha512, sha256, sha1, md5
import logging
import os

import semver
import requests


def get_latest_semver_hashicorp(versions):
    '''Scan all versions in list to find the latest one for HashiCorp products.'''

    latest = '0.0.0'

    for version in versions:
        # Drop anything that's closed-source (Enterprise), alpha or beta or
        # rc1, etc.
        if '+ent' not in version and '-alpha' not in version \
                and '-beta' not in version and '-connect' not in version \
                and '-oci' not in version and '-rc' not in version:
            logging.warning('{} {}'.format(latest, version))
            latest = semver.max_ver(latest, version)
        else:
            logging.warning('SKIPPING {}'.format(version))

    return latest


def get_latest_semver(versions):
    '''Scan all versions in list to find the latest one.'''

    latest = '0.0.0'

    for version in versions:
        logging.warning('{} {}'.format(latest, version))
        latest = semver.max_ver(latest, version)

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

        logging.warning('{},{},{},{}'.format(output, existing_size, new_size,
                                             new_type))

        # Check if the existing file is already the expected size
        if new_type != 'text/plain':
            if existing_size == new_size:
                return

        with open(output, 'wb') as outfile:
            for chunk in r.iter_content(1024):
                outfile.write(chunk)


def hash_file(filename, hash_method='sha512', chunk_size=2**16):
    '''Calculate the hash of a file using the specified method.'''

    # https://stackoverflow.com/questions/17731660/hashlib-optimal-size-of-chunks-to-be-used-in-md5-update

    # Any requested hash method that is invalid or that we forgot to import in
    # this module, will cause an exception
    file_hash = globals()[hash_method]()

    if os.path.isfile(filename) and os.access(filename, os.R_OK):
        with open(filename, 'rb') as filehandle:
            while True:
                chunk = filehandle.read(chunk_size)
                if not chunk:
                    break
                file_hash.update(chunk)
        return file_hash.hexdigest()
    else:
        return ''
