package config

// SourceConfig 数据源配置
type SourceConfig struct {
	BufferSize int32            `mapstructure:"buffer-size"`
	FileSource []*LogFileSource `mapstructure:"file-source"`
}

// LogFileSource 日志路径
type LogFileSource struct {
	Multiline *Multiline `mapstructure:"multiline"`
	AppId     string     `mapstructure:"app-id"`
	Path      string     `mapstructure:"path"`
}

type Multiline struct {
	Pattern string `mapstructure:"pattern"`
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
	LogLevel    string             `mapstructure:"log-level"`
	Source      *SourceConfig      `mapstructure:"source"`
	Destination *DestinationConfig `mapstructure:"destination"`
}
