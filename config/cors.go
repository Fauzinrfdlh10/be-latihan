package config

var allowedOrigins = []string{
	"http://localhost:5173",
	"https://my-fe-three.vercel.app",
	"https://be-latihan-production.up.railway.app",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}
