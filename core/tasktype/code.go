package tasktype

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strconv"
	"time"

	"github.com/labulaka521/crocodile/core/utils/resp"
)

var _ TaskRuner = DataCode{}

// DataCode run code
type DataCode struct {
	Lang     Lang     `json:"lang"`
	LangDesc string   `json:"langdesc" comment:"Lang"`
	Code     string   `json:"code" comment:"Code"`
	Params   []string `json:"params" comment:"Code"`
}

// Lang task type lang code
type Lang uint8

const (
	shell Lang = iota + 1
	python3
	golang
	python
	nodejs
	windowsbat
)

// String return Lanf str
func (l Lang) String() string {
	switch l {
	case shell:
		return "shell"
	case python:
		return "python"
	case python3:
		return "python3"
	case golang:
		return "golang"
	case nodejs:
		return "nodejs"
	case windowsbat:
		return "windowsbat"
	default:
		return "unknow lang"
	}
}

func getcmd(ctx context.Context, lang Lang, code string) (*exec.Cmd, string, error) {
	switch lang {
	case shell:
		return runshell(ctx, code)
	case python:
		return runpython(ctx, code)
	case python3:
		return runpython3(ctx, code)
	case golang:
		return rungolang(ctx, code)
	case nodejs:
		return runnodejs(ctx, code)
	case windowsbat:
		return runwindowsbat(ctx, code)
	default:
		return nil, "", fmt.Errorf("can not support lang: %d", lang)
	}
}

// Shell
// run shell code
func runshell(ctx context.Context, code string) (*exec.Cmd, string, error) {
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/sh"
	}
	tmpfile, err := ioutil.TempFile("", "*.sh")
	if err != nil {
		return nil, "", err
	}
	shellcodepath := tmpfile.Name()
	_, err = tmpfile.WriteString(code)
	if err != nil {
		return nil, "", err
	}

	tmpfile.Sync()
	tmpfile.Close()
	cmd := exec.CommandContext(ctx, shell, shellcodepath)
	return cmd, shellcodepath, nil
}

// Python
// run python code
func runpython(ctx context.Context, code string) (*exec.Cmd, string, error) {
	tmpfile, err := ioutil.TempFile("", "*.py")
	if err != nil {
		return nil, "", err
	}
	pythoncodepath := tmpfile.Name()
	_, err = tmpfile.WriteString(code)
	if err != nil {
		return nil, "", err
	}
	tmpfile.Sync()
	tmpfile.Close()
	cmd := exec.CommandContext(ctx, "python", pythoncodepath)
	return cmd, pythoncodepath, nil
}

// Python3
// run python code
func runpython3(ctx context.Context, code string) (*exec.Cmd, string, error) {
	tmpfile, err := ioutil.TempFile("", "*.py")
	if err != nil {
		return nil, "", err
	}
	python3codepath := tmpfile.Name()
	_, err = tmpfile.WriteString(code)
	if err != nil {
		return nil, "", err
	}
	tmpfile.Sync()
	tmpfile.Close()
	cmd := exec.CommandContext(ctx, "python3", python3codepath)
	return cmd, python3codepath, nil
}

// Javascript
// run python code
func runnodejs(ctx context.Context, code string) (*exec.Cmd, string, error) {
	tmpfile, err := ioutil.TempFile("", "*.js")
	if err != nil {
		return nil, "", err
	}
	nodejscodepath := tmpfile.Name()
	_, err = tmpfile.WriteString(code)
	if err != nil {
		return nil, "", err
	}
	tmpfile.Sync()
	tmpfile.Close()
	cmd := exec.CommandContext(ctx, "node", nodejscodepath)
	return cmd, nodejscodepath, nil
}

// Golang
const (
	modcontent = `module crocodile

go `
	modname   = "go.mod"
	gonamepre = "crocodile_"
)

// run golang code
func rungolang(ctx context.Context, code string) (*exec.Cmd, string, error) {
	// golang version must rather equal 1.11
	// GO111MODULE ust be on
	cmd := exec.CommandContext(context.Background(), "go", "version")

	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, "", err
	}

	pattern := `[0-1]\.[0-9]{1,2}`
	re := regexp.MustCompile(pattern)
	goversion := re.FindString(string(out))
	if goversion < "1.11" {
		err := errors.New("go version must rather equal go1.11 and enable go module")
		return nil, "", err
	}
	if os.Getenv("GO111MODULE") != "on" {
		os.Setenv("GO111MODULE", "on")
	}
	modcontent := modcontent + goversion + "\n"

	tmpdir, err := ioutil.TempDir("", "crocodile_")
	if err != nil {
		return nil, "", err
	}

	err = ioutil.WriteFile(path.Join(tmpdir, modname), []byte(modcontent), os.ModePerm)
	if err != nil {
		return nil, "", err
	}
	gonamefile := gonamepre + strconv.FormatInt(time.Now().Unix(), 10) + ".go"
	err = ioutil.WriteFile(path.Join(tmpdir, gonamefile), []byte(code), os.ModePerm)
	if err != nil {
		return nil, "", err
	}

	os.Chdir(tmpdir)

	gocmd := exec.CommandContext(ctx, "go", "run", gonamefile)

	return gocmd, tmpdir, nil
}

// Windows bat
// run bat code
func runwindowsbat(ctx context.Context, code string) (*exec.Cmd, string, error) {
	tmpfile, err := ioutil.TempFile("", "*.bat")
	if err != nil {
		return nil, "", err
	}
	batcodepath := tmpfile.Name()
	_, err = tmpfile.WriteString(code)
	if err != nil {
		return nil, "", err
	}

	tmpfile.Sync()
	tmpfile.Close()
	cmd := exec.CommandContext(ctx, "cmd", "/C", batcodepath)
	return cmd, batcodepath, nil
}

// Type return task run lang
func (ds DataCode) Type() string {
	return ds.Lang.String()
}

// Run implment TaskRuner
// run shell command
// return io.ReadCloser
func (ds DataCode) Run(ctx context.Context) io.ReadCloser {
	pr, pw := io.Pipe()
	go func() {
		var (
			exitCode = DefaultExitCode
			err      error
			codepath string
			cmd      *exec.Cmd
		)
		defer pw.Close()
		defer func() {
			now := time.Now().Local().Format("2006-01-02 15:04:05: ")
			pw.Write([]byte(fmt.Sprintf("%sTask Run Finished,Return Code:%5d", now, exitCode))) // write exitCode,total 5 byte
			if codepath != "" {
				_ = os.Remove(codepath)
			}
		}()
		cmd, codepath, err = getcmd(ctx, ds.Lang, ds.Code)
		if err != nil {
			pw.Write([]byte(err.Error()))
			return
		}
		cmd.Stdout = pw
		cmd.Stderr = pw
		stdin, err := cmd.StdinPipe()
		err = cmd.Start()
		if err != nil {
			pw.Write([]byte(err.Error()))
			return
		}
		// 输入参数
		for _, param := range ds.Params {
			io.WriteString(stdin, param+"\n")
		}

		err = cmd.Wait()
		if err != nil {
			// deal err
			// if context err,will change err to custom msg
			switch ctx.Err() {
			case context.DeadlineExceeded:
				pw.Write([]byte(resp.GetMsg(resp.ErrCtxDeadlineExceeded)))
			case context.Canceled:
				pw.Write([]byte(resp.GetMsg(resp.ErrCtxCanceled)))
			default:
				pw.Write([]byte(err.Error()))
			}

			// try to get the exit code
			if exitError, ok := err.(*exec.ExitError); ok {
				exitCode = exitError.ExitCode()
			}
		} else {
			exitCode = 0
		}

	}()
	return pr
}
