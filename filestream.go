package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

var (
	version       = "0.0.0"
	build         = "0"
	versionString = func() string {
		return fmt.Sprintf("%s@%s", version, build)
	}
)

// Args is
type Args struct {
	help    bool
	version bool
	command string
	files   []string
}

var (
	args Args
)

func init() {

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s %s:\n", os.Args[0], versionString())
		fmt.Fprintf(flag.CommandLine.Output(), "[OPTION ...] COMMAND FILE ...\n")
		flag.PrintDefaults()
	}

	flag.BoolVar(&args.version, "version", false, "version")
	flag.BoolVar(&args.help, "h", false, "help")

	flag.Parse() // flag.Parse

	if args.help {
		flag.Usage()
		os.Exit(1)
	}
	if args.version {
		flag.Usage()
		os.Exit(1)
	}

	if 2 > len(flag.Args()) {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {

	for idx, arg := range flag.Args() {
		switch idx {
		case 0:
			args.command = arg
		default:
			args.files = append(args.files, arg)
		}
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(args.command)
	stdin, err := cmd.StdinPipe()
	if nil != err {
		log.Fatal(err)
	}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	cmd.Start()
	if nil != err {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		for _, file := range args.files {
			data, err := ioutil.ReadFile(file)
			if nil != err {
				log.Fatal(err)
			}
			s := string(data)
			_, err = io.WriteString(stdin, s)
			// _, err = stdin.Write(data)
			if nil != err {
				log.Fatal(err)
			}
		}
	}()

	cmd.Wait()

	// out, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("%s\n", out)

	fmt.Fprintf(os.Stdout, "%s\n", stdout.String())
	fmt.Fprintf(os.Stderr, "%s\n", stderr.String())
}

// CopyTo is
func CopyTo(src io.ReadCloser, dest io.WriteCloser) (int, error) {

	var count int
	buff := make([]byte, 1<<13)

	for {
		n, err := src.Read(buff)
		if nil != err {
			return count, err
		}
		if 0 == n {
			return count, err
		}
		dest.Write(buff[:n])
		count += n
	}
}
