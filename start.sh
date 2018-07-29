#!/bin/bash
set -e -u -o pipefail

declare -r SCRIPT_PATH=$(readlink -f "$0")
declare -r SCRIPT_DIR=$(cd $(dirname "$SCRIPT_PATH") && pwd)


active_conda_env() {
  conda env list | grep ' \* ' | cut -f 1 -d ' '
}


main() {
  cd "$(git rev-parse --show-toplevel)"

  # conda activate notes
  # calling source activate doesn't work ... so make sure user has done it
  [[ $(active_conda_env) != "notes" ]] && {
    echo "ERROR: 'notes' not active; did you forget 'conda activate notes' ?"
    return 1
  }

  export GOPATH=$PWD/go
  export GOBIN=$GOPATH/bin
  export PATH=$PATH:$GOBIN

  set -x
  exec jupyter notebook "$@"
}

main "$@"
