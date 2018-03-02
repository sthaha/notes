#!/bin/bash
set -e -u -o pipefail

declare -r SCRIPT_PATH=$(readlink -f "$0")
declare -r SCRIPT_DIR=$(cd $(dirname "$SCRIPT_PATH") && pwd)


main() {
  cd $(git rev-parse --show-toplevel)
  source activate notes

  export GOPATH=$PWD/go
  export GOBIN=$GOPATH/bin
  export PATH=$PATH:$GOBIN

  exec jupyter notebook "$@"
}

main "$@"
