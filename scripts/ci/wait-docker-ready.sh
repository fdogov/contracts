#!/usr/bin/env bash
# This script waits for Docker startup.
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
  attempt=$(( 1 ))
  max_attempts=$(( 20 ))
  while ! docker version &>/dev/null; do
    echo_blue "Docker daemon is not ready... (attempt ${attempt})"

    if (( attempt++ == max_attempts )); then
      echo_red "Docker daemon is not ready after ${max_attempts} attempt"
      exit 1
    fi

    sleep 1
  done

  (( attempt-- ))
  echo_green "Docker daemon is ready after ${attempt} seconds"
}

main "${@}"
