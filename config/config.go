package config

import (
    "os"
    "fmt"
)

const (
    CONFIG_FILE_NAME = "gourl.conf"
)

var (
    configuration map[string]string
)

func init() {
    check()
    read()
}

func check() error {
    var (
        file *os.File
        err error
    )
    if file, err = os.Open(CONFIG_FILE_NAME); err != nil {
        file, err = os.Create(CONFIG_FILE_NAME)
        fmt.Fprintf(file, "domain:test.com\n")
        fmt.Fprintf(file, "url_len:4\n")
        fmt.Fprintf(file, "port:8080\n")
        file.Close()
    }

    return err
}

func read() {
    configuration = make(map[string]string)
    var option, value string

    file, _ := os.Open(CONFIG_FILE_NAME)
    defer file.Close()

    for {
        if n, err := fmt.Fscanf(file, "%s:%s", &option, &value); n==0 || err!=nil {
            break
        }
        configuration[option] = value
    }
}

func Value(option string) (string, bool) {
    value, ok := configuration[option]
    return value, ok
}
