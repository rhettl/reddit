package main

import (
    "os"
    "fmt"
    "log"
    "errors"
    "github.com/Tgeahre/reddit"
)

func osArg(i int) (string, error) {
    if len(os.Args) < 2 {
        return "", errors.New("You must enter the name of the reddit channel you want to read.")
    }
    if len(os.Args) < (1 + i) {
        return "", errors.New("The requested command line argument is not pressent.")
    }

    return os.Args[i], nil
}

func main() {
    forum, err := osArg(1)
    if err != nil {
        log.Fatal(err)
    }

    items, err := reddit.Get(forum)
    if err != nil {
        log.Fatal(err)
    }
    for _, item := range items {
        fmt.Println(item)
    }
}
