package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	"flag"
	"regexp"
	"strings"
)

func main() {
    flag.String("h","help","File Watcher")
	flag.Parse()
	if len(os.Args) == 1 {
		fmt.Printf(usage(filepath.Base(os.Args[0])))
		os.Exit(1)
	}
	root_path := os.Args[1]
	_,err := os.Stat(root_path)
	if err != nil {
		fmt.Println(MyError{fmt.Sprintf("NonExisting file path : %s",root_path)})
		fmt.Printf(usage(filepath.Base(os.Args[0])))
		os.Exit(1)
	}
	filepath.Walk(root_path, file_watcher)
}

func usage(arg string) string {
	return fmt.Sprintf("usage: %s <some_directory_base_path>\n",arg)
}



func file_watcher(path string, info os.FileInfo, err error) error {
	  sub_folders := strings.Split(path,"/")
	  last_sub_folder := sub_folders[len(sub_folders)-1]
	  matched, err := regexp.MatchString("\\d{6}\\_*", last_sub_folder)
      if info.IsDir() && matched {
	    _,rta_err := os.Stat(fmt.Sprintf("%s/RTAComplete.txt",path))
	  	mod_time := info.ModTime()
      	t0 := time.Now()
      	hours_duration := t0.Sub(mod_time).Hours()
//      	if rta_err == nil && hours_duration <= 1.0 {
      	if rta_err == nil {
	  		fmt.Printf("%50s\tduration in hours: %v\n",path, hours_duration)
      	}
	  	return filepath.SkipDir
	  }
	  return nil
}

// MyError is an error implementation that includes a time and message.
type MyError struct {
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v", e.What)
}