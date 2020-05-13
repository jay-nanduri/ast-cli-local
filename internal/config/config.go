package config

type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Network struct {
	EntrypointPort    string `yaml:"entrypointPort"`
	EntrypointTLSPort string `yaml:"entrypointTLSPort"`
	PrivateKeyFile    string `yaml:"privateKeyFile"`
	CertificateFile   string `yaml:"certificateFile"`
}

type ObjectStore struct {
	AccessKeyID     string `yaml:"accessKeyId"`
	SecretAccessKey string `yaml:"secretAccessKey"`
}

type MessageQueue struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type AccessControl struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Log struct {
	Level    string      `yaml:"level"`
	Rotation LogRotation `yaml:"rotation"`
}

type LogRotation struct {
	MaxSizeMB  string `yaml:"maxSizeMB"`
	MaxAgeDays string `yaml:"maxAgeDays"`
}

type AIOConfiguration struct {
	Database      Database      `yaml:"database"`
	Network       Network       `yaml:"network"`
	ObjectStore   ObjectStore   `yaml:"objectStore"`
	MessageQueue  MessageQueue  `yaml:"messageQueue"`
	AccessControl AccessControl `yaml:"accessControl"`
	Log           Log           `yaml:"log"`
}