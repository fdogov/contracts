#!/usr/bin/env bash
# This script runs "buf breaking" if not disabled.
set \
  -o errexit \
  -o nounset \
  -o pipefail

echo_red() {
  echo_color 31 "${@}"
}

echo_green() {
  echo_color 32 "${@}"
}

echo_blue() {
  echo_color 34 "${@}"
}

echo_color() {
  color="${1}"
  shift

  text="${@}"

  echo -e "\e[${color}m--> ${text}\e[0m"
}

main() {
  # Only run for MRs.
  if [[ "${CI_PIPELINE_SOURCE:=}" != "merge_request_event" ]] ; then
    echo_blue "Skip buf breaking check for non-MR pipeline."
    exit 0
  fi

  # If there is a "skip-breaking" label, skip the breaking check.
  if [[ "${CI_MERGE_REQUEST_LABELS:=}" =~ skip-breaking ]] ; then
    echo_blue "Skip buf breaking check for this MR."
    exit 0
  fi

  echo_blue "Run buf breaking ..."

  set +e
  ${CI_PROJECT_DIR:=.}/scripts/buf/run.sh buf breaking --against ".git#ref=${CI_MERGE_REQUEST_DIFF_BASE_SHA}"
  result=$?
  set -e

  if [[ ${result} -eq 0 ]]; then
    echo_green "No breaking changes found!"
    exit 0
  else
    echo_red "Found breaking changes! Please fix them before merging."
    echo_red "If you want to bypass this breaking check, add 'skip-breaking' label to the Merge Request."
    exit 1
  fi
}

main "${@}"


