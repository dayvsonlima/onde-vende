package main

import (
  "website/config"
)

func main() {
  config.Routes().Run(":8080")
}
