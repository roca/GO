package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"syscall"
	"time"
)

var commands []string

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all the machine's cores
	fmt.Println("The number of CPUs on this server is", runtime.NumCPU())

	flag.String("h", "help", "File Watcher")
	var run = flag.Bool("run", false, "run flag")
	var ruby_env = flag.String("ruby_env", "production", "ruby_env flag")

	flag.Parse()
	if len(os.Args) == 1 {
		fmt.Printf(usage(filepath.Base(os.Args[0])))
		os.Exit(1)
	}
	root_path := os.Args[len(os.Args)-1]
	_, err := os.Stat(root_path)
	if err != nil {
		fmt.Println(MyError{fmt.Sprintf("NonExisting file path : %s", root_path)})
		fmt.Printf(usage(filepath.Base(os.Args[0])))
		os.Exit(1)
	}
	fmt.Println("Folders found\n")
	filepath.Walk(root_path, file_watcher)
	//done := make(chan bool, len(commands))
	fmt.Println("-------------------------------------------\n")

	ruby_environment := fmt.Sprintf("ruby_env=%s", *ruby_env)
	for _, command := range commands {
		//go func() {
		//command_with_taskset := fmt.Sprintf("taskset -c %d %s", i, command)
		fmt.Println("executing: ", command)
		if *run {
			out := fmt.Sprintf("%s\n", exec_command(command, ruby_environment))
			fmt.Println(out)
		}
		//done <- true
		//}()
	}
	//for i := 0; i < len(commands); i++ {
	//	<-done
	//}

	fmt.Println("Done !")

}

func statTimes(name string) (atime, mtime, ctime time.Time, err error) {
	fi, err := os.Stat(name)
	if err != nil {
		return
	}
	mtime = fi.ModTime()
	stat := fi.Sys().(*syscall.Stat_t)
	atime = time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec))
	atime = time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec))
	ctime = time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec))
	return
}

func usage(arg string) string {
	return fmt.Sprintf("usage: %s <some_directory_base_path>\n", arg)
}

func file_watcher(path string, info os.FileInfo, err error) error {
	sub_folders := strings.Split(path, "/")
	last_sub_folder := sub_folders[len(sub_folders)-1]
	matched, err := regexp.MatchString("\\d{6}\\_*", last_sub_folder)

	if info.IsDir() && matched {
		atime, mtime, ctime, err := statTimes(fmt.Sprintf("%s/CompletedJobInfo.xml", path))

		if err != nil {
			fmt.Println(err)
			return filepath.SkipDir
		}

		t0 := time.Now()

		mhours_duration := t0.Sub(mtime).Hours()
		ahours_duration := t0.Sub(atime).Hours()
		chours_duration := t0.Sub(ctime).Hours()

		if mhours_duration >= .25 && chours_duration <= 2.0 {
			fmt.Printf("%s\tfile age in hours: a_age:%v  m_age:%v c_age:%v  \n", path, ahours_duration, mhours_duration, chours_duration)
			flowcell_command := fmt.Sprintf("/cm/shared/apps/blackjack/bin/flowcell run=%s", path)
			fmt.Printf("%s", exec_command(fmt.Sprintf("ls -l %s/CompletedJobInfo.xml", path)))
			//fmt.Println(flowcell_command, "\n")
			commands = append(commands, flowcell_command)
		}
		return filepath.SkipDir
	}
	return nil
}

func exec_command(command string, optionalParams ...string) []byte {

	command_arguments := strings.Split(command, " ")

	for _, options := range optionalParams {
		command_arguments = append(command_arguments, options)
	}

	cmd := exec.Command(command_arguments[0], command_arguments[1:]...)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(command_arguments)
		log.Fatal("ERROR: ", err)
		log.Fatal(fmt.Sprintf("%s\n", out))
	}

	return out

}

// MyError is an error implementation that includes a time and message.
type MyError struct {
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v", e.What)
}
