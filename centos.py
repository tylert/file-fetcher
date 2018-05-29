#!/usr/bin/env python3

from fetcher import fetch_file


def main():
    '''Main function.'''

    # 7.x (EoL is 2024-06-30)
    fetch_file('http://centos.mirror.iweb.ca/7.5.1804/isos/x86_64/CentOS-7-x86_64-Minimal-1804.iso',
               'centos-7.5-boot.iso')

    fetch_file('http://vault.centos.org/7.4.1708/isos/x86_64/CentOS-7-x86_64-Minimal-1708.iso',
               'centos-7.4-boot.iso')

    fetch_file('http://vault.centos.org/7.3.1611/isos/x86_64/CentOS-7-x86_64-Minimal-1611.iso',
               'centos-7.3-boot.iso')

    fetch_file('http://vault.centos.org/7.2.1511/isos/x86_64/CentOS-7-x86_64-Minimal-1511.iso',
               'centos-7.2-boot.iso')

    fetch_file('http://vault.centos.org/7.1.1503/isos/x86_64/CentOS-7-x86_64-Minimal-1503.iso',
               'centos-7.1-boot.iso')

    fetch_file('http://vault.centos.org/7.0.1406/isos/x86_64/CentOS-7.0-1406-x86_64-Minimal.iso',
               'centos-7.0-boot.iso')

    # 6.x (EoL is 2020-11-30)
    fetch_file('http://centos.mirror.iweb.ca/6.9/isos/x86_64/boot.iso',
               'centos-6.9-boot.iso')

    fetch_file('http://vault.centos.org/6.8/isos/x86_64/CentOS-6.8-x86_64-minimal.iso',
               'centos-6.8-boot.iso')

    fetch_file('http://vault.centos.org/6.7/isos/x86_64/CentOS-6.7-x86_64-minimal.iso',
               'centos-6.7-boot.iso')

    fetch_file('http://vault.centos.org/6.6/isos/x86_64/CentOS-6.6-x86_64-minimal.iso',
               'centos-6.6-boot.iso')

    fetch_file('http://vault.centos.org/6.5/isos/x86_64/CentOS-6.5-x86_64-minimal.iso',
               'centos-6.5-boot.iso')


if __name__ == '__main__':
    main()
