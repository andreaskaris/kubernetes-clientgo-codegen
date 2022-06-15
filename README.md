# What is this?

This is a minimal example for using the automatic code generation that is used for
https://github.com/kubernetes/sample-controller by means of the code generator at
https://github.com/kubernetes/code-generator/. deepcopy-gen, client-gen, lister-gen
and informer-gen will generate the client code that's needed to interact with a
custom resource via the kubernetes API.

None of this code (other than `main.go`) is mine and all credit goes to the original
authors, this repo shall just serve as a minimal example for what's needed to get
code-generation up and running.

This minimal example contains the following:
* `/hack/tools.go` will pull in the code-generator
* CRDs and a CR unter `/artifacts`
* `hack/update-codegen.sh` to call the code generators, and a boilerplate template that's needed for code generation
* `/pkg/apis/foo/*` contains the type definition with the annotations for the code generator

# How to use this?

For code generation to work, this code *must* be under `$GOPATH/src`. I tried a lot of variations to not have to do this (because my projects normally live outside of $GOPATH, call me a weirdo ...); none work. Hence, check this out into your gopath, under `$GOPATH/src/example.com/m`.

Then, run the following to generate the client-go code:
~~~
mv main.go main
go mod tidy
go vendor
hack/update-codegen.sh
mv main main.go
~~~

Now, add a custom resource definition and a single `Foo` resource:
~~~
export KUBECONFIG=kubeconfig
kubectl apply -f artifacts/crd.yaml
kubectl apply -f artifacts/example-foo.yaml
~~~

Now, compile and run `main.go`:
~~~
go mod tidy
go mod vendor
go run main.go
~~~

