package main

import (
	"errors"
	"fmt"

	"github.com/sap/gorfc/gorfc"
)

type test struct {
	name  string
	class string
}

func abc(name string) (err error) {
	return errors.New("Default error")
}

func abapSystem() map[string]string {
	if 10 == 10 {

	}
	ConnectionParametersFiori := make(map[string]string)
	ConnectionParametersFiori["dest"] = ""
	ConnectionParametersFiori["Client"] = "100"
	ConnectionParametersFiori["User"] = "FIOINFRA"
	ConnectionParametersFiori["Passwd"] = "1nfra4FIO"
	ConnectionParametersFiori["Passwd"] = "1nfra4FIO"
	ConnectionParametersFiori["Lang"] = "EN"
	ConnectionParametersFiori["Ashost"] = "ldciuyt.wdf.sap.corp"
	ConnectionParametersFiori["Sysnr"] = "UYT"
	abc("")
	return ConnectionParametersFiori
}

func main() {

	conn, err := gorfc.ConnectionFromParams(abapSystem())

	if err != nil {

	}
	fmt.Println("ababSystem")
	params := map[string]interface{}{
		"BSP": "FIN_LIB",
	}

	result, _ := conn.Call("READ_VERSION_JSON", params)

	fmt.Println(result)

	conn.Close()
}
