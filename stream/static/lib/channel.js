var Channel = (function() {

  function Channel() {

    /** @type {String} */
    this.id = null;
  }

  Channel.prototype.registerFeedHandler = function(feed, callback) {

  }

  Channel.prototype.onData = function(data) {
    // get handlers
    event = new StreamEvent(data)
  }

  Channel.prototype.getConnection = function() {

  }

  Channel.prototype.getWebSocketConnection = function() {

  }

  Channel.prototype.getLongPoolingConnection = function() {

  }

  return Channel;

})();
