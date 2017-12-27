#!/usr/bin/env python3

import json

# pip install requests
import requests

from fetcher import (find_latest, fetch_hashicorp_files)


if __name__ == '__main__':
    r = requests.get('https://releases.hashicorp.com/index.json')
    r.raise_for_status()

    packer = find_latest(json.loads(r.text)['packer']['versions'].keys())
    terraform = find_latest(json.loads(r.text)['terraform']['versions'].keys())
    terraform_provider_aws = find_latest(json.loads(r.text)['terraform-provider-aws']['versions'].keys())
    terraform_provider_template = find_latest(json.loads(r.text)['terraform-provider-template']['versions'].keys())
    terraform_provider_terraform = find_latest(json.loads(r.text)['terraform-provider-terraform']['versions'].keys())
    vagrant = find_latest(json.loads(r.text)['vagrant']['versions'].keys())

    fetch_hashicorp_files('packer', packer)
    fetch_hashicorp_files('terraform', terraform)
    fetch_hashicorp_files('terraform-provider-aws', terraform_provider_aws)
    fetch_hashicorp_files('terraform-provider-template', terraform_provider_template)
    fetch_hashicorp_files('terraform-provider-terraform', terraform_provider_terraform)
    fetch_hashicorp_files('vagrant', vagrant)
