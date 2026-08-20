package main

import (
    "embed"
    "portfolio/src"
)

//go:embed templates static
var binFS embed.FS


func main() {
    src.InitApp(binFS)

    frontend := src.LoadFrontend()
    backend := src.LoadBackend()

    src.Serve(frontend, backend)
}


