// Inspired by: https://github.com/lizrice/containers-from-scratch/blob/master/main.go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

func main() {
	if len(os.Args) == 1 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
		case "run":
			run()
		case "child":
			child()
		default:
			panic("Unsupported command.")
	}
}

func usage() {
	fmt.Printf("Usage: %v <cmd> <args>\n", os.Args[0])
}

func run() {
	fmt.Printf("Running %v\n", os.Args[2:])
	// TODO: Look into /proc/self/exe - Like fork and exec?
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	// TODO: Look into syscall commands below, Unix Time Sharing / New NAmespace, etc.
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	must(cmd.Run())
}

func child() {
	fmt.Printf("Running %v as PID %v\n", os.Args[2:], os.Getpid())
	//TODO: Look into cgroups
	cg()

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr

	must(syscall.Sethostname([]byte("container")))
	//TODO: Look into chroot, expects rootfs at this path
	must(syscall.Chroot("./fake-rootfs"))
	must(os.Chdir("/"))
	//TODO: Look into this, related to listing processes with `ps`
	must(syscall.Mount("proc", "proc", "proc", 0, ""))
	//must(syscall.Mount("thing", "mytmp", "tmpfs", 0, ""))
	
	must(cmd.Run())

	must(syscall.Unmount("proc", 0))
	//must(syscall.Unmount("thing", 0))
}

func cg() {
	cgroups := "/sys/fs/cgroup/"
	pids := filepath.Join(cgroups, "pids")
	os.Mkdir(filepath.Join(pids, "rmccorm4"), 0755)
	must(ioutil.WriteFile(filepath.Join(pids, "rmccorm4/pids.max"), []byte("20"), 0700))
	// Removes new cgroup after container exits
	must(ioutil.WriteFile(filepath.Join(pids, "rmccorm4/notify_on_release"), []byte("1"), 0700))
	must(ioutil.WriteFile(filepath.Join(pids, "rmccorm4/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
