# NPM Resource

A Concourse resource to manage a npm package.

## Source Configuration

* `package_name`: *Required.* npm package name.
* `registry`: *Optional.* npm registry url. (default is https://registry.npmjs.org/)

## Behaviour

### `check`: Check for new version of specified npm package.

Checks for new version of specified `package_name` from `registry`.

### `in`: Install npm package from registry.

#### Parameters

*There is nothing at the moment.*  
The assumed use case in `in` is cache npm package into private registry.

### `out`: Publish npm package to registry.

#### Parameters

* `username`: *Required.* npm registry login username.
* `password`: *Required.* npm registry login password.
* `email`: *Required.* npm registry login email.
* `path`: *Required.* Path to the package to be published. (including `package.json`) 
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
      package_name: sample-node
      registry: http://registry.private.npm

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
      username: foo
      password: bar
      email: baz@fox.qoo
      path: resource-npm-package
      tag: latest

```
