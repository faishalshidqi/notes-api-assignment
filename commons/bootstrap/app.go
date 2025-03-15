package bootstrap

type Application struct {
	Env *Env
	DB  Database
}

func App() Application {
	app := Application{}
	app.Env = NewEnv()
	app.DB = *NewMySQLDatabase(app.Env)
	return app
}
