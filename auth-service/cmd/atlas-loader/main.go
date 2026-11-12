package main

import (
	"fmt"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
	"github.com/Yeprem25/go-auth-microservice/api-gateway/internal/model"
)

func main() {
	stms, err := gormschema.New("postgres").Load(model.All()...)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Faild to load gorm chema %v", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stms)
}
