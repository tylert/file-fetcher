#!/usr/bin/env python3

from fetcher import fetch_file


def main():
    '''Main function.'''

    # 7.x (EoL is 2024-06-30)
    fetch_file('http://mirror.centos.org/centos/7.4.1708/os/x86_64/images/boot.iso',
               'centos-7.4-boot.iso')

    fetch_file('http://vault.centos.org/centos/7.3.1611/os/x86_64/images/boot.iso',
               'centos-7.3-boot.iso')

    fetch_file('http://vault.centos.org/centos/7.2.1511/os/x86_64/images/boot.iso',
               'centos-7.2-boot.iso')

    fetch_file('http://vault.centos.org/centos/7.1.1503/os/x86_64/images/boot.iso',
               'centos-7.1-boot.iso')

    fetch_file('http://vault.centos.org/centos/7.0.1406/os/x86_64/images/boot.iso',
               'centos-7.0-boot.iso')

    # 6.x (EoL is 2020-11-30)
    fetch_file('http://mirror.centos.org/centos/6.9/os/x86_64/images/boot.iso',
               'centos-6.9-boot.iso')

    fetch_file('http://vault.centos.org/centos/6.8/os/x86_64/images/boot.iso',
               'centos-6.8-boot.iso')

    fetch_file('http://vault.centos.org/centos/6.7/os/x86_64/images/boot.iso',
               'centos-6.7-boot.iso')

    fetch_file('http://vault.centos.org/6.6/os/x86_64/images/boot.iso',
               'centos-6.6-boot.iso')

    fetch_file('http://vault.centos.org/6.5/os/x86_64/images/boot.iso',
               'centos-6.5-boot.iso')


if __name__ == '__main__':
    main()
