package pkg

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"sync"
)

// ExecCommandWithOutput run a command than returns the output
func ExecCommandWithOutput(name, dir string, args ...string) (output string, err error) {
	command := exec.Command(name, args...)
	if dir != "" {
		command.Dir = dir
	}

	var stdoutIn io.ReadCloser
	if stdoutIn, err = command.StdoutPipe(); err != nil {
		return
	}
	if err = command.Start(); err != nil {
		return
	}

	var data []byte
	if data, err = ioutil.ReadAll(stdoutIn); err == nil {
		output = strings.TrimSpace(string(data))
	}
	return
}

// ExecCommandInDir run a command in a directory
func ExecCommandInDir(name, dir string, args ...string) (err error) {
	command := exec.Command(name, args...)
	if dir != "" {
		command.Dir = dir
	}

	//var stdout []byte
	//var errStdout error
	stdoutIn, _ := command.StdoutPipe()
	stderrIn, _ := command.StderrPipe()
	err = command.Start()
	if err != nil {
		return err
	}

	// cmd.Wait() should be called only after we finish reading
	// from stdoutIn and stderrIn.
	// wg ensures that we finish
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		_, _ = copyAndCapture(os.Stdout, stdoutIn)
		wg.Done()
	}()

	_, _ = copyAndCapture(os.Stderr, stderrIn)

	wg.Wait()

	err = command.Wait()
	return
}

func execCommand(name string, arg ...string) (err error) {
	return ExecCommandInDir(name, "", arg...)
}

func copyAndCapture(w io.Writer, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			_, err := w.Write(d)
			if err != nil {
				return out, err
			}
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
}
