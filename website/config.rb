#-------------------------------------------------------------------------
# Configure Middleman
#-------------------------------------------------------------------------

set :base_url, "https://www.elasticfeed.io/"

set :latest_version, "1.0"

activate :hashicorp do |h|
  h.version      = ENV["ELASTICFEED_VERSION"]
  h.bintray_repo = 'elasticfeed/elasticfeed'
  h.bintray_user = 'elasticfeed'
  h.bintray_key  = ENV['BINTRAY_API_KEY']
end
