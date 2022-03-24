projectversionparser
======================
[![build](https://github.com/SIMBAChain/pyprojectversionparser/actions/workflows/build.yaml/badge.svg?branch=main)](https://github.com/SIMBAChain/pyprojectversionparser/actions/workflows/build.yaml)  [![release](https://github.com/SIMBAChain/pyprojectversionparser/actions/workflows/release.yaml/badge.svg?branch=main)](https://github.com/SIMBAChain/pyprojectversionparser/actions/workflows/release.yaml) [![Coverage Status](https://coveralls.io/repos/github/SIMBAChain/pyprojectversionparser/badge.svg?branch=main)](https://coveralls.io/github/SIMBAChain/pyprojectversionparser?branch=main)

Simple tool for quickly grabbing the version from a pyproject.toml or package.json file

[Get the latest release here](https://github.com/SIMBAChain/pyprojectversionparser/releases/tag/v0.0.2)

### Running

Running without any arguments will look for project files in the current directory and will parse the first one it finds.

Running with an argument will use that as the filename to try to parse.

```bash
$ projectversionparser ./pyproject.toml
PACKAGE_VERSION=0.1.0
PACKAGE_NAME=my-package-name
```

To set the output as env variables, use either:

```bash
$ source <(projectversionparser)
```

or 
```bash
$ projectversionparser > envfile
$ source ./envfile
```