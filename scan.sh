#!/bin/sh

export SKIP_TLS_VERIFICATION=true

# Local
#export GUARDRAILS_API_HOST=https://api.dev.guardrails.io
#alias scan="/Users/haohoang/workspace/05-gr/guardrails-cli/bin/guardrails-amd64-darwin scan -t d9b82da26c4f2b9a6377dade679c57f151e8706eebe7d48fe2b3985809ece895"

# Staging
export GUARDRAILS_API_HOST=https://api.staging.k8s.guardrails.io
alias scan="/Users/haohoang/workspace/05-gr/guardrails-cli/bin/guardrails-amd64-darwin scan -t 883046ce8d25e88ee9c50034ba5d3ed037fa58ede5d78068b5d86c82665c8b87"

# Azure
#export GUARDRAILS_API_HOST=https://api.staging.k8s.guardrails.io
#alias scan="/Users/haohoang/workspace/05-gr/guardrails-cli/bin/guardrails-amd64-darwin scan -t 4891fd3784dc815964af1bed8d96acb95342f49f74af788afee691de38fef4ed"

execute() {
    local path=${1}
    echo "scanning: ${path} ..."
    scan -p ${path}
}

main() {
    current_context=$(kubectl config current-context)
    echo "Context: $current_context"
    execute $@
}

# ex: ./scan.sh /Users/haohoang/workspace/05-gr/tmp/python-test-01
main $@