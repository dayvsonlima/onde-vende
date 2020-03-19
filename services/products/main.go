package main

import (
  "products/config"
)

func main() {
  config.Routes().Run(":8080")
}
