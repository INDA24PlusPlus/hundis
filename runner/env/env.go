package env

type Configuration struct {
	S3Id     string `env:"S3_ID"`
	S3Secret string `env:"S3_SECRET"`
}
