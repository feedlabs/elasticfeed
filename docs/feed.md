Feed API
--------

Message
-------
```
message, _ = stream.NewStreamMessage()
data := `{"channel":"iO5wshd5fFE5YXxJ/hfyKQ==:17","event":"CM_Action_Abstract:SEND:31","data":{"action":{"actor":{"_type":33,"_id":{"id":"1"},"id":1,"displayName":"user1","visible":true,"_class":"Feed_Model_User"},"verb":13,"type":31,"_class":"Feed_Action_Feed"},"model":{"_type":33,"_id":{"id":"1"},"id":1,"displayName":"user1","visible":true,"_class":"Feed_Model_User"},"data":{"action":"add","clientData":{"type":"photo","imageList":[{"url":"http://lorempixel.com/500/500/?46577","title":"Image"},{"url":"http://lorempixel.com/500/500/?81474","title":"Image"}],"likeCount":0,"id":19}}}}`
message.Publish(data)
```

POST
----
```
curl -XPOST -H "Content-type:application/json" -d '{"Data":"{\"channel\":\"iO5wshd5fFE5YXxJ/hfyKQ==:17\",\"event\":\"CM_Action_Abstract:SEND:31\",\"data\":{\"action\":{\"actor\":{\"_type\":33,\"_id\":{\"id\":\"1\"},\"id\":1,\"displayName\":\"user1\",\"visible\":true,\"_class\":\"Feed_Model_User\"},\"verb\":13,\"type\":31,"_class\":\"Feed_Action_Feed\"},\"model\":{\"_type\":33,\"_id\":{\"id\":\"1\"},\"id\":1,\"displayName\":\"user1\",\"visible\":true,\"_class\":\"Feed_Model_User\"},\"data\":{\"action\":\"add\",\"clientData\":{\"type\":\"photo\",\"imageList\":[{\"url\":\"http://lorempixel.com/500/500/?4651277\",\"title\":\"Image\"},{\"url\":\"http://lorempixel.com/500/500/?8551474\",\"title\":\"Image\"}],\"likeCount\":0,\"id\":19}}}}"}' http://localhost:10100/v1/feed
```

Subscribe
---------
```
  message, _ = stream.NewStreamMessage()

	channels := []string{"socket-redis-down"}
	message.Subscribe(channels, func(timeout bool, message string, channel string) {
		if !timeout {
			fmt.Println("publish:", message, " channel:", channel)
		} else {
			fmt.Println("error: sub timedout")
		}
	})
```
