package main

//go:generate go-bindata ./config

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/naoina/toml"
)

type config struct {
	Download []string
	Install  []string
}

func main() {
	flag.Parse()

	var conf config
	if err := toml.Unmarshal(reposData(), &conf); err != nil {
		log.Fatalf("%v", err)
	}
	download(conf.Download)
	install(conf.Install)
}

func reposData() (buf []byte) {
	var err error
	if flag.NArg() == 1 {
		buf, err = ioutil.ReadFile(flag.Arg(0))
	} else {
		buf, err = Asset("config/repos.toml")
	}
	if err != nil {
		log.Fatalf("%v", err)
	}
	return buf
}

func download(repos []string) {
	var wg sync.WaitGroup
	for _, repo := range repos {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			log.Printf("Updating %s ...", url)
			err := exec.Command("go", "get", "-u", url).Run()
			log.Printf("Updated %s: %v", url, err)
		}(repo)
	}
	wg.Wait()
}

func install(repos []string) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		log.Fatalf("No GOPATH set.")
	}
	for _, repo := range repos {
		path := filepath.Join(gopath, "src", repo)
		if err := os.Chdir(path); err != nil {
			log.Fatalf("%v", err)
		}
		err := exec.Command("go", "install").Run()
		log.Printf("Installed %s: %v", repo, err)
	}
}
