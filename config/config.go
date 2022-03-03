package config

type Server struct {
	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	// zap
	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
	//	Timer
	Timer Timer `mapstructure:"timer" json:"timer" yaml:"timer"`
	//	System
	System System `mapstructure:"system" json:"system" yaml:"system"`
	// OSS-Local
	Local Local `mapstructure:"local" json:"local" yaml:"local"`
	//	JWT
	JWT JWT `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}
