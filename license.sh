#!/usr/bin/env bash

header='// Derivative of "The Go Programming Language" © 2016 examples by
// Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
'
dir=${1:-.}
tmp=$(mktemp)

for file in $(find "${dir}" -type f -not -path '*/\.*' -not -iname readme.md); do
  if ! grep -q 'Derivative of "The Go Programming Language"' "${file}"; then
    {
    echo "${header}"
    cat "${file}"
    } > "${tmp}"
    mv "${tmp}" "${file}"
  fi
done
rm -f "${tmp}"