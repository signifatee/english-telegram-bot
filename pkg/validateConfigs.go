package pkg

import "os"

func ValidateConfigs() (err bool, msg string) {

	_, err = os.LookupEnv("TELEGRAM_API_TOKEN")
	if err == false {
		return false, "No TELEGRAM_API_TOKEN specified"
	}

	return err, ""

}
