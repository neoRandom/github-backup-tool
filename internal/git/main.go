package git

import (
	"fmt"
	"os"
	"os/exec"
)

func MirrorClone(url, path string) error {
	fmt.Printf(
		"\nRunning: 'git clone --mirror %v %v'...\n",
		url, path,
	)

	cmd := exec.Command(
		"git",
		"clone",
		"--mirror",
		url,
		path,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func Clone(url, path string) error {
	fmt.Printf(
		"\nRunning: 'git clone %v %v'\n",
		url, path,
	)

	cmd := exec.Command(
		"git",
		"clone",
		url,
		path,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
