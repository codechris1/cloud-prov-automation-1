package apivalues

import (
	// logs "loggers"
	libs "libraries"
	"encoding/json"
)

type ApiValues struct {
	ApiValues []ApiValue `json:"apiValues"`
}

type ApiValue struct {
	ParamID int `json:"paramID"`
	ParamName string `json:"paramName"`
	Values []string `json:"values"`
}

func GetApiValues() ApiValues {
	config := libs.LoadConfig()
	file := config["configsPath"] + "/" + config["apiVaulesFile"] 

	byteValue := libs.OpenFile(file)

    var result ApiValues

    json.Unmarshal(byteValue, &result)

	return result
}
