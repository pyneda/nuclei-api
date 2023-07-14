package config

import (
	"github.com/spf13/viper"
)

func SetDefaultConfig() {
	// follow-redirects: enable following redirects for http templates
	viper.SetDefault("configs.follow-redirects", false)
	// follow-host-redirects: follow redirects on the same host
	viper.SetDefault("configs.follow-host-redirects", false)
	// max-redirects: max number of redirects to follow for http templates
	viper.SetDefault("configs.max-redirects", 10)
	// disable-redirects: disable redirects for http templates
	viper.SetDefault("configs.disable-redirects", false)
	// report-config: nuclei reporting module configuration file
	// viper.SetDefault("configs.report-config", "")
	// header: custom header/cookie to include in all http request in header:value format (cli, file)
	viper.SetDefault("configs.header", []string{})
	// var: custom vars in key=value format
	// viper.SetDefault("configs.var", map[string]string{})
	// resolvers: file containing resolver list for nuclei
	viper.SetDefault("configs.resolvers", "")
	// system-resolvers: use system DNS resolving as error fallback
	viper.SetDefault("configs.system-resolvers", false)
	// disable-clustering: disable clustering of requests
	viper.SetDefault("configs.disable-clustering", false)
	// offline-http: enable passive HTTP response processing mode
	viper.SetDefault("configs.offline-http", false)
	// force-http2: force http2 connection on requests
	viper.SetDefault("configs.force-http2", false)
	// env-vars: enable environment variables to be used in template
	viper.SetDefault("configs.env-vars", false)
	// client-cert: client certificate file (PEM-encoded) used for authenticating against scanned hosts
	viper.SetDefault("configs.client-cert", "")
	// client-key: client key file (PEM-encoded) used for authenticating against scanned hosts
	viper.SetDefault("configs.client-key", "")
	// client-ca: client certificate authority file (PEM-encoded) used for authenticating against scanned hosts
	viper.SetDefault("configs.client-ca", "")
	// show-match-line: show match lines for file templates, works with extractors only
	viper.SetDefault("configs.show-match-line", false)
	// ztls: use ztls library with autofallback to standard one for tls13
	viper.SetDefault("configs.ztls", false)
	// sni: tls sni hostname to use (default: input domain name)
	viper.SetDefault("configs.sni", "")
	// sandbox: sandbox nuclei for safe templates execution
	viper.SetDefault("configs.sandbox", false)
	// interface: network interface to use for network scan
	viper.SetDefault("configs.interface", "")
	// attack-type: type of payload combinations to perform (batteringram,pitchfork,clusterbomb)
	viper.SetDefault("configs.attack-type", "")

	// interactsh default configuration
	viper.SetDefault("interactsh.enabled", true)
	viper.SetDefault("interactsh.url", "")
	viper.SetDefault("interactsh.token", "")
	viper.SetDefault("interactsh.cache-size", 5000)
	viper.SetDefault("interactsh.eviction", 60)
	viper.SetDefault("interactsh.poll-duration", 5)
	viper.SetDefault("interactsh.cooldown-period", 5)

	// rate limit configuration
	viper.SetDefault("rate-limit.rate-limit", 150)
	viper.SetDefault("rate-limit.rate-limit-minute", 0)
	viper.SetDefault("rate-limit.bulk-size", 25)
	viper.SetDefault("rate-limit.template-threads", 25)
	viper.SetDefault("rate-limit.headless-bulk-size", 10)
	viper.SetDefault("rate-limit.headless-template-threads", 10)

	// uncover configuration
	viper.SetDefault("uncover.enabled", false)
	viper.SetDefault("uncover.uncover-query", []string{})
	viper.SetDefault("uncover.uncover-engine", []string{})
	viper.SetDefault("uncover.uncover-field", "ip:port")
	viper.SetDefault("uncover.uncover-limit", 100)
	viper.SetDefault("uncover.uncover-ratelimit", 60)

}
