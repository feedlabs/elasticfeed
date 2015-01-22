var elasticfeed = (function() {

  var defaults = {
    transport: 'ws'
  }

  var actions = {
    entry: ['add', 'delete', 'update'],
    feed: ['load-init', 'load-more', 'reload', 'reset'],
    metric: ['click', 'skip', 'back', 'scroll']
  }

  function ef() {
    this.channel = null
  }

  // elastic feed communication

  ef.prototype.getFeedRoom = function() {
    // can exist many feed-rooms
  }

  ef.prototype.getMetricRoom = function() {
    // should exist only single metric room
  }

  ef.prototype.getChannel = function() {
    // should exist only single channel connection
  }

  // Transportation

  ef.prototype.detectTransportMode = function() {

  }

  ef.prototype.newWebSocket = function() {

  }

  ef.prototype.newLongPooling = function() {

  }

  return ef;

})();


var $ = {
  getJSON: function(url, callback) {
    xhr = new XMLHttpRequest;
    xhr.onreadystatechange = function() {
      if (xhr.readyState == 4 && xhr.status == 200) {
        if (xhr.responseText != "") {
          data = JSON.parse(xhr.responseText);
          callback.call(this, data, xhr.readyState)
        } else {
          callback.call(this, null, xhr.readyState)
        }
      } else {
        callback.call(this, null, xhr.readyState)
      }
    }
    xhr.open("GET", url)
    xhr.send();
  },

  each: function(obj, callback) {
    for (i = 0; i < obj.length; i++) {
      value = callback.call(obj[i], i, obj[i]);

      if (value === false) {
        break;
      }
    }
  },

  queryString: function(obj) {
    return Object.keys(obj).map(function(key) {
      return encodeURIComponent(key) + '=' + encodeURIComponent(obj[key]);
    }).join('&');
  },

  post: function(url, data, callback) {
    xhr1 = new XMLHttpRequest;
    xhr1.onreadystatechange = function() {
      if (xhr1.readyState == 4 && xhr1.status == 200) {
        callback.call(this, xhr1.responseText)
      }
    }
    dataString = this.queryString(data)
    xhr1.open("POST", url + "?" + dataString, true);
    xhr1.send(dataString);
  }
}

var ef = {
  ws: function(username) {
    socket = new WebSocket('ws://localhost:10100/ws/join?uname=' + username);

    socket.onmessage = function(event) {
      var data = JSON.parse(event.data);
      console.log(data);
      switch (data.Type) {
        case 0: // JOIN
          if (data.User == username) {
            console.log('joined the room');
          } else {
            console.log(data.User + " joined the chat room");
          }
          break;
        case 1: // LEAVE
          console.log(data.User + " left the chat room");
          break;
        case 2: // MESSAGE
          console.log(data.User + ", " + data.Content);
          break;
      }
    };

    return socket;
  },

  lp: function(username) {
    var lastReceived = 0;
    var isWait = false;

    var fetch = function() {
      if (isWait) {
        return;
      }
      isWait = true;
      $.getJSON("http://localhost:10100/lp/fetch?lastReceived=" + lastReceived, function(data, code) {

        if (code == 4) {
          isWait = false
        }

        if (data == null) {
          return;
        }

        $.each(data, function(i, event) {
          switch (event.Type) {
            case 0: // JOIN
              if (event.User == username) {
                console.log('joined the room');
              } else {
                console.log(event.User + " joined the chat room");
              }
              break;
            case 1: // LEAVE
              console.log(event.User + " left the chat room");
              break;
            case 2: // MESSAGE
              console.log(event.User + ", " + event.Content);
              break;
          }

          lastReceived = event.Timestamp;
        });
        isWait = false;
      });
    }

    setInterval(fetch, 3000);
    fetch()

    return fetch;
  }
}
