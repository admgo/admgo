package config

type Config struct {
	// Username of database
	Username string
	// Password of database
	Password string
	// Host of database
	Host string `json:",default=127.0.0.1"`
	// Port of database
	Port string `json:",default=3306"`
	// Database of database
	Database string
	// Charset of database
	Charset string `json:",default=utf8mb4"`
	// MaxOpenConns is the max open connect count
	MaxOpenConns int `json:",default=1000"`
	// MaxIdleConns is the max idle connect count
	MaxIdleConns int `json:",default=5"`
	// MaxLifetime of database, unit: ms
	MaxLifetime int `json:",default=3600"`
	// TimeOut of database, unit: ms
	TimeOut int `json:",default=50000"`
	// WriteTimeOut of database, unit: ms
	WriteTimeOut int `json:",default=50000"`
	// ReadTimeOut of database, unit: ms
	ReadTimeOut int `json:",default=50000"`
}
