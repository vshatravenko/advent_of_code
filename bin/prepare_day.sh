#!/usr/bin/env bash

set -e

# Usage: ./prepare_day.sh *lang* *day* *year*

if (( $# != 3 )); then
    echo "Invalid number of arguments!"
    printf "Usage: ./prepare_day.sh *lang* *day* *year*\ne.g. ./prepare_day.sh go 04 2023"
    exit 1
fi

src="tpl/$1"
dest="$3/$2"

if [ ! -d "${src}" ]; then
    echo "${src} is missing, available languages: $(ls tpl/)"
    exit 1
fi

mkdir -p "${dest}"
cp -R "${src}" "${dest}"

echo "${dest}/${1} is ready!"
