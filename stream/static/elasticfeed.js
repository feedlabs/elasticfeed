/*
 * Author: Feed Labs
 */

function includeJs(jsFilePath) {
  var js = document.createElement("script");

  js.type = "text/javascript";
  js.src = jsFilePath;

  document.body.appendChild(js);
}

includeJs('lib/feed.js');
includeJs('lib/entry.js');
includeJs('lib/channel.js');
includeJs('lib/event.js');

(function(window) {

  var elasticfeed = {

    /** @type {Object} */
    channel: null,

    getFeed: function(id) {
      return Feed;
    },

    getChannel: function(id) {
      if(this.channel == null) {
        this.channel = new Channel()
      }

      return this.channel;
    },

    newFeed: function(options) {

    },

    newChannel: function(options) {

    },

    _uniqueId: function() {
      return '_' + Math.random().toString(36).substr(2, 9);
    },

    load: function(url, callback) {
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

  }; // elasticfeed

  // ===========================================================================

  if ("function" === typeof define) {
    define(function(require) {
      return elasticfeed;
    });
  } else {
    window.elasticfeed = elasticfeed;
  }
}(window));


// Helpers
// =======

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
