package printer

import "fmt"

func GetRed(msg interface{}) string {
	return fmt.Sprintf("\u001B[0;31m%v\u001B[0m", msg)
}

func GetGreen(msg interface{}) string {
	return fmt.Sprintf("\u001B[0;32m%v\u001B[0m", msg)
}

func GetYellow(msg interface{}) string {
	return fmt.Sprintf("\u001B[0;33m%v\u001B[0m", msg)
}

func GetBlue(msg interface{}) string {
	return fmt.Sprintf("\u001B[0;34m%v\u001B[0m", msg)
}

func GetMagenta(msg interface{}) string {
	return fmt.Sprintf("\u001B[0;35m%v\u001B[0m", msg)
}

func GetCyan(msg interface{}) string {
	return fmt.Sprintf("\u001B[0;36m%v\u001B[0m", msg)
}

func GetGrey(msg interface{}) string {
	return fmt.Sprintf("\u001B[0;31m%v\u001B[0m", msg)
}
