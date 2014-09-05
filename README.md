api
===

#### Development
Create symbolic link from you directory to `GOPATH`
```
cd to-your-project
mkdir -p $GOPATH/src/github.com/feedlabs
ln -s $(pwd) $GOPATH/src/github.com/feedlabs/api
```

#### Run
`go run api.go`

#### Dependencies
* `go get github.com/feedlabs/feedify`
