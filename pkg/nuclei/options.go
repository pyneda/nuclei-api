package nuclei

import (
	"github.com/projectdiscovery/nuclei/v2/pkg/types"
	"github.com/spf13/viper"
)

func getDefaultOptions() *types.Options {
	options := types.DefaultOptions()
	options.FollowRedirects = viper.GetBool("configs.follow-redirects")
	options.FollowHostRedirects = viper.GetBool("configs.follow-host-redirects")
	options.MaxRedirects = viper.GetInt("configs.max-redirects")
	options.DisableRedirects = viper.GetBool("configs.disable-redirects")
	options.ReportingConfig = viper.GetString("configs.report-config")
	options.CustomHeaders = viper.GetStringSlice("configs.header")
	// options.Vars = viper.GetStringMapString("configs.var")
	options.ResolversFile = viper.GetString("configs.resolvers")
	options.SystemResolvers = viper.GetBool("configs.system-resolvers")
	options.DisableClustering = viper.GetBool("configs.disable-clustering")
	options.OfflineHTTP = viper.GetBool("configs.offline-http")
	options.ForceAttemptHTTP2 = viper.GetBool("configs.force-http2")
	options.EnvironmentVariables = viper.GetBool("configs.env-vars")
	options.ClientCertFile = viper.GetString("configs.client-cert")
	options.ClientKeyFile = viper.GetString("configs.client-key")
	options.ClientCAFile = viper.GetString("configs.client-ca")
	options.ShowMatchLine = viper.GetBool("configs.show-match-line")

	if !viper.GetBool("interactsh.enabled") {
		options.NoInteractsh = true
	}
	options.InteractshURL = viper.GetString("interactsh.url")
	options.InteractshToken = viper.GetString("interactsh.token")
	options.InteractionsCacheSize = viper.GetInt("interactsh.cache-size")
	options.InteractionsEviction = viper.GetInt("interactsh.eviction")
	options.InteractionsPollDuration = viper.GetInt("interactsh.poll-duration")
	options.InteractionsCoolDownPeriod = viper.GetInt("interactsh.cooldown-period")

	options.RateLimit = viper.GetInt("rate-limit.rate-limit")
	options.RateLimitMinute = viper.GetInt("rate-limit.rate-limit-minute")
	options.BulkSize = viper.GetInt("rate-limit.bulk-size")
	options.TemplateThreads = viper.GetInt("rate-limit.template-threads")
	options.HeadlessBulkSize = viper.GetInt("rate-limit.headless-bulk-size")
	options.HeadlessTemplateThreads = viper.GetInt("rate-limit.headless-template-threads")

	options.Uncover = viper.GetBool("uncover.enabled")
	options.UncoverQuery = viper.GetStringSlice("uncover.uncover-query")
	options.UncoverEngine = viper.GetStringSlice("uncover.uncover-engine")
	options.UncoverField = viper.GetString("uncover.uncover-field")
	options.UncoverLimit = viper.GetInt("uncover.uncover-limit")
	options.UncoverRateLimit = viper.GetInt("uncover.uncover-ratelimit")
	return options
}
