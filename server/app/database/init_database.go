package database

func InitDatabase() error {
	err := InitMySQL()
	if err != nil {
		return err
	}

	InitRedis()
	return nil
}
