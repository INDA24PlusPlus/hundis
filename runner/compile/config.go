package compile

type Config struct {
	SourceName       string
	OutputName       string
	NsjailConfigName string
	ShouldCompile    bool
}

var ConfigRepository = initConfigRepository()

type configMap map[string]*Config

func initConfigRepository() configMap {
	return configMap{
		"cpp": {
			SourceName:       "submission.cpp",
			OutputName:       "submission",
			NsjailConfigName: "/configs/cpp.cfg",
			ShouldCompile:    true,
		},
		"python": {
			SourceName:       "submission.py",
			OutputName:       "submission.py",
			NsjailConfigName: "",
			ShouldCompile:    false,
		},
	}
}

func GetConfig(name string) (*Config, bool) {
	config, exists := ConfigRepository[name]
	return config, exists
}

func ListConfigs() []string {
	names := make([]string, 0, len(ConfigRepository))
	for name := range ConfigRepository {
		names = append(names, name)
	}
	return names
}
