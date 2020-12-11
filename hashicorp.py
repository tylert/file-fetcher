#!/usr/bin/env python

import json

import requests

from fetcher import get_latest_semver_hashicorp


def main():
    '''Main function.'''

    r = requests.get('https://releases.hashicorp.com/index.json')
    r.raise_for_status()

    consul = get_latest_semver_hashicorp(json.loads(r.text)['consul']['versions'].keys())
    packer = get_latest_semver_hashicorp(json.loads(r.text)['packer']['versions'].keys())
    terraform = get_latest_semver_hashicorp(json.loads(r.text)['terraform']['versions'].keys())
    terraform_provider_aws = get_latest_semver_hashicorp(json.loads(r.text)['terraform-provider-aws']['versions'].keys())
    terraform_provider_http = get_latest_semver_hashicorp(json.loads(r.text)['terraform-provider-http']['versions'].keys())
    vagrant = get_latest_semver_hashicorp(json.loads(r.text)['vagrant']['versions'].keys())
    vault = get_latest_semver_hashicorp(json.loads(r.text)['vault']['versions'].keys())

    print('consul:')
    print('  release: {}'.format(consul))
    print('packer:')
    print('  release: {}'.format(packer))
    print('terraform:')
    print('  release: {}'.format(terraform))
    print('terraform-provider-aws:')
    print('  release: {}'.format(terraform_provider_aws))
    print('terraform-provider-http:')
    print('  release: {}'.format(terraform_provider_http))
    print('vagrant:')
    print('  release: {}'.format(vagrant))
    print('vault:')
    print('  release: {}'.format(vault))


if __name__ == '__main__':
    main()
