#!/usr/bin/env python

import json
import logging

import requests
from semver import max_ver


def get_latest_semver_hashicorp(versions):
    '''Scan all versions in list to find the latest one for HashiCorp products.'''

    latest = '0.0.0'

    for version in versions:
        # Drop anything that's closed-source (Enterprise), alpha, beta, rc,
        # etc.
        if '+ent' not in version and '-alpha' not in version \
                and '-beta' not in version and '-connect' not in version \
                and '-oci' not in version and '-rc' not in version:
            logging.warning(f'{latest} {version}')
            latest = max_ver(latest, version)
        else:
            logging.warning(f'SKIPPING {version}')

    return latest


def main():
    '''Main function.'''

    r = requests.get('https://releases.hashicorp.com/terraform/index.json')
    r.raise_for_status()

    print(f'terraform:')
    print(f'  release: {get_latest_semver_hashicorp(json.loads(r.text)["versions"].keys())}')


if __name__ == '__main__':
    main()
