package helpers

import "fmt"

func CheckError(err error) error {
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}