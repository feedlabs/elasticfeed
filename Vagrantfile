Vagrant.configure('2') do |config|
  config.ssh.forward_agent = true
  config.vm.box = 'cargomedia/debian-7-amd64-cm'

  config.vm.hostname = 'api.feedlabs.dev'
  if Vagrant.has_plugin? 'vagrant-dns'
    config.dns.tld = 'dev'
    config.dns.patterns = [/^.*feedlabs.dev$/]
  end

  synced_folder_type = ENV.fetch('SYNC_TYPE', 'nfs')
  synced_folder_type = nil if 'vboxsf' == synced_folder_type

  config.vm.network :private_network, ip: '10.10.11.10'
  config.vm.synced_folder '.', '/home/vagrant/feedlabs', :type => 'nfs'

  config.librarian_puppet.puppetfile_dir = 'puppet'
  config.librarian_puppet.placeholder_filename = '.gitkeep'
  config.librarian_puppet.resolve_options = {:force => true}
  config.vm.provision :puppet do |puppet|
    puppet.module_path = 'puppet/modules'
    puppet.manifests_path = 'puppet/manifests'
  end

  config.vm.provision 'shell', inline: [
    'cd /home/vagrant/feedlabs/src'
  ].join(' && ')
end
