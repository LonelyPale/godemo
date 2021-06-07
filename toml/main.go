package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Config struct {
	Age        int
	Cats       []string
	Pi         float64
	Perfection []int
	DOB        time.Time
}

func main() {
	read2()
}

func write() {
	fmt.Println(string(1234567890))

	f, err := os.Create("toml/test.toml")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	config := &Config{DOB: time.Now()}
	if err := toml.NewEncoder(f).Encode(config); err != nil {
		log.Fatal(err)
	}
}

type Man struct {
	Songs []Song
}

type Song struct {
	NameAbc  string `toml:"name_abc"`
	Duration duration
}

type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

func read() {
	bs, _ := ioutil.ReadFile("/Users/wyb/project/github/godemo/toml/test-read.toml")
	var favorites Man
	if _, err := toml.Decode(string(bs), &favorites); err != nil {
		log.Fatal(err)
	}

	for _, s := range favorites.Songs {
		fmt.Printf("%s (%s)\n", s.NameAbc, s.Duration)
	}
}

type Conf struct {
	Groups []Group `toml:"groups"`
}

type Group struct {
	TmpDir   string `toml:"tmp_dir"`
	FinalDir string `toml:"final_dir"`
}

func read2() {
	bs, _ := ioutil.ReadFile("/Users/wyb/project/github/godemo/toml/test-read2.toml")
	var conf Conf
	if _, err := toml.Decode(string(bs), &conf); err != nil {
		log.Fatal(err)
	}

	for _, s := range conf.Groups {
		fmt.Printf("%s (%s)\n", s.TmpDir, s.FinalDir)
	}
}
