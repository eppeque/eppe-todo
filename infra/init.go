package infra

func InitData() error {
	if err := assignSecret(); err != nil {
		return err
	}

	database, err := createDatabase()

	if err != nil {
		return err
	}

	Db = database
	return nil
}

func assignSecret() error {
	secret, err := initSecret()

	if err != nil {
		return err
	}

	SecretKey = secret
	return nil
}
