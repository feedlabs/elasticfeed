var Channel = (function() {

  const JOIN = 0
  const LEAVE = 1
  const MESSAGE = 2

  var defaultOptions = {
    id: null,
    transport: 'ws',
    connectOnInit: true
  }

  function Channel(options) {

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
    this._handlers = {};

    /** @type {WebSocket} */
    this._socket = null;

    this._xhr = []
  }

  // Handlers

  /**
   * @param {Event} event
   * @param {Function} callback
   */
  Channel.prototype.on = function(name, callback) {
    switch (name) {
      case 'join':
        type = JOIN
        break;
      case 'leave':
        type = LEAVE
        break;
      case 'message':
        type = MESSAGE
        break;
      default:
        return false;
        break;
    }
    if (this._handlers[type] == undefined) {
      this._handlers[type] = []
    }
    this._handlers[type].push(callback);
    return true;
  }

  // Events

  /**
   * @param {Event} event
   */
  Channel.prototype.onData = function(event) {
    switch (event.type) {
      case JOIN:
        this.onJoin(event.user, event.ts)
        break;
      case LEAVE:
        this.onLeave(event.user, event.ts)
        break;
      case MESSAGE:
        this.onMessage(event.user, event.ts, event.content)
        break;
    }
  }

  Channel.prototype.onJoin = function(chid, timestamp) {
    for (var i in this._handlers[JOIN]) {
      this._handlers[JOIN][i].call(this, chid, timestamp);
    }
  }

  Channel.prototype.onLeave = function(chid, timestamp) {
    for (var i in this._handlers[LEAVE]) {
      this._handlers[LEAVE][i].call(this, chid, timestamp);
    }
  }

  Channel.prototype.onMessage = function(chid, timestamp, data) {
    systemEvent = new Event(data);

    for (var i in this._handlers[MESSAGE]) {
      this._handlers[MESSAGE][i].call(this, chid, timestamp, systemEvent);
    }
  }

  // Connection

  Channel.prototype.isWebSocket = function() {
    return this._socket != undefined;
  }

  Channel.prototype.getConnection = function() {
  }

  Channel.prototype.getWebSocketConnection = function() {
    var self = this;

    if (this._socket == null) {
      this._socket = new WebSocket('ws://localhost:10100/stream/ws/join?chid=' + this.id);

      this._socket.onmessage = function(event) {
        event = new Event(JSON.parse(event.data))
        self.onData(event)
      };
    }

    return {
      send: function(data) {
        self._socket.send(JSON.stringify(data));
      }
    };
  }

  Channel.prototype.getLongPoolingConnection = function() {

    self = this;

    if (this._socket == null) {
      var lastReceived = 0;
      var isWait = false;

      this.getJSON('http://localhost:10100/stream/lp/join?chid=' + this.id, function(data) {
        if (data == null) {
          return;
        }
        event = new Event(data['response']);
        self.onData(event);
      })

      var fetch = function() {
        if (isWait) {
          return;
        }
        isWait = true;
        self.getJSON("http://localhost:10100/stream/lp/fetch?lastReceived=" + lastReceived, function(data, code) {

          if (code == 4) {
            isWait = false
          }

          if (data == null) {
            return;
          }

          self.each(data, function(i, event) {
            event = new Event(event)
            self.onData(event)

            lastReceived = event.GetTimestamp();
          });
          isWait = false;
        });
      }

      this._socket = setInterval(fetch, 3000);
      fetch();
    }

    return {
      send: function(data) {
        self.post("/stream/lp/post", {chid: self.id, data: JSON.stringify(data)}, function(data) {
          response_json = JSON.parse(data);
          event = new Event(JSON.parse(response_json['response']));
          self.onData(event);
        });
      }
    };
  }

  Channel.prototype.geServerSentEventsConnection = function() {

    es = new EventSource("http://localhost:8001/stream/sse/join");
    es.onmessage = function(event) {
    };

    es.addEventListener("some-event", function(event) {
    }, false);

    return {
      send: function(data) {
        self.post("/stream/sse/post", {chid: self.id, data: JSON.stringify(data)}, function(data) {
          response_json = JSON.parse(data);
          event = new Event(JSON.parse(response_json['response']));
          self.onData(event);
        });
      }
    }
  }

  // HTTP

  Channel.prototype.__cleanup = function() {
    for (var i in this._xhr) {
      if (this._xhr[i].xhr.readyState == 4) {
        delete this._xhr[i];
      }
    }
  }

  Channel.prototype.getJSON = function(url, callback) {

    this.__cleanup();

    var pos = this._xhr.length;

    this._xhr[pos] = {
      xhr: new XMLHttpRequest(),
      cb: callback
    };

    var self = this;
    this._xhr[pos].xhr.onreadystatechange = function() {
      if (self._xhr[pos].xhr.readyState == 4 && self._xhr[pos].xhr.status == 200) {
        if (self._xhr[pos].xhr.responseText != "") {
          data = JSON.parse(self._xhr[pos].xhr.responseText);
          self._xhr[pos].cb.call(this, data, self._xhr[pos].xhr.readyState)
        } else {
          self._xhr[pos].cb.call(this, null, self._xhr[pos].xhr.readyState)
        }
      } else {
        self._xhr[pos].cb.call(this, null, self._xhr[pos].xhr.readyState)
      }
    }
    this._xhr[pos].xhr.open("GET", url, true)
    this._xhr[pos].xhr.send('');
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

    this.__cleanup();

    var pos = this._xhr.length;

    this._xhr[pos] = {
      xhr: new XMLHttpRequest(),
      cb: callback
    };

    var self = this;
    this._xhr[pos].xhr.onreadystatechange = function() {
      if (self._xhr[pos].xhr.readyState == 4 && self._xhr[pos].xhr.status == 200) {
        self._xhr[pos].cb.call(this, self._xhr[pos].xhr.responseText)
      }
    }
    dataString = this.queryString(data)
    this._xhr[pos].xhr.open("POST", url + "?" + dataString, true);
    this._xhr[pos].xhr.send(dataString);
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
