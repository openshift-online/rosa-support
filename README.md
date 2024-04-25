# ROSA Support
`rosa-support` is a tool designed to be a plugin to the [ROSA CLI](https://github.com/openshift/rosa) by adding useful commands for engineers to support and develop the CLI

## Build 
To build the binary run 
``` 
$ make build
```
To install the binary run onto your `$GOPATH`
``` 
$ make install
```

## Releases a new CLI Version
Releasing a new version requires submitting an MR for review/merge with an update to Version constant in [version.go](pkg/version/version.go).
Additionally. update the [CHANGES.md](CHANGES.md) file to include the new version and describe all changes included

Below is an example CHANGES.md update:
```
## 0.0.1

- Added an awesome change
```

Submit an MR for review/merge with the CHANGES.md and version.go update.

Finally, create and submit a new tag with the new version following the below example:

```shell
git checkout main
git pull
git tag -a -m 'Release 0.0.1' v0.0.1
git push origin v0.0.1
```

Note that a repository administrator may need to push the tag to the repository due to access restrictions.