package checks

import (
	"fmt"
	"os"

	"github.com/NodeSpy/procfs"
)

func check(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func getCPUStat() {
	fs, err := procfs.NewFS("/proc")
	check(err)
	num, err := fs.CPUInfo()
	load, err := fs.LoadAvg()

	check(err)
	fmt.Println("Number of CPUs: ", len(num))
	fmt.Println("1 minute Load Average:", load.Load1)
	fmt.Println("5 minute Load Average:", load.Load5)
	fmt.Println("15 minute Load Average:", load.Load15)

}

func Start() {
	fmt.Println("getting CPU stats")
	getCPUStat()
}
