var Channel = (function() {

  const ACTION_JOIN = 0
  const ACTION_LEAVE = 1
  const ACTION_MESSAGE = 2

  const AUTHENTICATED = 3
  const AUTHENTICATION_REQUIRED = 4
  const AUTHENTICATION_FAILED = 5
  const LOGGED_OUT = 6

  var defaultOptions = {
    id: null,
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
    this.id = _uniqueId();

    /** @type {String} */
    this.url = null

    /** @type {Object} */
    this.options = _extend(defaultOptions, options);

    if (this.options.id != null) {
      this.id = this.options.id;
    }

    if (this.options.url != null) {
      this.url = this.options.url;
    }

    /** @type {Object} */
    this.credential = _extend(defaultCredential, credential);

    /** @type {Object} */
    this._handlers = [];
  }

  /**
   * @param {StreamEvent} event
   * @param {Function} callback
   */
  Channel.prototype.registerHandler = function(options, callback) {
    options = {
      id: null,
      action_group: "feed|entry",
      action_type: "add|del|update|*"
    }

    this._handlers.push({options: options, cb: callback});
  }

  /**
   * @param {StreamEvent} event
   */
  Channel.prototype.onData = function(event) {
    for (var i in this._handlers) {
      this._handlers[i].cb.call(this, data);
    }

    this.EventToString(event);
  }

  Channel.prototype.Authenticate = function(credential) {
  }

  Channel.prototype.getConnection = function() {
  }

  Channel.prototype.getWebSocketConnection = function() {
    this._socket = new WebSocket('ws://localhost:10100/ws/join?chid=' + this.id);

    self = this
    this._socket.onmessage = function(event) {
      event = new StreamEvent(JSON.parse(event.data))
      self.onData(event)
    };

    self = this
    return {
      send: function(data) {
        self._socket.send(JSON.stringify(data))
      }
    };
  }

  Channel.prototype.getLongPoolingConnection = function() {
    var lastReceived = 0;
    var isWait = false;

    this.getJSON('http://localhost:10100/lp/join?chid=' + this.id, function() {
    })

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
          event = new StreamEvent(event)
          self.onData(event)

          lastReceived = event.GetTimestamp();
        });
        isWait = false;
      });
    }

    setInterval(fetch, 3000);
    fetch()

    return {
      send: function(data) {
        self.post("/lp/post", {chid: self.id, data: JSON.stringify(data)}, function(status) {
        });
      }
    };
  }

  Channel.prototype.load = function(url, callback) {
    var xhr;
    if (typeof XMLHttpRequest !== 'undefined') {
      xhr = new XMLHttpRequest();
    } else {
      var versions = ["MSXML2.XmlHttp.5.0", "MSXML2.XmlHttp.4.0", "MSXML2.XmlHttp.3.0", "MSXML2.XmlHttp.2.0", "Microsoft.XmlHttp"];
      for (var i = 0, len = versions.length; i < len; i++) {
        try {
          xhr = new ActiveXObject(versions[i]);
          break;
        }
        catch (e) {
        }
      }
    }
    xhr.onreadystatechange = ensureReadiness;
    function ensureReadiness() {
      if (xhr.readyState < 4) {
        return;
      }
      if (xhr.status !== 200) {
        return;
      }
      if (xhr.readyState === 4) {
        callback(xhr);
      }
    }

    xhr.open('GET', url, true);
    xhr.send('');
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

  Channel.prototype.EventToString = function(event) {
    switch (event.Type) {
      case ACTION_JOIN:
        console.log(event.User + " joined the chat room");
        break;
      case ACTION_LEAVE:
        console.log(event.User + " left the chat room");
        break;
      case ACTION_MESSAGE:
        console.log(event.User + ", " + event.PrintContent());
        break;
    }
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

  var _uniqueId = function() {
    return '_' + Math.random().toString(36).substr(2, 36);
  }

  return Channel;

})();
