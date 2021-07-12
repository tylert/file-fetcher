#!/usr/bin/env bash

repo="${1}"

release=$(curl --head --silent "https://github.com/${repo}/releases/latest" | grep location: | awk -F"/" '{ printf "%s", $NF }' | tr -d '\r')

echo "${release}"
