<p align="center">
    <p align="center">
        <img src="https://www.guardrails.io/assets/images/logo-color.png" alt="GuardRails" title="GuardRails" height="100px" align="center"/>
    </p>
    <h1 align="center"><b>GuardRails CLI</b></h1>
</p>

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/guardrailsio/guardrails-cli)
![GitHub](https://img.shields.io/github/license/guardrailsio/guardrails-cli)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/guardrailsio/guardrails-cli?sort=semver)
[![GuardRails CLI CI](https://github.com/guardrailsio/guardrails-cli/actions/workflows/ci.yaml/badge.svg)](https://github.com/guardrailsio/guardrails-cli/actions/workflows/ci.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/guardrailsio/guardrails-cli)](https://goreportcard.com/report/github.com/guardrailsio/guardrails-cli) 
[![codecov](https://codecov.io/gh/guardrailsio/guardrails-cli/branch/main/graph/badge.svg)](https://codecov.io/gh/guardrailsio/guardrails-cli)
[![GuardRails badge](https://api.guardrails.io/v2/badges/145474?token=5cdcd3c9f602bdf5dd4ec8a7d19e2e6599e571a73e8c9751f5b6d04deaf68aa6)](https://dashboard.guardrails.io/gh/guardrailsio/repos/145474)

The GuardRails CLI allows you to interact with [GuardRails](https://www.guardrails.io) via the command line.

## Table of Contents

- [Pre-Requisites](#pre-requisites)
- [Installation](#installation)
- [Usage](#usage)
- [Documentation](#documentation)
- [License](#license)

## Pre-Requisites

To use the GuardRails CLI, you require an active GuardRails account and a CLI token.

More information on how to get started can be found [here](https://www.guardrails.io/docs/en/getting-started).

Your GuardRails account CLI token can be obtained under `Settings`->`CLI Authentication` on the GuardRails dashboard.

## Installation

### Installation scripts (Linux / OSX)

Just paste this command, and you're good to go. We're assuming you're using `bash`, but you can change it accordingly based on the shell you're using. You might be asked for a password for `sudo` in the installation process.

```
curl -fsSL https://raw.githubusercontent.com/guardrailsio/guardrails-cli/main/etc/scripts/install.sh | bash
```

### Brew (Linux / OSX)

Alternatively, you can also install `guardrails` via `brew`:

```
brew tap guardrailsio/guardrails
brew install guardrails
```

### Windows

You require [scoop](https://scoop.sh) before installing `guardrails`. The rest will be similar to the installation scripts for Linux / OSX. Execute the below command in your powershell:

```
iex ((new-object net.webclient).DownloadString('https://raw.githubusercontent.com/guardrailsio/guardrails-cli/main/etc/scripts/install.ps1'))
```

## Usage

Here are the main GuardRails CLI commands:
- `scan`    : Scans a repository for vulnerabilities and outputs results
- `version` : Displays the build version
 
For more information on all the options and available arguments, please check the help menu with: `guardrails --help`

### How to read the results

The CLI will output the total number of detected vulnerabilities.
Vulnerabilities are grouped by category, i.e., `Hard-Coded Secrets`.

For each item within a category, the following information is shown:
- A severity index (see table below).
- The type of vulnerability containing a hyperlink to fixing advice in our documentation.
- The file path and line number.

Example: `(M) Hard-coded Secret - awesome-product/config.js:2`

Here we're looking at a vulnerability of type `Hard-coded secret` with a `Medium` severity in the file `awesome-product/config.js` at line `2`.

For Vulnerable Libraries specifically, the type of vulnerability will be replaced by the dependency name and version.

Example: `(C) pkg:gem/mypackage@2.5.2 - awesome-product/Gemfile.lock:14`

Here we're looking at the vulnerable `mypackage` dependency in version `2.5.2` with a `Critical` severity declared in the file `awesome-product/Gemfile.lock` at line `14`.

#### Severity index table

| Index | Severity      |
|-------|---------------|
| (N/A) | Not available |
| (I)   | Informational |
| (L)   | Low           |
| (M)   | Medium        |
| (H)   | High          |
| (C)   | Critical      |

## Documentation

https://www.guardrails.io/docs/en/cli/introduction

## License

The GuardRails CLI is released under the Apache 2.0 license. See [LICENSE.txt](https://github.com/guardrailsio/guardrails-cli/blob/main/LICENSE.txt)
