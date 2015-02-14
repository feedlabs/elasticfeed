#-------------------------------------------------------------------------
# Configure Middleman
#-------------------------------------------------------------------------

set :base_url, "https://www.elasticfeed.io/"

activate :feedlabs do |h|
  h.version      = ENV["ELASTICFEED_VERSION"]
  h.bintray_repo = 'elasticfeed/elasticfeed'
  h.bintray_user = 'elasticfeed'
  h.bintray_key  = ENV['BINTRAY_API_KEY']
end
