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
## Install from online
* Call below command to install the rosa-support command line tool
```$ go install github.com/openshift-online/rosa-support@latest```

* Put the go binary path to PATH env variable
```$ export PATH=$PATH:/$GOPATH/bin:/usr/local/bin```
## How to use the binary of rosa-support to create resources
* Check the help message

	`$ rosa-support -h`

* Create a vpc on the indicated region
    `$ rosa-support create vpc --region us-west-2 --name <your-alias>-vpc`
	* When you have a vpc on the indicated region and you want to re-use it, if cannot find create it
	`$ rosa-support create vpc --region us-west-2 --name <your-alias>-vpc --find-existing`

* Prepare a pair of subnets on the indicated zone of the region.
	* NOTE: If the subnets had been existing, it will reuse the exsiting subnets in the zone
	`$ ./rosa-support create subnets --region us-west-2 --availability-zones us-west-2a --vpc-id <vpc id>`

* Prepare proxy server
	* *--region* is required where created the vpc-id
    * *--vpc-id* is required which should be vpc id used to launch cluster
	* *--availability-zone* is required on which zone to launch the proxy instance
    * *--ca-file* is required which is a abs path used to record the ca-bundle file generated when launch cluster
    * *--keypair-name* is required to generate temporary used to launch the proxy instance
    * *--private-key-path* is required to record the generated private ssh key

	`$ rosa-support create proxy --region <region> --vpc-id <vpc-id>  --availability-zone <az> --ca-file <abs path of ca file> --keypair-name <keypair name> --private-key-path  <path>`

* Prepare addtional security groups, output the sg IDs with comma seperated

	`$ rosa-support create security-groups --region <region> --vpc-id <vpc id> --count <security group number>`

* Tag a resource

	`$ rosa-support tag --resource-id <resource id> --region <region> --tags aaa:ddd,fff:bbb`

* Delete a tag from a resource

	`$ rosa-support delete tag --resource-id <resource id> --region <region> --tag-key <key> --tag-value <value>`

* Clean the vpc and the resources, there is a flag --total-clean supported to do a total clean even the resources is not created by this package

	`$ rosa-support delete vpc --vpc-id <vpc id> --region <region>`

Note that a repository administrator may need to push the tag to the repository due to access restrictions.