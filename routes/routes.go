package routes

import "log"

type Bar struct {
    Baz int
}

func PrintBaz() {
    baz := Bar{42}
    log.Printf("Bar struct: %v", baz)
}
