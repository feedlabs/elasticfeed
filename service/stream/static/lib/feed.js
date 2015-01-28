var Feed = (function() {

  const ACTION_GROUP_TYPE = 1

  const ACTION_RELOAD = 1
  const ACTION_RESET = 2
  const ACTION_DATA_INIT = 3
  const ACTION_DATA_MORE = 4
  const ACTION_HIDE = 5
  const ACTION_SHOW = 6

  const AUTHENTICATED = 100
  const AUTHENTICATION_REQUIRED = 101
  const AUTHENTICATION_FAILED = 102
  const LOGGED_OUT = 103

  /** @type {Feed} */
  var localCache = {}

  /** @type {Object} */
  var globalOptions = {
    feedId: '',
    outputContainerId: 'defaultContainerId',
    defaultElementLayout: '',
    defaultElementCount: 0
  };

  /** @type {Object} */
  var globalCredential = {
    username: null,
    token: null,
    method: 'basic'
  };

  function Feed(options, channel) {

    /** @type {String} */
    this.id = null;

    /** @type {Channel} */
    this.channel = channel;

    /** @type {Array} */
    this.entryList = [];

    /** @type {Object} */
    if (this.channel.options.transport == 'ws') {
      this.socket = this.channel.getWebSocketConnection();
    } else if (this.channel.options.transport == 'lp') {
      this.socket = this.channel.getLongPoolingConnection();
    }

    this.options = _extend(globalOptions, options);
    this._stylerFunction = options.styler || this._stylerFunction;
    this.outputContainer = document.getElementById(this.options.outputContainerId);

    this.bindChannel(this.channel);
  }

  Feed.prototype.on = function(type, callback) {

  }

  // Events callbacks

  Feed.prototype.onReload = function(callback) {
  }

  Feed.prototype.onReset = function(callback) {
  }

  Feed.prototype.onEntryAdd = function(callback) {
  }

  Feed.prototype.onEntryDelete = function(callback) {
  }

  Feed.prototype.onEntryUpdate = function(callback) {
  }

  Feed.prototype.onEvent = function(eventName, callback) {
  }

  Feed.prototype.onData = function(callback) {
  }

  // Entries management

  Feed.prototype.addEntry = function(data) {
    this.entryList.push(new Entry(data))
  }

  Feed.prototype.deleteEntry = function(id) {
  }

  Feed.prototype.updateEntry = function(id, data) {
  }

  Feed.prototype.findEntry = function(id) {
  }

  // Handlers

  Feed.prototype.bindChannel = function(channel) {
    channel.on('message', function(chid, ts, data) {
      // should detect type of message
      // if feed addressed then check id
      // trigger action if needed
    });
  }

  // Stylers

  Feed.prototype._stylerFunction = function(data) {
    return JSON.stringify(data.Data);
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
