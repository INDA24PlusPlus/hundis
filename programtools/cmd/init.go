/*
Copyright © 2025 Arvid Kristoffersson arvid.kristoffersson@icloud.com
*/

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// Folder and file structures for problem types, as well as default content

// Standard
var standardFolders = []string{
	"data",
	"grader",
	"scorer",
	"statement",
	"data/sample",
	"data/secret",
	"solutions",
	"solutions/100",
	"data/secret/group1",
}

var standardFiles = []string{
	"config.yaml",
	"metadata.yaml",
	"/grader/grading.go",
	"/scorer/scoring.go",
	"/statement/problem.en.tex",
	"/data/secret/group1/1.in",
	"/data/secret/group1/1.ans",
}

var standardContent = map[string]string{
	"config.yaml": `limits:
  time_limit: 4
  character_limit: 100000
  memory_limit: 1073741824

scoring:
  group_grades: [100]
  on_reject: stop
  include_sample: no

type: standard`,
	"metadata.yaml": `author: V.A. Cant
author_id: 0
admin_id: 0
source: duck competition 2025
rights_owner: duck competition
license: cc by-sa`,
	"/grader/grading.go": `package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	argus := os.Args[1:]

	max_grade, err := strconv.Atoi(argus[0])
	if err != nil {
		log.Println("ERROR: max_grade, err := strconv.Atoi(argus[0]) FAILED")
		return
	}

	for i := 1; i < len(argus); i++ {
		argint, err := strconv.Atoi(argus[i])
		if err != nil {
			log.Println("ERROR: argint, err := strconv.Atoi(argus[i]) FAILED " + argus[i])
			return
		}
		if argint == -1 {
			fmt.Println(-1)
			return
		}
		if argint != 100 {
			fmt.Println(0)
			return
		}
	}

	fmt.Println(max_grade)
}
`,
	"/scorer/scoring.go": `package main

import (
	"fmt"
	"os"
	"strings"
)

// 0 ----> 100
func main() {
	output := os.Args[1]
	ans := os.Args[2]
	outputlist := strings.Fields(output)
	anslist := strings.Fields(ans)
	if len(outputlist) != len(anslist) {
		fmt.Println(-1)
		return
	}
	for i, v := range outputlist {
		if v != anslist[i] {
			fmt.Println(0)
			return
		}
	}
	fmt.Println(100)
}`,
	"problem.en.tex": "",
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a problem folder structure",
	Long:  `This command creates a problem folder structure in the current directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			commandfail("No file or problem-type given")
			return
		}
		if len(args) == 1 {
			commandfail("No problem-type given")
			return
		}
		problemName := strings.ToLower(args[0])
		problemType := strings.ToLower(args[1])
		if problemType == "standard" {
			createFileSys(standardFolders, standardFiles, standardContent, problemName)
		} else {
			commandfail("couldn't find problemType " + problemType)
			return
		}

	},
}

func createFileSys(folders []string, files []string, content map[string]string, problemName string) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	err = os.MkdirAll(problemName, os.ModePerm)

	if err != nil {
		commandfail("Failed creating folder " + problemName)
		return
	}

	for _, folder := range folders {
		nfolder := problemName + "/" + folder
		path := filepath.Join(cwd, nfolder)
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Printf("Failed to create %s: %v\n", nfolder, err)
		} else {
			fmt.Printf("Created: %s\n", nfolder)
		}
	}

	for _, file := range files {
		nfile := problemName + "/" + file
		path := filepath.Join(cwd, nfile)
		fmt.Println(nfile)
		err := os.WriteFile(path, []byte(content[file]), os.ModePerm)
		if err != nil {
			fmt.Printf("Failed to create %s: %v\n", nfile, err)
		} else {
			fmt.Printf("Created: %s\n", nfile)
		}
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
}
