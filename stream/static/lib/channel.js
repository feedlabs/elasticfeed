var Channel = (function() {

  var defaultOptions = {
    transport: 'ws',
    connectOnInit: true
  }

  var defaultCredential = {
    username: null,
    token: null,
    method: 'basic'
  }

  function Channel(options, credential) {

    /** @type {String} */
    this.id = null;

    /** @type {Object} */
    this.options = _extend(defaultOptions, options);

    /** @type {Object} */
    this.credential = _extend(defaultCredential, credential);
  }

  Channel.prototype.registerFeedHandler = function(feed, callback) {

  }

  Channel.prototype.onData = function(data) {
    // get handlers
    event = new StreamEvent(data)
  }

  Channel.prototype.Authenticate = function(credential) {
  }

  Channel.prototype.getConnection = function() {
  }

  Channel.prototype.getWebSocketConnection = function(username) {
    this._socket = new WebSocket('ws://localhost:10100/ws/join?uname=' + username);

    this._socket.onmessage = function(event) {
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

    self = this
    return {
      send: function(data) {
        self._socket.send(data)
      }
    };
  }

  Channel.prototype.getLongPoolingConnection = function(username) {
    var lastReceived = 0;
    var isWait = false;

    this.getJSON('http://localhost:10100/lp/join?uname=' + username, function() {})

    self = this;
    var fetch = function() {
      if (isWait) {
        return;
      }
      isWait = true;
      self.getJSON("http://localhost:10100/lp/fetch?lastReceived=" + lastReceived, function(data, code) {

        if (code == 4) {
          isWait = false
        }

        if (data == null) {
          return;
        }

        self.each(data, function(i, event) {
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

    return {
      send: function(data) {
        self.post("/lp/post", data, function(status) { });
      }
    };
  }

  // HTTP

  Channel.prototype.getJSON = function(url, callback) {
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
  }

  Channel.prototype.each = function(obj, callback) {
    for (i = 0; i < obj.length; i++) {
      value = callback.call(obj[i], i, obj[i]);

      if (value === false) {
        break;
      }
    }
  }

  Channel.prototype.queryString = function(obj) {
    return Object.keys(obj).map(function(key) {
      return encodeURIComponent(key) + '=' + encodeURIComponent(obj[key]);
    }).join('&');
  }

  Channel.prototype.post = function(url, data, callback) {
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

  // Helpers

  var _extend = function(a, b) {
    var c = {}, prop;
    for (prop in a) {
      if (a.hasOwnProperty(prop)) {
        c[prop] = a[prop];
      }
    }
    for (prop in b) {
      if (b.hasOwnProperty(prop)) {
        c[prop] = b[prop];
      }
    }
    return c;
  }

  return Channel;

})();
