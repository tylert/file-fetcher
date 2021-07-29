#!/usr/bin/env python

import json
import logging

import requests
import semver


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
            latest = semver.max_ver(latest, version)
        else:
            logging.warning(f'SKIPPING {version}')

    return latest


def main():
    '''Main function.'''

    r = requests.get('https://releases.hashicorp.com/index.json')
    r.raise_for_status()

    print(f'consul:')
    print(f'  release: {get_latest_semver_hashicorp(json.loads(r.text)["consul"]["versions"].keys())}')
    print(f'nomad:')
    print(f'  release: {get_latest_semver_hashicorp(json.loads(r.text)["nomad"]["versions"].keys())}')
    print(f'packer:')
    print(f'  release: {get_latest_semver_hashicorp(json.loads(r.text)["packer"]["versions"].keys())}')
    print(f'terraform:')
    print(f'  release: {get_latest_semver_hashicorp(json.loads(r.text)["terraform"]["versions"].keys())}')
    print(f'terraform-provider-aws:')
    print(f'  release: {get_latest_semver_hashicorp(json.loads(r.text)["terraform-provider-aws"]["versions"].keys())}')
    print(f'terraform-provider-http:')
    print(f'  release: {get_latest_semver_hashicorp(json.loads(r.text)["terraform-provider-http"]["versions"].keys())}')
    print(f'terraform-provider-vsphere:')
    print(f'  release: {get_latest_semver_hashicorp(json.loads(r.text)["terraform-provider-vsphere"]["versions"].keys())}')
    print(f'vagrant:')
    print(f'  release: {get_latest_semver_hashicorp(json.loads(r.text)["vagrant"]["versions"].keys())}')
    print(f'vault:')
    print(f'  release: {get_latest_semver_hashicorp(json.loads(r.text)["vault"]["versions"].keys())}')
    print(f'vault-k8s:')
    print(f'  release: {get_latest_semver_hashicorp(json.loads(r.text)["vault-k8s"]["versions"].keys())}')
    print(f'vault-ssh-helper:')
    print(f'  release: {get_latest_semver_hashicorp(json.loads(r.text)["vault-ssh-helper"]["versions"].keys())}')
    print(f'waypoint')
    print(f'  release: {get_latest_semver_hashicorp(json.loads(r.text)["waypoint"]["versions"].keys())}')
    print(f'waypoint-entrypoint')
    print(f'  release: {get_latest_semver_hashicorp(json.loads(r.text)["waypoint-entrypoint"]["versions"].keys())}')


if __name__ == '__main__':
    main()
