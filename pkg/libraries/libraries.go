package libraries

import (
    "encoding/json"
    "io/ioutil"
    "os"
	logs "loggers"
)

func OpenFile(file string) []byte {
    openFile, err := os.Open(file)

    if err != nil {
        logs.Error(err)
    }

	logs.Info("Successfully loaded file " + file)

    defer openFile.Close()
    byteValue, _ := ioutil.ReadAll(openFile)

    return byteValue
}

func LoadConfig() map[string]string {
    byteValue := OpenFile("configs/configuration.json")

    var result map[string]string
    json.Unmarshal([]byte(byteValue), &result)

	return result
}
