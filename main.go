package main

//go:generate go-bindata ./config

import (
	"github.com/naoina/toml"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync"
	""
)

type Config struct {
	Download []string
	Install  []string
}

func main() {
	cpus := runtime.NumCPU()
	log.Printf("cpus %v", cpus)
	runtime.GOMAXPROCS(cpus)

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		log.Fatalf("No GOPATH set.")
	}

	buf, _ := Asset("config/repos.toml")

	var repos Config
	if err := toml.Unmarshal(buf, &repos); err != nil {
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
