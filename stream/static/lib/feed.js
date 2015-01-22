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

  Feed.prototype.addEntry = function() {
    this.entryList.push(new Entry())
  }

  Feed.prototype.findEntry = function(id) {

  }

  Feed.prototype.getChannel = function() {

  }

  Feed.prototype._stylerFunction = function(data) {
    return JSON.stringify(data.Data);
  }

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
