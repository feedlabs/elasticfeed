node default {

  include 'monit'

  class {'neo4j': }

  class {'nsq::lookup': }

  class {'nsq::producer':
    nsqlookupd_tcp_addresses => ['127.0.0.1:4160'],
    require => Class['nsq::lookup'],
  }

  class {'nsq::admin':
    nsqlookupd_http_addresses => ['127.0.0.1:4161'],
  }

  nsq::glue::file {'archive':
    topic => 'archive',
    nsqlookupd_http_addresses => '127.0.0.1:4161',
  }

  nsq::glue::file {'test':
    topic => 'test',
    nsqlookupd_http_addresses => '127.0.0.1:4161',
  }

}
