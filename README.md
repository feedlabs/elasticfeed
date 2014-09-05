api
===

what should be done
-------------------
![structure](https://cloud.githubusercontent.com/assets/1843523/4171633/2f18586a-353f-11e4-90e9-b68d8a7d27a1.png)

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
