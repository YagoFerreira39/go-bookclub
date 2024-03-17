package main

import (
	_ "github.com/YagoFerreira39/go-bookclub/src/externals/infrastructure/http_config"
	httpconfig "github.com/YagoFerreira39/go-bookclub/src/externals/infrastructure/http_config"
)

func main() {
	httpconfig.RegisterHttpConfig()
}
