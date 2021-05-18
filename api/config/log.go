package config

// LogConfig 日志配置文件
type LogConfig struct {
	LogDirector      string
	LogInfoFilename  string
	LogInfoFilePath  string
	LogErrorFilename string
	LogErrorFilePath string
	LogMaxSize       int
	LogMaxBackups    int
	LogMaxAge        int
	LogLevel         string
}
