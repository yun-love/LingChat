package config

import (
	"os"
	"strconv"
)

// Config 应用程序的总体配置
type Config struct {
	Server   Server         `json:"server" yaml:"server"`
	Data     Data           `json:"data" yaml:"data"`
	Chat     ChatConfig     `json:"chat" yaml:"chat"`
	Backend  BackendConfig  `json:"backend" yaml:"backend"`
	Vits     VitsConfig     `json:"vits" yaml:"vits"`
	Emotion  EmotionConfig  `json:"emotion" yaml:"emotion"`
	TempDirs TempDirsConfig `json:"temp_dirs" yaml:"temp_dirs"`
}

type Server struct {
	JWTSecret string `json:"jwt_secret" yaml:"jwt_secret"`
}

type Data struct {
	DataBase DataBase `json:"database" yaml:"database"`
}

type DataBase struct {
	Driver      string `json:"driver" yaml:"driver"`
	Source      string `json:"source" yaml:"source"`
	AutoMigrate bool   `json:"auto_migrate" yaml:"auto_migrate"`
}

// ChatConfig 聊天API配置
type ChatConfig struct {
	APIKey  string `json:"api_key,omitempty" yaml:"api_key,omitempty"`
	BaseURL string `json:"base_url" yaml:"base_url"`
	Model   string `json:"model" yaml:"model"`
}

// BackendConfig 后端服务配置
type BackendConfig struct {
	LogDir   string `json:"log_dir" yaml:"log_dir"`
	BindAddr string `json:"bind_addr" yaml:"bind_addr"`
	Addr     string `json:"addr" yaml:"addr"`
	Port     int    `json:"port" yaml:"port"`
}

// VitsConfig 语音合成配置
type VitsConfig struct {
	APIURL    string `json:"api_url" yaml:"api_url"`
	SpeakerID int    `json:"speaker_id" yaml:"speaker_id"`
}

// EmotionConfig 情感分类配置
type EmotionConfig struct {
	URL string `json:"url" yaml:"url"`
}

// TempDirsConfig 临时目录配置
type TempDirsConfig struct {
	VoiceDir string `json:"voice_dir" yaml:"voice_dir"`
}

func GetConfigFromEnv() *Config {
	// 从环境变量读取整数值
	vitsSpkID, _ := strconv.Atoi(os.Getenv("VITS_SPEAKER_ID"))
	backendPort, _ := strconv.Atoi(os.Getenv("BACKEND_PORT"))

	autoMigrate, _ := strconv.ParseBool(os.Getenv("AUTO_MIGRATE"))

	// 创建并返回配置结构体
	return &Config{
		Server: Server{
			JWTSecret: os.Getenv("JWT_SECRET"),
		},
		Data: Data{
			DataBase{
				Driver:      os.Getenv("DATABASE_DRIVER"),
				Source:      os.Getenv("DATABASE_SOURCE"),
				AutoMigrate: autoMigrate,
			},
		},
		Chat: ChatConfig{
			APIKey:  os.Getenv("CHAT_API_KEY"),
			BaseURL: os.Getenv("CHAT_BASE_URL"),
			Model:   os.Getenv("MODEL_TYPE"),
		},
		Backend: BackendConfig{
			LogDir:   os.Getenv("BACKEND_LOG_DIR"),
			BindAddr: os.Getenv("BACKEND_BIND_ADDR"),
			Addr:     os.Getenv("BACKEND_ADDR"),
			Port:     backendPort,
		},
		Vits: VitsConfig{
			APIURL:    os.Getenv("VITS_API_URL"),
			SpeakerID: vitsSpkID,
		},
		Emotion: EmotionConfig{
			URL: os.Getenv("EMOTION_PREDICT_URL"),
		},
		TempDirs: TempDirsConfig{
			VoiceDir: os.Getenv("TEMP_VOICE_DIR"),
		},
	}
}
