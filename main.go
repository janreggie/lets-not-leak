package main

import (
	"fmt"
	"log"

	"github.com/jessevdk/go-flags"
	"gopkg.in/gographics/imagick.v3/imagick"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	var opts Options
	_, err := flags.Parse(&opts)

	if err != nil {
		log.Fatalln("cannot parse args", err)
	}

	fmt.Println("Saving output to", opts.Output)
	log.Fatalln("Unimplemented.")
}
