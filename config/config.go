package config

import (
    "os"
    "fmt"
)

const (
    CONFIG_FILE_NAME = "gourl.conf"
)

var (
    ConfigValues map[string]string
)

func init() {
    if file, err := os.Open(CONFIG_FILE_NAME); err != nil {
        file, _ = os.Create(CONFIG_FILE_NAME)
        fmt.Fprintf(file, "domain:test.com\n")
        fmt.Fprintf(file, "url_len:4\n")
        fmt.Fprintf(file, "port:8080\n")

        file.Close()
    }
    file, _ := os.Open(CONFIG_FILE_NAME)
    ConfigValues = make(map[string]string)
    var key, value string
    for i := 0; i < 3; i++ {
        fmt.Fscanf(file, "%s:%s", &key, &value)
        ConfigValues[key] = value
        //fmt.Println(key, value)
    }
    file.Close()

}
