package config

// ConfigFile returns the default path to the configuration file. On
// Unix-like systems this is the ".elasticfeedconfig" file in the home directory.
// On Windows, this is the "elasticfeed.config" file in the application data
// directory.
func ConfigFile() (string, error) {
	return configFile()
}

// ConfigDir returns the configuration directory for Elasticfeed.
func ConfigDir() (string, error) {
	return configDir()
}
