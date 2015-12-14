package execute

import (
	"bufio"
	"io"
	"os/exec"
)

// Execute run a command and return the stdout and stderr as string array.
func Execute(name string, args ...string) ([]string, []string, error) {
	cmd := exec.Command(name, args...)
	// create the pipes for stderr and stdout...
	stderr, errStderr := cmd.StderrPipe()
	if errStderr != nil {
		return []string{}, []string{}, errStderr
	}
	stdout, errStdout := cmd.StdoutPipe()
	if errStdout != nil {
		return []string{}, []string{}, errStdout
	}
	// execute the command
	cmd.Start()
	// read the stderr pipe
	stderrText := readerToStringArray(stderr)
	stdoutText := readerToStringArray(stdout)
	return stdoutText, stderrText, nil
}

func readerToStringArray(in io.Reader) []string {
	var out []string
	r := bufio.NewReader(in)
	line, _, err := r.ReadLine()
	for err == nil {
		out = append(out, string(line)) //string(line) + "\n"
		// fmt.Println(string(line))
		line, _, err = r.ReadLine()
	}
	return out
}
