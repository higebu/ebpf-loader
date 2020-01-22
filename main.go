package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/newtools/ebpf"
)

var (
	obj  = flag.String("obj", "./src/example.o", "ebpf object file path")
	path = flag.String("dir", "", "path to pin the objects")
)

func main() {
	flag.Parse()
	spec, err := ebpf.LoadCollectionSpec(*obj)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	opts := ebpf.CollectionOptions{
		Programs: ebpf.ProgramOptions{LogLevel: 1, LogSize: 10240 * 10240},
	}
	coll, err := ebpf.NewCollectionWithOptions(spec, opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if *path != "" {
		coll.Pin(*path, 0755)
	}
	defer coll.Close()
	fmt.Printf("%s\n", coll)
}
