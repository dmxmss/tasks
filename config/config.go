package config

type (
	Config struct {
		app *App
		database *Database
	}

	App struct {
		Address string
		Port string
	}

	Database struct {
		Name string	
		User string
		Port string
		Password string
	}
)
