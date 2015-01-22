var Channel = (function() {

  var globalCredential = {
    username: null,
    token: null,
    method: 'basic'
  }

  function Channel(credential) {

    /** @type {String} */
    this.id = null;

    /** @type {Object} */
    this.credential = _extend(globalCredential, credential);
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

  Channel.prototype.getWebSocketConnection = function() {
  }

  Channel.prototype.getLongPoolingConnection = function() {
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
