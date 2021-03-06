#!/bin/bash

BIN="$(cd "$(dirname "$0")" ; pwd)"
SRC="$(dirname "${BIN}")"
PROJECT="$(dirname "${SRC}")"

source "${BIN}/verbose.sh"

BIND="${PROJECT}"
if [ ".$1" = '.--bind' ]
then
  BIND="$2"
  shift 2
fi

mkdir -p ~/.cache/nix

docker run -ti --rm -v "jeroenvm-nix-store:/nix/store" -v "${HOME}/.cache/nix:/root/.cache/nix" -v "${BIND}:${BIND}" -w "$(pwd)" 'jeroenvm/archetype-nix-go' "$@"
