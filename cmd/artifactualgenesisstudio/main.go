// cmd/artifactualgenesisstudio/main.go
package main

import (
"flag"
"log"
"os"

"artifactualgenesisstudio/internal/artifactualgenesisstudio"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := artifactualgenesisstudio.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
