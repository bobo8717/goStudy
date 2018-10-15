package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"
)

/*
	4种方式
*/

type interval []time.Duration

func (i *interval) String() string {
	return fmt.Sprint(*i)
}

//-deltaT 10s,15s
func (i *interval) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("interval flag already set")
	}
	for _, dt := range strings.Split(value, ",") {
		duration, err := time.ParseDuration(dt)
		if err != nil {
			return err
		}
		*i = append(*i, duration)
	}
	return nil
}

var intervalFlag interval

func init() {
	flag.Var(&intervalFlag, "deltaT", "comma-separated list of intervals to use between events")
}

//go run main.go -int 222
//go run main.go -int=222
//go run main.go --int 222
//go run main.go --int=222
func main() {
	var ip = flag.Int("int", 1234, "help message for flagname")

	var flagvar int
	flag.IntVar(&flagvar, "intvar", 1234, "help message for flagname")

	//StringVar  共享一个变量完成简写的配置
	var gopherType string
	const (
		defaultGopher = "pocket"
		usage         = "the variety of gopher"
	)
	flag.StringVar(&gopherType, "gopher_type", defaultGopher, usage)
	flag.StringVar(&gopherType, "g", defaultGopher, usage+" (shorthand)")

	flag.Parse()

	fmt.Print("Int:", *ip)
	fmt.Print("IntVar:", flagvar)
	fmt.Print("StringVar:", gopherType)
	fmt.Print("Var:", intervalFlag)

}
