package config

// SourceConfig 数据源配置
type SourceConfig struct {
	Type string     `mapstructure:"type"`
	Path []*LogPath `mapstructure:"path"`
}

// LogPath 日志路径
type LogPath struct {
	AppId string `mapstructure:"app-id"`
	Path  string `mapstructure:"path"`
}

// DestinationConfig 接收源配置
type DestinationConfig struct {
	Kafka *KafkaConfig `mapstructure:"kafka"`
}

// KafkaConfig Kafka配置
type KafkaConfig struct {
	BrokerList []string `mapstructure:"broker-list"`
	TopicName  string   `mapstructure:"topic-name"`
}

// CollectorConfig 日志收集器配置
type CollectorConfig struct {
	Source      *SourceConfig      `mapstructure:"source"`
	Destination *DestinationConfig `mapstructure:"destination"`
}
