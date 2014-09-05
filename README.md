feedify
=======

what should be done
-------------------
![structure](https://cloud.githubusercontent.com/assets/1843523/4164978/36a80e50-34fd-11e4-8620-fa8b30d37d0a.png)


#### Development
Create symbolic link from you directory to `GOPATH`
```
cd to-your-project
mkdir -p $GOPATH/src/github.com/feedlabs
ln -s $(pwd) $GOPATH/src/github.com/feedlabs/feedify
```

#### Run
`go run test/feedify.go`

#### Dependencies
* `go get github.com/astaxie/beego`
* `go get github.com/fzzy/radix/redis`
* `go get github.com/barakmich/glog`
