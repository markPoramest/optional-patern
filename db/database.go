package db

type Database struct {
	Host     string
	Driver   string
	Port     int
	UserName string
	Password string
}

type Options func(*Database)

func WithHost(host string) Options {
	return func(db *Database) {
		db.Host = host
	}
}

func WithDriver(driver string) Options {
	return func(db *Database) {
		db.Driver = driver
	}
}

func WithPort(port int) Options {
	return func(db *Database) {
		db.Port = port
	}
}

func WithUserNamePassword(userName, password string) Options {
	return func(db *Database) {
		db.UserName = userName
		db.Password = password
	}
}

func New(args ...Options) *Database {
	db := &Database{}

	for _, opt := range args {
		opt(db)
	}

	return db
}
