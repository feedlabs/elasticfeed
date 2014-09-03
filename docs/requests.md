API requests
------------


#### Examples
```
curl feed.dev:10111/v1/feed/
{}

curl -XPOST -H "content-type:application/json" -d '{"Data":"new feed page"}' feed.dev:10111/v1/feed
{
  "id": "1409731977215584330"
}

curl -XPOST -H "content-type:application/json" -d '{Data:"new feed entry"}' feed.dev:10111/v1/feed/1409731977215584330/entry
{
  "id": "1409731990598877748"
}

curl -XPUT -H "content-type:application/json" -d '{"Data":"updated new feed page"}' feed.dev:10111/v1/feed/1409731977215584330
"update success!"

curl -XPUT -H "content-type:application/json" -d '{"Data":"updated new feed entry"}' feed.dev:10111/v1/feed/1409731977215584330/entry/1409731990598877748
"update success!"

curl -XDELETE feed.dev:10111/v1/feed/1409731977215584330/entry/1409731990598877748
"delete success!"

curl -XDELETE feed.dev:10111/v1/feed/1409731977215584330
"delete success!"

curl feed.dev:10111/v1/feed/
{}
```
