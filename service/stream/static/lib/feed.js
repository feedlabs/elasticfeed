var Feed = (function() {

  const SYSTEM_FEED_MESSAGE = 1

  const RELOAD = 1
  const EMPTY = 2
  const ENTRY_NEW = 3
  const ENTRY_INIT = 4
  const ENTRY_MORE = 5
  const HIDE = 6
  const SHOW = 7
  const ENTRY_MESSAGE = 8

  const AUTHENTICATED = 100
  const AUTHENTICATION_REQUIRED = 101
  const AUTHENTICATION_FAILED = 102
  const LOGGED_OUT = 103

  /** @type {Feed} */
  var localCache = {}

  var globalOptions = {

    /** @type {Function} */
    stylerFunction: function(data) {
      return JSON.stringify(data);
    },

    /** @type {Function} */
    renderFunction: function(data) {
      return JSON.stringify(data);
    }
  };

  var globalCredential = {

    /** @type {String} */
    username: null,

    /** @type {String} */
    token: null,

    /** @type {String} */
    method: 'basic'
  };

  function Feed(id, options, channel) {

    /** @type {String} */
    this.id = id;

    /** @type {String} */
    this.feedId = id.split(/[::]/)[0];

    /** @type {String} */
    this.appId = id.split(/[::]/)[1];

    /** @type {String} */
    this.orgId = id.split(/[::]/)[2];

    /** @type {Channel} */
    this.channel = channel;

    /** @type {Array} */
    this.entryList = [];

    this.loadInit();

    /** @type {Object} */
    if (this.channel.options.transport == 'ws') {
      this.socket = this.channel.getWebSocketConnection();
    } else if (this.channel.options.transport == 'lp') {
      this.socket = this.channel.getLongPoolingConnection();
    }

    /** @type {Object} */
    this.options = _extend(globalOptions, options);

    /** @type {Function} */
    this.stylerFunction = this.options.stylerFunction;

    /** @type {Function} */
    this.renderFunction = this.options.renderFunction;

    /** @type {DOM} */
    this.outputContainer = document.getElementById(this.options.outputContainerId);

    this.bindChannel(this.channel);

    /** @type {Object} */
    this._handlers = {};

    /** @type {Object} */
    this._state = {
      initiated: false
    };
  }

  Feed.prototype.on = function(name, callback) {
    switch (name) {
      case 'reload':
        type = RELOAD
        break;
      case 'empty':
        type = EMPTY
        break;
      case 'entry':
        type = ENTRY_NEW
        break;
      case 'entry-init':
        type = ENTRY_INIT
        break;
      case 'entry-more':
        type = ENTRY_MORE
        break;
      case 'hide':
        type = HIDE
        break;
      case 'show':
        type = SHOW
        break;
      case 'entry-message':
        type = ENTRY_MESSAGE
        break;
      case 'authenticated':
        type = AUTHENTICATED
        break;
      case 'authentication-required':
        type = AUTHENTICATION_REQUIRED
        break;
      case 'authentication-failed':
        type = AUTHENTICATION_FAILED
        break;
      case 'logout':
        type = LOGGED_OUT
        break;
      default:
        break;
    }
    if (this._handlers[type] == undefined) {
      this._handlers[type] = []
    }
    this._handlers[type].push(callback);

    return callback;
  }

  Feed.prototype.off = function(callback) {
    for (var i in this._handlers) {
      for (var x in this._handlers[i]) {
        if (this._handlers[i][x] == callback) {
          delete this._handlers[i][x];
          return;
        }
      }
    }
  }

  Feed.prototype.onData = function(feedEvent) {
    switch (feedEvent.type) {
      case RELOAD:
        this.onReload(feedEvent.ts)
        break;
      case EMPTY:
        this.onEmpty(feedEvent.ts)
        break;
      case ENTRY_NEW:
        this.onEntryNew(feedEvent.ts, feedEvent.content)
        break;
      case ENTRY_INIT:
        this.onEntryInit(feedEvent.ts, feedEvent.content)
        break;
      case ENTRY_MORE:
        this.onEntryMore(feedEvent.ts, feedEvent.content)
        break;
      case HIDE:
        this.onHide(feedEvent.ts)
        break;
      case SHOW:
        this.onShow(feedEvent.ts)
        break;
      case ENTRY_MESSAGE:
        this.onEntryMessage(feedEvent.ts, feedEvent.content)
        break;
      case AUTHENTICATED:
        this.onAuthenticated(feedEvent.ts, feedEvent.content)
        break;
      case AUTHENTICATION_REQUIRED:
        this.onAuthenticationRequired(feedEvent.ts, feedEvent.content)
        break;
      case AUTHENTICATION_FAILED:
        this.onAuthenticationFailed(feedEvent.ts, feedEvent.content)
        break;
      case LOGGED_OUT:
        this.onLogout(feedEvent.ts, feedEvent.content)
        break;
    }
  }

  // Events callbacks

  Feed.prototype.onReload = function(timestamp) {
    for (var i in this._handlers[RELOAD]) {
      this._handlers[RELOAD][i].call(this, timestamp);
    }
  }

  Feed.prototype.onEmpty = function(timestamp) {
    for (var i in this._handlers[EMPTY]) {
      this._handlers[EMPTY][i].call(this, timestamp);
    }
  }

  Feed.prototype.onEntryNew = function(timestamp, data) {
    entry = new Entry(data, {styler: this.stylerFunction});

    this.addEntry(entry)

    for (var i in this._handlers[ENTRY_NEW]) {
      this._handlers[ENTRY_NEW][i].call(this, timestamp, entry);
    }
  }

  Feed.prototype.onEntryInit = function(timestamp, entries) {
    for (var i in entries) {
      this.onEntryNew(timestamp, entries[i]);
    }

    for (var i in this._handlers[ENTRY_INIT]) {
      this._handlers[ENTRY_INIT][i].call(this, timestamp, entries);
    }
  }

  Feed.prototype.onEntryMore = function(timestamp, data) {
    entries = JSON.parse(data);

    for (var i in this._handlers[ENTRY_MORE]) {
      this._handlers[ENTRY_MORE][i].call(this, timestamp, entries);
    }
  }

  Feed.prototype.onHide = function(timestamp) {
    for (var i in this._handlers[HIDE]) {
      this._handlers[HIDE][i].call(this, timestamp);
    }
  }

  Feed.prototype.onShow = function(timestamp) {
    for (var i in this._handlers[SHOW]) {
      this._handlers[SHOW][i].call(this, timestamp);
    }
  }

  Feed.prototype.onEntryMessage = function(timestamp, content) {
    entryEvent = new Event(content);

    for (var i in this._handlers[ENTRY_MESSAGE]) {
      this._handlers[ENTRY_MESSAGE][i].call(this, timestamp, entryEvent);
    }
  }

  Feed.prototype.onAuthenticated = function(timestamp, content) {
    for (var i in this._handlers[AUTHENTICATED]) {
      this._handlers[AUTHENTICATED][i].call(this, timestamp);
    }
  }

  Feed.prototype.onAuthenticationRequired = function(timestamp, content) {
    for (var i in this._handlers[AUTHENTICATION_REQUIRED]) {
      this._handlers[AUTHENTICATION_REQUIRED][i].call(this, timestamp);
    }
  }

  Feed.prototype.onAuthenticationFailed = function(timestamp, content) {
    for (var i in this._handlers[AUTHENTICATION_FAILED]) {
      this._handlers[AUTHENTICATION_FAILED][i].call(this, timestamp);
    }
  }

  Feed.prototype.onLogout = function(timestamp, content) {
    for (var i in this._handlers[LOGGED_OUT]) {
      this._handlers[LOGGED_OUT][i].call(this, timestamp);
    }
  }

  // Feed management

  Feed.prototype.reload = function() {
    this.empty();
    this.socket.send({action: ENTRY_INIT, feedId: this.feedId, appId: this.appId, orgId: this.orgId});
  }

  Feed.prototype.loadMore = function() {
    this.socket.send({action: ENTRY_MORE, feedId: this.feedId, appId: this.appId, orgId: this.orgId, state: {}});
  }

  Feed.prototype.loadInit = function() {
    var self = this;
    this.channel.on('join', function() {
      if (self._state.initiated == true) {
        return;
      }

      self.socket.send({action: ENTRY_INIT, feedId: self.feedId, appId: self.appId, orgId: self.orgId});
      self._state.initiated = true;
    });
  }

  // Entries management

  Feed.prototype.addEntry = function(entry) {

    // types
    // add by: timestamp up/down; always to top; always to bottom

    entry.setParent(this);
    this.entryList.push(entry);

    this.outputContainer.innerHTML = '<div id="' + entry.getViewId() + '"></div>' + this.outputContainer.innerHTML;

    entry.render();
  }

  Feed.prototype.deleteEntry = function(entry) {
    entry.delete();
  }

  Feed.prototype.updateEntry = function(entry, data) {
  }

  Feed.prototype.empty = function() {
    for (var i in this.entryList) {
      this.deleteEntry(this.entryList[i]);
      delete this.entryList[i];
    }
    this.entryList = []
  }

  Feed.prototype.findEntry = function(id) {
  }

  // UI

  Feed.prototype.render = function(id) {
    for (var i in this.entryList) {
      this.entryList[i].render();
    }
  }

  // Handlers

  Feed.prototype.bindChannel = function(channel) {
    var self = this;
    channel.on('message', function(chid, ts, systemEvent) {
      if (systemEvent.type == SYSTEM_FEED_MESSAGE) {
        feedEvent = new Event(systemEvent.content);
        if (feedEvent.user == self.feedId || feedEvent.user == '*') {
          self.onData(feedEvent);
        }
      }
    });
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

  return Feed;

})();
