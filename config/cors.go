package config

var allowedOrigins = []string{
	"http://localhost:5173",
	"https://my-fe-three.vercel.app",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}
