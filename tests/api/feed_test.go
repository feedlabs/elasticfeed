curl feed.dev:10111/v1/feed/

curl -XPOST -H "content-type:application/json" -d '{"Data":"halo entry"}' feed.dev:10111/v1/feed


curl -H "content-type:application/json" -d '{Data:"halo entry"}' feed.dev:10111/v1/feed/123/entry


curl -XDELETE feed.dev:10111/v1/feed/1409693489600564956/entry/1409693548994215890

curl -XPUT feed.dev:10111/v1/feed/1409693489600564956/entry/1409693548994215890
