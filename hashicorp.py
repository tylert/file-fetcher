#!/usr/bin/env python

import json

# pip install requests
import requests

from fetcher import find_latest, fetch_file


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


def main():
    '''Main function.'''

    r = requests.get('https://releases.hashicorp.com/index.json')
    r.raise_for_status()

    consul = find_latest(json.loads(r.text)['consul']['versions'].keys())
    consul_aws = find_latest(json.loads(r.text)['consul-aws']['versions'].keys())
    consul_esm = find_latest(json.loads(r.text)['consul-esm']['versions'].keys())
    consul_k8s = find_latest(json.loads(r.text)['consul-k8s']['versions'].keys())
    consul_replicate = find_latest(json.loads(r.text)['consul-replicate']['versions'].keys())
    consul_template = find_latest(json.loads(r.text)['consul-template']['versions'].keys())
    envconsul = find_latest(json.loads(r.text)['envconsul']['versions'].keys())
    packer = find_latest(json.loads(r.text)['packer']['versions'].keys())
    terraform = find_latest(json.loads(r.text)['terraform']['versions'].keys())
    terraform_provider_aws = find_latest(json.loads(r.text)['terraform-provider-aws']['versions'].keys())
    terraform_provider_template = find_latest(json.loads(r.text)['terraform-provider-template']['versions'].keys())
    terraform_provider_terraform = find_latest(json.loads(r.text)['terraform-provider-terraform']['versions'].keys())
    vagrant = find_latest(json.loads(r.text)['vagrant']['versions'].keys())
    vault = find_latest(json.loads(r.text)['vault']['versions'].keys())
    vault_ssh_helper = find_latest(json.loads(r.text)['vault-ssh-helper']['versions'].keys())

    fetch_hashicorp_files('consul', consul)
    fetch_hashicorp_files('consul-aws', consul_aws)
    fetch_hashicorp_files('consul-esm', consul_esm)
    fetch_hashicorp_files('consul-k8s', consul_k8s)
    fetch_hashicorp_files('consul-replicate', consul_replicate)
    fetch_hashicorp_files('consul-template', consul_template)
    fetch_hashicorp_files('envconsul', envconsul)
    fetch_hashicorp_files('packer', packer)
    fetch_hashicorp_files('terraform', terraform)
    fetch_hashicorp_files('terraform-provider-aws', terraform_provider_aws)
    fetch_hashicorp_files('terraform-provider-template', terraform_provider_template)
    fetch_hashicorp_files('terraform-provider-terraform', terraform_provider_terraform)
    fetch_hashicorp_files('vagrant', vagrant)
    fetch_hashicorp_files('vault', vault)
    fetch_hashicorp_files('vault-ssh-helper', vault_ssh_helper)


if __name__ == '__main__':
    main()
