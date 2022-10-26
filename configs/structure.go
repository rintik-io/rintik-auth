package configs

// ConfigStruct : struct config
type ConfigStruct struct {
	Logger    LumberjackLoggerConfig `mapstructure:"logger"`
	Databases DatabasesConfig        `mapstructure:"databases"`
	Services  ServicesConfig         `mapstructure:"services"`
	Etc       map[string]interface{} `mapstructure:"etc"`
}

// LumberjackLoggerConfig :
type LumberjackLoggerConfig struct {
	DailyRotate   bool `mapstructure:"dailyRotate"`
	CompressLog   bool `mapstructure:"compressLog"`
	LogToTerminal bool `mapstructure:"logToTerminal"`
}

// DatabasesConfig :
type DatabasesConfig struct {
	Redis   RedisConfig   `mapstructure:"redis"`
	Mongo   MongoConfig   `mapstructure:"mongo"`
	Postgre PostgreConfig `mapstructure:"postgre"`
	Oracle  OracleConfig  `mapstructure:"oracle"`
	Sqlite  SqliteConfig  `mapstructure:"sqlite"`
}

// ServicesConfig :
type ServicesConfig struct {
	CronJob    CronJobConfig    `mapstructure:"cronjob"`
	Daemon     DaemonConfig     `mapstructure:"daemon"`
	Graphql    GraphqlConfig    `mapstructure:"graphql"`
	Restapi    RestapiConfig    `mapstructure:"restapi"`
	Kafka      KafkaConfig      `mapstructure:"kafka"`
	Prometheus PrometheusConfig `mapstructure:"prometheus"`
}

// RedisConfig :
type RedisConfig struct {
	Host      string `mapstructure:"host"`
	Auth      string `mapstructure:"auth"`
	DB        int    `mapstructure:"db"`
	MaxIdle   int    `mapstructure:"max_idle"`
	MaxActive int    `mapstructure:"max_active"`
	Status    bool   `mapstructure:"status"`
}

type SqliteConfig struct {
	Path   string `mapstructure:"path"`
	File   string `mapstructure:"file"`
	Status bool   `mapstructure:"status"`
}

// MongoConfig : mongo config
type MongoConfig struct {
	Host    map[int]string `mapstructure:"host"`
	User    string         `mapstructure:"user"`
	Pass    string         `mapstructure:"pass"`
	DB      string         `mapstructure:"db"`
	Srv     bool           `mapstructure:"srv"`
	Cluster bool           `mapstructure:"cluster"`
	RsName  string         `mapstructure:"rs_name"`
	Status  bool           `mapstructure:"status"`
}

// PostgreConfig :
type PostgreConfig struct {
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
	User   string `mapstructure:"user"`
	Pass   string `mapstructure:"pass"`
	DB     string `mapstructure:"db"`
	Schema string `mapstructure:"schema"`
	Status bool   `mapstructure:"status"`
}

// OracleConfig :
type OracleConfig struct {
	Host    map[int]string `mapstructure:"host"`
	User    string         `mapstructure:"user"`
	Pass    string         `mapstructure:"pass"`
	DB      string         `mapstructure:"db"`
	Timeout string         `mapstructure:"timeout"`
	Status  bool           `mapstructure:"status"`
}

// CronJobConfig :
type CronJobConfig struct {
	Status bool                         `mapstructure:"status"`
	Jobs   map[string]CronJobJobsConfig `mapstructure:"jobs"`
}

// CronJobJobsConfig :
type CronJobJobsConfig struct {
	Status bool     `mapstructure:"status"`
	Every  string   `mapstructure:"every"`
	Hours  []string `mapstructure:"hours"`
}

// DaemonConfig :
type DaemonConfig struct {
	Sleep     int  `mapstructure:"sleep"`
	WaitGroup bool `mapstructure:"waitGroup"`
	Status    bool `mapstructure:"status"`
}

// Graphql : GraphQL Config
type GraphqlConfig struct {
	Port   string `mapstructure:"port"`
	Status bool   `mapstructure:"status"`
}

// RestapiConfig : REST API Config
type RestapiConfig struct {
	Port     string        `mapstructure:"port"`
	BasePath string        `mapstructure:"base_path"`
	Swagger  SwaggerConfig `mapstructure:"swagger"`
	Status   bool          `mapstructure:"status"`
}

// SwaggerConfig : REST API Swagger Config
type SwaggerConfig struct {
	Title       string   `mapstructure:"title"`
	Description string   `mapstructure:"description"`
	Schemes     []string `mapstructure:"schemes"`
}

// KafkaConfig : kafka config
type KafkaConfig struct {
	Brokers     map[int]string       `mapstructure:"brokers"`
	Assignor    string               `mapstructure:"assignor"`
	Version     string               `mapstructure:"version"`
	Verbose     bool                 `mapstructure:"verbose"`
	DialTimeout int                  `mapstructure:"dialTimeout"`
	Consumer    KafkaConsumerConfig  `mapstructure:"consumer"`
	Publisher   KafkaPublisherConfig `mapstructure:"publisher"`
}

// KafkaConsumerConfig :
type KafkaConsumerConfig struct {
	Type   string `mapstructure:"type"`
	Topic  string `mapstructure:"topic"`
	Group  string `mapstructure:"group"`
	Oldest bool   `mapstructure:"oldest"`
	Status bool   `mapstructure:"status"`
}

// KafkaPublisherConfig :
type KafkaPublisherConfig struct {
	RetryMax   int  `mapstructure:"retrymax"`
	Timeout    int  `mapstructure:"timeout"`
	Idempotent bool `mapstructure:"idempotent"`
	Status     bool `mapstructure:"status"`
}

// PrometheusConfig :
type PrometheusConfig struct {
	Status bool `mapstructure:"status"`
}
