package main

import (
	"fmt"
	"os"
	"os/exec"
	"bufio"
	"strings"
	"github.com/atotto/clipboard"
)

func run(args []string) int{

	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "[ERROR] too few arguments\n")
		return 1
	}else if len(args) > 1 {
		fmt.Fprintf(os.Stderr, "[ERROR] too many arguments\n")
		return 1
	}

	var filename string = args[0];
	exec.Command("ebb", filename).Run()
	var bb_filename string = filename[:len(filename)-4] + ".bb"


	file, err := os.Open(bb_filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] cant open file\n")
		return 1
	}
	defer file.Close()

	var bb string
	// 一行ずつ読み出し
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Index(line, "BoundingBox") != -1{
			var space_index int = strings.Index(line, " ") + 1
			var max int = len(line)
			bb = line[space_index:max]
			clipboard.WriteAll(bb)

		}
	}


	exec.Command("rm","-f", bb_filename).Run()
	return 0
}

func main() {

	os.Exit(run(os.Args[1:]));

}
