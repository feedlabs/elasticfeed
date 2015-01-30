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
includeJs('lib/event/feed.js');
includeJs('lib/event/entry.js');

(function(window) {

  /** @type {Object} */
  var defaultOptions = {
    channel: {
      url: 'localhost',
      transport: 'ws'
    },
    stylerFunction: function(data) {
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
      if (id == undefined) {
        return false;
      }

      if (this.feedList[id] == undefined) {
        opts = _extend(this.options, options || {});
        channel = this.getChannel(opts.channel);

        this.feedList[id] = new Feed(id, opts, channel);
      }

      return this.feedList[id];
    },

    /**
     * Returns Channel defined per API url
     * @param options
     * @param credential
     * @returns {*}
     */
    getChannel: function(options, credential) {
      if (options.url == undefined) {
        return false;
      }

      if (this.channelList[options.url] == undefined) {
        this.channelList[options.url] = new Channel(options, credential)
      }

      return this.channelList[options.url];
    },

    findFeed: function(id) {
      if (this.feedList[id] == undefined) {
        return false;
      }
      return this.feedList[id];
    },

    findChannel: function(url) {
      if (this.channelList[url] == undefined) {
        return false;
      }
      return this.channelList[url];
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
