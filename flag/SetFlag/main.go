package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
)

//FlagSet
type URLValue struct {
	URL *url.URL
}

func (v URLValue) String() string {
	if v.URL != nil {
		return v.URL.String()
	}
	return ""
}
func (v URLValue) Set(s string) error {
	if u, err := url.Parse(s); err != nil {
		return err
	} else {
		*v.URL = *u
	}
	return nil
}

var u = &url.URL{}

func main() {
	fs1 := flag.NewFlagSet("one", flag.ExitOnError)
	fs1.Var(&URLValue{u}, "ourl", "URL to parse")

	//var fs2Value string
	fs2 := flag.NewFlagSet("two", flag.ExitOnError)
	fs2Value := fs2.String("turl", "http://www.qq.com", "URL to parse")

	var err error
	switch os.Args[1] {
	case "one":
		err = fs1.Parse(os.Args[2:])
	case "two":
		err = fs2.Parse(os.Args[2:])
	default:
		fmt.Println("default!")
		os.Exit(1)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf(`{scheme: %q, host: %q, path: %q}`, u.Scheme, u.Host, u.Path)
	fmt.Printf(`fs2.String %s`, *fs2Value)
}
