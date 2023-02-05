# Nested Modules

Shows how you can use one repository for nested modules written on Go. Each submodule has unique tag,

```sh
git clone git@github.com:tullo/nested-modules.git
cd nested-modules
git tag -l
# app/v0.0.2
# lib/v0.0.1
```

Start example app

```sh
go get github.com/tullo/nested-modules/lib
go run app/main.go 
08:57:12
```

## General Best Practice

The general best practice is to put your `go.mod` file at the root of the repo, and have only one module per repo. Then all of the packages within the repo will be in the same module and share the same dependencies.

> It is possible to define multiple modules within a single repository, but doing so adds a lot of maintenance overhead: if you have multiple modules, they are treated as completely separate; you have to explicitly manage the dependencies between them, and you have to explicitly tag separate releases for them.

[For most projects, that extra overhead comes with very little benefit](https://stackoverflow.com/questions/67448535/how-do-i-create-a-nested-go-module-within-a-repository)
