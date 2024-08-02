package tip

import (
	"fmt"
	"os"
	"reflect"
)

func Tip(c interface{}) {
	val := reflect.ValueOf(c)

	// Ensure the input is a pointer to a struct
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		panic(fmt.Errorf("expected a pointer to a struct"))
	}

	val = val.Elem()

	modeField := val.FieldByName("Mode")
	listenOnField := val.FieldByName("ListenOn")
	hostField := val.FieldByName("Host")
	portField := val.FieldByName("Port")
	adress := ""
	if hostField.IsValid() && portField.IsValid() {
		adress = hostField.String() + portField.String()
	}
	if listenOnField.IsValid() {
		adress = listenOnField.String()
	}
	fmt.Printf("-----------------------------------------------------------\n")
	fmt.Printf("           _                           admgo-user:0.0.1\n")
	fmt.Printf("          | |                          Running in %s mode\n", modeField)
	fmt.Printf("  __ _  __| |_ __ ___   __ _  ___      Adress: %s\n", adress)
	fmt.Printf(" / _` |/ _` | '_ ` _ \\ / _` |/ _ \\     \n")
	fmt.Printf("| (_| | (_| | | | | | | (_| | (_) |    \n")
	fmt.Printf(" \\__,_|\\__,_|_| |_| |_|\\__, |\\___/     PID: %d\n", os.Getpid())
	fmt.Printf("                        __/ |          \n")
	fmt.Printf("                       |___/           https://admgo.com.cn\n")
	fmt.Printf("-----------------------------------------------------------\n")
}
