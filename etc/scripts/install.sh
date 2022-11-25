#!/usr/bin/env bash

VERSION="v0.0.0"

function install_guardrails() {
  if [[ "$OSTYPE" == "linux"* ]]; then
      case $(uname -m) in
          aarch64) ARCH=arm64;;
          armv7l)  ARCH=arm;;
          *)       ARCH=$(uname -m);;
      esac
      set -x
      curl -fsSL https://github.com/guardrailsio/guardrails-cli/releases/download/${VERSION}/guardrails_${VERSION}_linux_${ARCH}.tar.gz | tar -xzv guardrails
      sudo mv guardrails /usr/local/bin/guardrails
  elif [[ "$OSTYPE" == "darwin"* ]]; then
      ARCH=$(uname -m)
      set -x
      curl -fsSL https://github.com/guardrailsio/guardrails-cli/releases/download/${VERSION}/guardrails_${VERSION}_darwin_${ARCH}.tar.gz | tar -xzv guardrails
      sudo mv guardrails /usr/local/bin/guardrails
  else
      set +x
      echo "OS is not supported: $OSTYPE"
      exit 1
  fi

  set +x
}

install_guardrails