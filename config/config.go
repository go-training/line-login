package config

type (
	// Config stores the configuration settings.
	Config struct {
		HTTP struct {
			Host string
			Port string `envconfig:"PORT" default:"8080"`
			Root string `default:"/"`
		}
		Line struct {
			ChannelID     string
			ChannelSecret string
		}
	}
)
