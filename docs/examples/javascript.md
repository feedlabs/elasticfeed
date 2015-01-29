### JS client

#### Usage
```
  window.onload = function() {
    elasticfeed.init({
      channel: {
        url: 'ws://localhost:80',
        transport: 'ws'
      }
    });

    feed = elasticfeed.initFeed('000001', {
      outputContainerId: 'my-elastic-feed-1',
      stylerFunction: function(data) {
        return '<div style="height:50px; border:1px dotted; border-color: blue;">' + data + '</div>';
      }
    });

    feed.channel.on('join', function(chid, ts) {
      feed1.addEntry(chid + " joined the chat room");
    });

    feed.channel.on('leave', function(chid, ts) {
      feed1.addEntry(chid + " left the chat room");
    });

    window['socket'] = feed1.socket;
  }
```

#### Test data broadcast
```
  // single feed
  socket.send({Type:1, Timestamp:1111111, Content: {Type:3, Timestamp:22222, Id: "000001", Content: "data-examples"}})

  // all feeds in the view on the chanel
  socket.send({Type:1, Timestamp:1111111, Content: {Type:3, Timestamp:22222, Id: "*", Content: "data-examples"}})
```
