# [WIP]NPM Resource

A Concourse resource to manage a npm package.

## Source Configuration

This resource uses either authentication type for login
**By Token**

* `token`: *Required.* npm registry token.
```
//localhost:4873/:_authToken=[fill the token parameter by this value]
```

**By Password** (not yet implemented)

* `username`: npm registry username.
* `password`: npm registry password.
* `email`: user email for adduser.

When both are listed, `token` is used.

* `package_name`: *Required.* npm package name.
* `registry`: *Optional.* npm registry url. (default is https://registry.npmjs.org/)

## Behaviour

### `check`: Check for new version of specified npm package.

Checks for new versions of specified `package_name` from `registry`.

### `in`: Install npm pachage from registry.

#### Parameters

*There is nothing at the moment.*

### `out`: Publish npm package to registry.

#### Parameters

* `path`: *Required.* Path to the package to be published. (including `package.json`) 
* `version`: *Optional.* [Version](https://docs.npmjs.com/cli/version) of the package to publish.
* `tag`: *Optional.* package tag.

## Pipeline example

```yaml
---
resource_types:
  - name: npm-resource
    type: docker-image
    source:
      repository: idahobean/npm-resource

resources:
  - name: resource-npm-package
    type: git
    source:
      uri: https://github.com/idahobean/sample-node.git

  - name: private-npm-registry
    type: npm-resource
    source:
      token: {{ npm-token }}
      package_name: sample-node
      registry: http://registry.private.npm/

jobs:
- name: job-publish-package
  public: true
  serial: true
  plan:
  - get: resource-npm-package 
    task: build
    file: resource-npm-package/build.yml
  - put: private-npm-registry
    params:
      path: resource-npm-package
      version: patch

```
