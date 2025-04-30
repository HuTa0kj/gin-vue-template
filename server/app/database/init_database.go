package database

func InitDatabase() error {
	err := InitMySQL()
	if err != nil {
		return err
	}
	
	err = InitRedis()
	if err != nil {
		return err
	}
	return nil
}
