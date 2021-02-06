package main

import (
    "fmt"
    "github.com/ozgio/strutil"
    "os"
    "time"
)

func main () {

    value := ""
    if len(os.Args) > 1 {
        value = os.Args[1]
    } else {
        value = os.Getenv("HELLO_ARG")
    }
    if len(value) == 0 {
        value = "-h"
    }

    if value == "-h" || value == "--help" {
        fmt.Println("USAGE: Provide a string and we'll give you a personal hello!")
        value = "default"
    }
    
    output := "Hello, " + value + "!"
    output, _ = strutil.DrawCustomBox(output, 40, strutil.Center, strutil.SimpleBox9Slice(), "\n")
    
    for {
        fmt.Println(output)
        time.Sleep(90 * time.Second)
    }

}

//For builds with CNB/Paketo:
//    Using base run-image, two options (command overrides env var):
//        docker run -e HELLO_ARG=world hello:base "hello sunshine"
//    Using tiny run image, use:
//        docker run -e HELLO_BOX_ARG=world hello:tiny

