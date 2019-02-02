#!/usr/bin/env python3

from fetcher import fetch_file


def main():
    '''Main function.'''

    # 7.x (EoL is 2024-06-30)
    fetch_file('http://centos.mirror.iweb.ca/7/isos/x86_64/sha256sum.txt',
               'SHA256SUMS-centos-7.6-minimal.txt')
    fetch_file('http://centos.mirror.iweb.ca/7/isos/x86_64/sha256sum.txt.asc',
               'SHA256SUMS-centos-7.6-minimal.txt.asc')
    fetch_file('http://centos.mirror.iweb.ca/7/isos/x86_64/CentOS-7-x86_64-Minimal-1810.iso',
               'centos-7.6-minimal.iso')

    # 6.x (EoL is 2020-11-30)
    fetch_file('http://centos.mirror.iweb.ca/6.10/isos/x86_64/sha256sum.txt',
               'SHA256SUMS-centos-6.10-minimal.txt')
    fetch_file('http://centos.mirror.iweb.ca/6.10/isos/x86_64/sha256sum.txt.asc',
               'SHA256SUMS-centos-6.10-minimal.txt.asc')
    fetch_file('http://centos.mirror.iweb.ca/6.10/isos/x86_64/CentOS-6.10-x86_64-minimal.iso',
               'centos-6.10-minimal.iso')


if __name__ == '__main__':
    main()
