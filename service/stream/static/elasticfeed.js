function includeJs(jsFilePath) {
  var js = document.createElement("script");

  js.type = "text/javascript";
  js.src = jsFilePath;

  document.body.appendChild(js);
}

includeJs('lib/feed.js');
includeJs('lib/entry.js');
includeJs('lib/channel.js');
includeJs('lib/event/channel.js');
includeJs('lib/event/system.js');

(function(window) {

  /** @type {Object} */
  var defaultOptions = {
    channel: {
      url: 'localhost',
      transport: 'ws'
    },
    styler: function(data) {
    }
  }

  var elasticfeed = {

    /** @type {Object} */
    options: {},

    /** @type {Object} */
    channelList: {},

    /** @type {Object} */
    feedList: {},

    init: function(options) {
      this.options = _extend(defaultOptions, options);
    },

    initFeed: function(id, options) {
      feed_options = _extend(this.options, options || {});
      channel = this.getChannel(feed_options.channel);

      return new Feed(feed_options, channel);
    },

    /**
     * Returns Feed object
     * @param options
     * @returns {*}
     */
    getFeed: function(options) {
      if (options.id == undefined) {
        return null;
      }

      if (this.feedList[id] == undefined) {
        this.feedList[id] = new Feed(options)
      }

      return this.feedList[options.id];
    },

    /**
     * Returns Channel defined per API url
     * @param options
     * @param credential
     * @returns {*}
     */
    getChannel: function(options, credential) {
      if (options.url == undefined) {
        return null;
      }

      if (this.channelList[options.url] == undefined) {
        this.channelList[options.url] = new Channel(options, credential)
      }

      return this.channelList[options.url];
    },

    findFeed: function(id) {
      return this.getFeed({id: id});
    },

    findChannel: function(url) {
      return this.getChannel({url: url});
    }

  };

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

  if ("function" === typeof define) {
    define(function(require) {
      return elasticfeed;
    });
  } else {
    window.elasticfeed = elasticfeed;
  }

}(window));

// Helpers

Element.prototype.remove = function() {
  this.parentElement.removeChild(this);
};

NodeList.prototype.remove = HTMLCollection.prototype.remove = function() {
  for (var i = 0, len = this.length; i < len; i++) {
    if (this[i] && this[i].parentElement) {
      this[i].parentElement.removeChild(this[i]);
    }
  }
};
