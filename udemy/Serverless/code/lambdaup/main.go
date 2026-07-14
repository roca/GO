package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

type LProject struct {
	Name   string
	Bucket string
	Role   string
	path   string
}

// NewProject ...
func NewProject(fname string) (LProject, error) {

	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return LProject{}, err
	}

	var res LProject
	err = json.Unmarshal(data, &res)
	if err != nil {
		return LProject{}, err
	}

	res.path = path.Dir(fname)
	if strings.HasPrefix(res.Role, "arn:") {
		// Happy
		return res, nil
	}

	rmp, err := RoleMap()
	if err != nil {
		return LProject{}, err
	}

	nRole, ok := rmp[res.Role]
	res.Role = nRole
	if !ok {
		return LProject{}, errors.New("Role Not found: " + res.Role)
	}

	return res, nil
}

// UploadLambda ...
func (lp LProject) UploadLambda(name string) error {
	fpath := path.Join(lp.path, name)

	os.Setenv("GOOS", "linux")
	os.Setenv("GOARCH", "amd64")

	fmt.Println("Building: " + fpath + ".go")
	_, err := run("go", "build", "-o", fpath, fpath+".go")
	if err != nil {
		return err
	}

	fmt.Println("Zipping: " + fpath + ".zip")
	_, err = run("zip", "-j", fpath+".zip", fpath)
	if err != nil {
		return err
	}

	_, err = run("zip", "-j", fpath+".zip", fpath)
	if err != nil {
		return err
	}

	lamdbaName := lp.Name + "_" + name

	upcmd := exec.Command("aws", "s3", "cp", fpath+".zip", "s3://"+lp.Bucket+"/"+lamdbaName+".zip")
	upOut, err := upcmd.StdoutPipe()
	if err != nil {
		return err
	}

	fmt.Println("Starting Upload of " + lamdbaName)
	err = upcmd.Start()
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, upOut)

	err = upcmd.Wait()
	if err != nil {
		return err
	}

	fl, err := NewFunctionList()
	if err != nil {
		return err
	}

	if fl.HasFunction(lamdbaName) {
		fmt.Println("Updating Function")
		resp, err := run("aws", "lambda", "update-function-code",
			"--function-name", lamdbaName,
			"--s3-bucket", lp.Bucket,
			"--s3-key", lamdbaName+".zip")
		if err != nil {
			return err
		}
		fmt.Println(string(resp))
		return nil
	}

	fmt.Println("Createing Function")
	resp, err := run("aws", "lambda", "create-function",
		"--function-name", lamdbaName,
		"--runtime", "go1.x",
		"--role", lp.Role,
		"--handler", name,
		"--code", "S3Bucket="+lp.Bucket+",S3Key="+lamdbaName+".zip")
	if err != nil {
		return err
	}

	fmt.Println(string(resp))

	return nil
}

func main() {

	lname := flag.String("n", "", "Name of Lambda")
	confloc := flag.String("c", "project.json", "Location of Config file")
	flag.Parse()

	proj, err := NewProject(*confloc)
	if err != nil {
		log.Fatal(err)
	}

	err = proj.UploadLambda(*lname)
	if err != nil {
		log.Fatal(err)
	}

}
