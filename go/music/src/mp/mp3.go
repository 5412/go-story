package mp

import (
    "fmt"
    "time"
)

type MP3Player struct {
    stat int
    process int
}

func (p *MP3Player) Play(source string) {
    fmt.Println("Playing MP3 music", source)
    p.process = 0

    for p.process < 100 {
        time.Sleep(100 * time.Millisecond)
        fmt.Print(".")
        p.process += 10
    }

    fmt.Println("\nFinished playing", source)
}
