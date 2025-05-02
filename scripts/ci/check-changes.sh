#!/usr/bin/env bash
# This script checks for Git changes. If it founds any, this means that the
# task which is checked is not reproducible.
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
  echo_blue "Check git worktree for changes ..."

  git_status="$(git status --porcelain=v1)"

  if [[ -z "${git_status}" ]]; then
    echo_green "No changes found."
    return 0
  fi

  echo_red "Found changes!"
  echo "${git_status}"

  git update-index -q --refresh
  git diff-index -p --color=always HEAD --

  return 1
}

main "${@}"
