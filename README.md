pyprojectversionparser
======================
[![build](https://github.com/SIMBAChain/pyprojectversionparser/actions/workflows/build.yaml/badge.svg?branch=main)](https://github.com/SIMBAChain/pyprojectversionparser/actions/workflows/build.yaml)  [![release](https://github.com/SIMBAChain/pyprojectversionparser/actions/workflows/release.yaml/badge.svg?branch=main)](https://github.com/SIMBAChain/pyprojectversionparser/actions/workflows/release.yaml)

Mini tool for quickly grabbing the version from a pyproject.toml file

[Get the latest release here](https://github.com/SIMBAChain/pyprojectversionparser/releases/tag/v0.0.2)

### Running
```bash
$ pyprojectversionparser ./pyproject.toml envfile
PACKAGE_VERSION=0.1.0
PACKAGE_NAME=my-package-name
$ cat envfile
PACKAGE_VERSION=0.1.0
PACKAGE_NAME=nd-nft-server
```