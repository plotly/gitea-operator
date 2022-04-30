# Gitea Operator

An Operator that installs Gitea. Installation is performed by creating a custom resource of kind `Gitea`. You can uninstall Gitea by removing this resource.
The Operator will also watch all Gitea resources and reinstall them if they are deleted.

This operator use old version of the operator-sdk ( 0.1.1 ) , also we need to use golang version 1.10 because it use old library for the `controller-runtime`.
See this file with the information : [pkg/apis/integreatly/v1alpha1/register.go](./pkg/apis/integreatly/v1alpha1/register.go)

## Tags information 

Version v0.0.6 == containe the build version 
Version v0.0.6-1 == is the same but with the fix for the controller-runtime, I did it to be able to do `go get github.com/plotly/gitea-operator@v0.0.6-1`

## What the operator provide 

* Secrets : 
    * admin-core-secret
    * gitea-init
    * gitea-admin-secret
    * gitea-config

* Deployment : 
    * postgres
    * gitea

* pvc : 
    * gitea-postgres-pvc
    * gitea-repos

* Service : 
    * gitea-postgresql   ClusterIP   10.43.57.158   <none>        5432/TCP
    * gitea-http         ClusterIP   None           <none>        3000/TCP
    * gitea-ssh          ClusterIP   None           <none>        22/TCP

* ingress :
    * gitea-ingress


## Operator object configuration

```
apiVersion: integreatly.org/v1alpha1
kind: Gitea
metadata:
  name: gitea-cluster
spec:
  hostname: "example.dash.surf"
```

## Setup Dev Env

To setup your dev env you can use those cmd : 

```sh
$ make dockerBuildEnv/build
$ make dockerBuildEnv/run
```

The `dockerBuildEnv/run` will mount local directory in the POD : 
* \$PWD:/go/src/github.com/integr8ly/gitea-operator"  : The code of the operator
* \${HOME}/.kube:/root/.kube"  : your Kubernetes configuration to be able to run the operator on the cluster
* "/var/run/docker.sock:/var/run/docker.sock" : To be able to build the docker image

When your env is running you will have to run those command : ( TODO: Need to check if I can add it in the Dockerfile )

```
root@a92c9dee272c:/go/src/github.com/integr8ly/gitea-operator# export USER=root

root@a92c9dee272c:/go/src/github.com/integr8ly/gitea-operator# make setup/dep
Installing golang dependencies
Installing dep
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  5230  100  5230    0     0  11761      0 --:--:-- --:--:-- --:--:-- 11886
ARCH = amd64
OS = linux
Will install into /go/bin
Fetching https://github.com/golang/dep/releases/latest..
Release Tag = v0.5.4
Fetching https://github.com/golang/dep/releases/tag/v0.5.4..
Fetching https://github.com/golang/dep/releases/download/v0.5.4/dep-linux-amd64..
Setting executable permissions.
Moving executable to /go/bin/dep
setup complete
```

```
root@a92c9dee272c:/go/src/github.com/integr8ly/gitea-operator# make setup/travis
Installing Operator SDK
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   675  100   675    0     0   1759      0 --:--:-- --:--:-- --:--:--  1785
100 37.4M  100 37.4M    0     0  10.5M      0  0:00:03  0:00:03 --:--:-- 12.6M
```

### Run the operator in Dev mode

You need to add the CRD to your cluster

```
$ kubectl apply -f deploy/crds/crd.yaml
```

Now you can create your Custom Ressouce : 

```
$ kubectl apply -f example-gitea.yaml
```

Now you can do you code change and run the operator from the container : 

```sh
root@a92c9dee272c:/go/src/github.com/integr8ly/gitea-operator# make code/run
2022/04/11 14:18:22 Go Version: go1.10.8
2022/04/11 14:18:22 Go OS/Aprobarch: linux/amd64
2022/04/11 14:18:22 operator-sdk Version: 0.0.7
2022/04/11 14:18:22 Registering Components.
2022/04/11 14:18:22 Starting the Cmd.
2022/04/11 14:18:22 Reconciling Gitea gitea/example-gitea
2022/04/11 14:18:22 Gitea image is up to date: quay.io/integreatly/gitea:1.10.3
2022/04/11 14:18:52 Reconciling Gitea gitea/example-gitea
2022/04/11 14:18:53 Gitea image is up to date: quay.io/integreatly/gitea:1.10.3
```

To try something else simply delete the cr and recreate it :

```
$ kubectl delete gitea example-gitea
$ kubectl apply -f example-gitea.yaml
```

### Build docker image 

in the dev docker container  : 

```
# make image/build
```

Push it to quay

```
# make image/push
```

## Troubleshooting

If you have an error with this librari `controller-runtime` change the import, check the file : [pkg/apis/integreatly/v1alpha1/register.go](./pkg/apis/integreatly/v1alpha1/register.go)
