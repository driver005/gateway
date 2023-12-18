package config

import (
	"github.com/spf13/viper"
)

func Init() {
	//Stripe
	// viper.SetDefault("stripe.secret_key", "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")
	viper.SetDefault("stripe.secret_key", "sk_test_51MJoIJDuTp4RMkBZ7gkfWY0Un5XLA60A4cpkn86lARUhCuflEj2BIN0iD80J7mRwZCQodp0azQ4BC5Wd71nZxOEF00g9d402Tk")
	viper.SetDefault("stripe.public_key", "pk_live_51MJoIJDuTp4RMkBZQ0AcBuEmHGbIjWqwHcwHDdDBPMLevQcUk8OQnXWjJthdwKsJKF8LrfApwaiv1MrcV4ShhAHk00ZbuZp49x")

	// Logger Defaults
	viper.SetDefault("logger.level", "debug")
	viper.SetDefault("logger.encoding", "console")
	viper.SetDefault("logger.color", true)
	viper.SetDefault("logger.dev_mode", true)
	viper.SetDefault("logger.disable_caller", false)
	viper.SetDefault("logger.disable_stacktrace", false)

	// Pidfile
	viper.SetDefault("pidfile", "")

	// Profiler config
	viper.SetDefault("profiler.enabled", false)
	viper.SetDefault("profiler.host", "")
	viper.SetDefault("profiler.port", "6060")

	// Server Configuration
	viper.SetDefault("server.container", false)
	viper.SetDefault("server.production", false)
	viper.SetDefault("server.host", "")
	viper.SetDefault("server.port", "3002")
	viper.SetDefault("server.tls", false)
	viper.SetDefault("server.devcert", false)
	viper.SetDefault("server.certfile", "server.crt")
	viper.SetDefault("server.keyfile", "server.key")
	viper.SetDefault("server.log_requests", true)
	viper.SetDefault("server.log_cors", false)
	viper.SetDefault("server.profiler_enabled", false)
	viper.SetDefault("server.profiler_path", "/debug")
	viper.SetDefault("server.sentry", "")
	viper.SetDefault("server.mock.authorizer", true)

	// gorm Settings
	viper.SetDefault("storage.type", "graphql")
	viper.SetDefault("storage.endpoint", "")

	// Merchant Settings -- Apple Pay
	viper.SetDefault("apple.merchant.id", "merchant.com.buytecheckout")
	viper.SetDefault("apple.merchant.name", "Buyte Apple Pay Checkout")
	viper.SetDefault("apple.merchant.domain", "go.buytecheckout.com")
	viper.SetDefault("apple.certs", "")

	// Merchant Settings -- Google Pay
	viper.SetDefault("google.merchant.id", "05174216476243863888")
	viper.SetDefault("google.merchant.name", "Buyte Google Pay Checkout")
	viper.SetDefault("google.merchant.domain", "go.buytecheckout.com")

	// Lambda Functions Settings
	viper.SetDefault("func.region", "ap-southeast-2")
	viper.SetDefault("func.adyen_cse", "buyte-dev-adyen_cse")

	// Proxy Settings
	viper.SetDefault("proxy.token", "06935288689c5a2935575744fcacff3e6bc92900")

	// User Pool
	viper.SetDefault("config.clientId", "")
	viper.SetDefault("config.userPoolId", "")

	// Stripe Settings
	viper.SetDefault("stripe.live.secret", "")
	viper.SetDefault("stripe.live.public", "")
	viper.SetDefault("stripe.test.secret", "")
	viper.SetDefault("stripe.test.public", "")
}
