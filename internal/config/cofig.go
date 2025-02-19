package config

import (
	"errors"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
	"remnawawe-json/internal/utils"
)

type Config struct {
	V2rayTemplatePath    string
	V2RayTemplate        map[string]interface{}
	V2rayMuxEnabled      bool
	V2rayMuxTemplatePath string
	V2RayMuxTemplate     map[string]interface{}
	RemnaweweURL         string
	AppPort              string
}

var conf Config

func GetConfig() Config {
	return conf
}

func InitConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Warn("Env file not found")
	}

	conf.V2rayTemplatePath = os.Getenv("V2RAY_TEMPLATE_PATH")
	if conf.V2rayTemplatePath == "" {
		conf.V2rayTemplatePath = "/app/templates/v2ray/default.json"
	}
	if _, err := os.Stat(conf.V2rayTemplatePath); os.IsNotExist(err) {
		slog.Error("File does not exist: %s", conf.V2rayTemplatePath)
		panic(err)
	}
	data, err := os.ReadFile(conf.V2rayTemplatePath)
	if err != nil {
		slog.Error("Error reading file:")
		panic(err)
	}
	conf.V2RayTemplate = utils.ConvertJsonStringIntoMap(string(data))

	conf.V2rayMuxEnabled = os.Getenv("V2RAY_MUX_ENABLED") == "true"

	if conf.V2rayMuxEnabled {
		conf.V2rayMuxTemplatePath = os.Getenv("V2RAY_MUX_TEMPLATE_PATH")
		if conf.V2rayMuxTemplatePath == "" {
			conf.V2rayMuxTemplatePath = "/app/templates/mux/default.json"
		}
		if _, err := os.Stat(conf.V2rayMuxTemplatePath); os.IsNotExist(err) {
			slog.Error("Mux template file does not exist: %s", conf.V2rayMuxTemplatePath)
			panic(err)
		}
		muxData, err := os.ReadFile(conf.V2rayMuxTemplatePath)
		if err != nil {
			slog.Error("Error reading V2ray Mux template file")
			panic(err)
		}
		conf.V2RayMuxTemplate = utils.ConvertJsonStringIntoMap(string(muxData))

	}

	conf.RemnaweweURL = os.Getenv("REMNAWAWE_URL")
	if conf.RemnaweweURL == "" {
		slog.Error("remnawawe url not found")
		panic(errors.New("remnawawe url not found"))
	}
	conf.AppPort = os.Getenv("APP_PORT")
	if conf.AppPort == "" {
		slog.Error("app port not found")
		panic(errors.New("app port not found"))
	}
}
