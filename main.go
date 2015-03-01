package main

//go:generate go-bindata ./config

import (
	"flag"
	"github.com/naoina/toml"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync"
)

type Config struct {
	Download []string
	Install  []string
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		log.Fatalf("No GOPATH set.")
	}

	flag.Parse()
	args := flag.Args()

	var buf []byte
	var err error

	if len(args) == 1 {
		f, err := os.Open(args[0])
		if err != nil {
			log.Fatalf("%v", err)
		}
		defer f.Close()
		buf, err = ioutil.ReadAll(f)
		if err != nil {
			log.Fatalf("%v", err)
		}
	} else {
		buf, err = Asset("config/repos.toml")
		if err != nil {
			log.Fatalf("%v", err)
		}
	}

	var repos Config
	if err = toml.Unmarshal(buf, &repos); err != nil {
		log.Fatalf("%v", err)
	}

	// download
	var wg sync.WaitGroup
	for _, repo := range repos.Download {
		repo := repo
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Printf("Updating %s ...", repo)
			err := exec.Command("go", "get", "-u", repo).Run()
			log.Printf("Updated %s: %v", repo, err)
		}()
	}
	wg.Wait()

	//install
	for _, repo := range repos.Install {
		path := filepath.Join(gopath, "src", repo)
		os.Chdir(path)
		err := exec.Command("go", "install").Run()
		log.Printf("Installed %s: %v", repo, err)
	}
}
