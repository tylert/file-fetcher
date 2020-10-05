#!/usr/bin/env python

import json

import requests

from fetcher import find_latest


def find_hashicorp_version(tool, version):
    '''Fetch a given tool version from Hashicorp.'''

    print('https://releases.hashicorp.com/{}/{}/{}_{}_SHA256SUMS'.format(tool, version, tool, version),
          '{}_{}_SHA256SUMS.txt'.format(tool, version))
    print('https://releases.hashicorp.com/{}/{}/{}_{}_SHA256SUMS.sig'.format(tool, version, tool, version),
          '{}_{}_SHA256SUMS.sig'.format(tool, version))

    if tool == 'vagrant':
        print('https://releases.hashicorp.com/{}/{}/{}_{}_x86_64.deb'.format(tool, version, tool, version),
              '{}_{}_x86_64.deb'.format(tool, version))
    else:
        print('https://releases.hashicorp.com/{}/{}/{}_{}_linux_amd64.zip'.format(tool, version, tool, version),
              '{}_{}_linux_amd64.zip'.format(tool, version))


def main():
    '''Main function.'''

    r = requests.get('https://releases.hashicorp.com/index.json')
    r.raise_for_status()

    consul = find_latest(json.loads(r.text)['consul']['versions'].keys())
    # consul_aws = find_latest(json.loads(r.text)['consul-aws']['versions'].keys())
    # consul_esm = find_latest(json.loads(r.text)['consul-esm']['versions'].keys())
    # consul_k8s = find_latest(json.loads(r.text)['consul-k8s']['versions'].keys())
    # consul_replicate = find_latest(json.loads(r.text)['consul-replicate']['versions'].keys())
    # consul_template = find_latest(json.loads(r.text)['consul-template']['versions'].keys())
    # envconsul = find_latest(json.loads(r.text)['envconsul']['versions'].keys())
    packer = find_latest(json.loads(r.text)['packer']['versions'].keys())
    terraform = find_latest(json.loads(r.text)['terraform']['versions'].keys())
    terraform_provider_aws = find_latest(json.loads(r.text)['terraform-provider-aws']['versions'].keys())
    # terraform_provider_template = find_latest(json.loads(r.text)['terraform-provider-template']['versions'].keys())
    # terraform_provider_terraform = find_latest(json.loads(r.text)['terraform-provider-terraform']['versions'].keys())
    vagrant = find_latest(json.loads(r.text)['vagrant']['versions'].keys())
    vault = find_latest(json.loads(r.text)['vault']['versions'].keys())
    # vault_ssh_helper = find_latest(json.loads(r.text)['vault-ssh-helper']['versions'].keys())

    find_hashicorp_version('consul', consul)
    # find_hashicorp_version('consul-aws', consul_aws)
    # find_hashicorp_version('consul-esm', consul_esm)
    # find_hashicorp_version('consul-k8s', consul_k8s)
    # find_hashicorp_version('consul-replicate', consul_replicate)
    # find_hashicorp_version('consul-template', consul_template)
    # find_hashicorp_version('envconsul', envconsul)
    find_hashicorp_version('packer', packer)
    find_hashicorp_version('terraform', terraform)
    find_hashicorp_version('terraform-provider-aws', terraform_provider_aws)
    # find_hashicorp_version('terraform-provider-template', terraform_provider_template)
    # find_hashicorp_version('terraform-provider-terraform', terraform_provider_terraform)
    find_hashicorp_version('vagrant', vagrant)
    find_hashicorp_version('vault', vault)
    # find_hashicorp_version('vault-ssh-helper', vault_ssh_helper)


if __name__ == '__main__':
    main()
