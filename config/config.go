package config

// SourceConfig 数据源配置
type SourceConfig struct {
	FileSource []*LogFileSource `mapstructure:"file-source"`
}

// LogFileSource 日志路径
type LogFileSource struct {
	AppId     string          `mapstructure:"app-id"`
	Path      string          `mapstructure:"path"`
	Multiline *Multiline      `mapstructure:"multiline"`
	Analyzer  *AnalyzerConfig `mapstructure:"analyzer"`
}

type Multiline struct {
	Enable  bool   `mapstructure:"enable"`
	Pattern string `mapstructure:"pattern"`
}

//日志解析器配置
type AnalyzerConfig struct {
	Parser []*LogParserConfig `mapstructure:"parser"`
}

type LogParserConfig struct {
	AppId   string   `mapstructure:"app-id"`
	Type    string   `mapstructure:"type"`
	Pattern string   `mapstructure:"pattern"`
	Field   []string `mapstructure:"field"`
}

// KafkaConfig Kafka配置
type KafkaConfig struct {
	BrokerList []string `mapstructure:"broker-list"`
	TopicName  string   `mapstructure:"topic-name"`
	GroupId    string   `mapstructure:"group-id"`
}

//es 配置
type ElasticConfig struct {
	Url      string `mapstructure:"url"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

//日志管理配置
type ManagerConfig struct {
}

// AppConfig 日志收集器配置
type AppConfig struct {
	Port       int             `mapstructure:"port"`
	LogLevel   string          `mapstructure:"log-level"`   //日志等级
	BufferSize int32           `mapstructure:"buffer-size"` //缓冲对内大小
	Kafka      *KafkaConfig    `mapstructure:"kafka"`       //kafka 配置
	Elastic    *ElasticConfig  `mapstructure:"elastic"`     //es配置
	Source     *SourceConfig   `mapstructure:"source"`      //日志源
	Analyzer   *AnalyzerConfig `mapstructure:"analyzer"`    //日志解析器
	Manager    *ManagerConfig  `mapstructure:"manager"`     //日志管理
}
