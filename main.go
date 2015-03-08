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
	flag.Parse()

	var conf Config
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
}

func install(repos []string) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		log.Fatalf("No GOPATH set.")
	}

	for _, repo := range repos {
		path := filepath.Join(gopath, "src", repo)
		os.Chdir(path)
		err := exec.Command("go", "install").Run()
		log.Printf("Installed %s: %v", repo, err)
	}
}
