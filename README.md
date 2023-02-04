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
