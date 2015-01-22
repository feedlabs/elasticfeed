var Feed = (function() {

  /** @type {Feed} */
  var localCache = {}

  /** @type {Object} */
  var globalOptions = {
    feedId: '',
    outputContainerId: 'defaultContainerId',
    defaultElementLayout: '',
    defaultElementCount: 0
  };

  function Feed(options, stylerFunction) {

    /** @type {String} */
    this.id = null;

    /** @type {String} */
    this.channelId = null;

    /** @type {Array} */
    this.entryList = [];

    this.options = _extend(globalOptions, options);
    this._stylerFunction = stylerFunction || this._stylerFunction;
    this.outputContainer = document.getElementById(this.options.outputContainerId);
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

  Feed.prototype.registerHandlers = function() {
    // bind to channel data
  }

  // Channel management

  Feed.prototype.getChannel = function() {
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
