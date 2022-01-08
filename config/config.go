package config

// SourceConfig 数据源配置
type SourceConfig struct {
	FileSource []*LogFileSource `mapstructure:"file-source"`
}

// LogFileSource 日志路径
type LogFileSource struct {
	AppId     string     `mapstructure:"app-id"`
	Path      string     `mapstructure:"path"`
	Multiline *Multiline `mapstructure:"multiline"`
	Analyzer  *Analyzer  `mapstructure:"analyzer"`
}

type Multiline struct {
	Enable  bool   `mapstructure:"enable"`
	Pattern string `mapstructure:"pattern"`
}

type Analyzer struct {
	Type    string `mapstructure:"type"`
	Pattern string `mapstructure:"pattern"`
}

// KafkaConfig Kafka配置
type KafkaConfig struct {
	BrokerList []string `mapstructure:"broker-list"`
	TopicName  string   `mapstructure:"topic-name"`
	GroupId    string   `mapstructure:"group-id"`
}

type ElasticConfig struct {
	Url      string `mapstructure:"url"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type ManagerServer struct {
	Port string `mapstructure:"port"`
}

// AppConfig 日志收集器配置
type AppConfig struct {
	LogLevel   string         `mapstructure:"log-level"`   //日志等级
	BufferSize int32          `mapstructure:"buffer-size"` //缓冲对内大小
	Source     *SourceConfig  `mapstructure:"source"`      //日志源
	Kafka      *KafkaConfig   `mapstructure:"kafka"`
	Elastic    *ElasticConfig `mapstructure:"elastic"`
	Manager    *ManagerServer `mapstructure:"manager"`
}
