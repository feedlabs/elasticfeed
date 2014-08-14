feedify
=======

what should be done
-------------------
![structure](https://cloud.githubusercontent.com/assets/1843523/3927887/bb8c4260-2401-11e4-944f-3f89eeefb72f.png)


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
`go get github.com/astaxie/beego`
