/*
Copyright © 2025 Arvid Kristoffersson arvid.kristoffersson@icloud.com
*/

package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// Error messages

func commandfail(text string) {
	fmt.Println("ERROR: " + text)
}

func warning(text string) {
	fmt.Println("------------ WARNING -----------")
	fmt.Println(text)
	fmt.Println("--------------------------------")
}

// Random utility

func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	return !errors.Is(error, os.ErrNotExist)
}

func isProcessRunning(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}

	// Try sending signal 0 (non-destructive)
	err = process.Signal(syscall.Signal(0))
	return err == nil
}

func enableUnlimitedRecursion() {
	runcmd := exec.Command("bash", "-c", "ulimit -s unlimited")
	err := runcmd.Run()
	if err != nil {
		commandfail(`Failed running "ulimit -s unlimited"`)
		return
	}
}

// YAML Config Struct
type Config struct {
	Limits struct {
		TimeLimit      int `yaml:"time_limit"`
		CharacterLimit int `yaml:"character_limit"`
		MemoryLimit    int `yaml:"memory_limit"`
	} `yaml:"limits"`
	Scoring struct {
		GroupGrades   []int  `yaml:"group_grades"`
		OnReject      string `yaml:"on_reject"`
		IncludeSample string `yaml:"include_sample"`
	} `yaml:"scoring"`

	Type string `yaml:"type"`
}

func groupScores(problemName string) []int {
	data, err := os.ReadFile(problemName + "/config.yaml")
	if err != nil {
		commandfail("Failed reading config.yaml in groupScores")
		return []int{}
	}

	var config Config

	if err := yaml.Unmarshal(data, &config); err != nil {
		commandfail("Failed parsing config.yaml in groupScores, we: " + err.Error())
		return []int{}
	}

	return config.Scoring.GroupGrades
}

func onReject(problemName string) string {
	data, err := os.ReadFile(problemName + "/config.yaml")
	if err != nil {
		commandfail("Failed reading config.yaml in onReject")
		return ""
	}

	var config Config

	if err := yaml.Unmarshal(data, &config); err != nil {
		commandfail("Failed parsing config.yaml in onReject")
		return ""
	}

	return config.Scoring.OnReject
}

func timeLimit(problemName string) int {
	data, err := os.ReadFile(problemName + "/config.yaml")
	if err != nil {
		commandfail("Failed reading config.yaml in timeLimit")
		return 0
	}

	var config Config

	if err := yaml.Unmarshal(data, &config); err != nil {
		commandfail("Failed parsing config.yaml in timeLimit")
		return 0
	}

	return config.Limits.TimeLimit
}

func includeSample(problemName string) string {
	data, err := os.ReadFile(problemName + "/config.yaml")
	if err != nil {
		commandfail("Failed reading config.yaml in includeSample")
		return ""
	}

	var config Config

	if err := yaml.Unmarshal(data, &config); err != nil {
		commandfail("Failed parsing config.yaml in includeSample")
		return ""
	}

	return config.Scoring.IncludeSample
}

func compileCPP(dir string, exeFile string) bool {
	compileCmd := exec.Command("g++", dir, "-o", exeFile)
	compileCmd.Stdout = os.Stdout
	compileCmd.Stderr = os.Stderr

	err := compileCmd.Run()

	if err != nil {
		fmt.Println("Error starting process:", err)
		return false
	}

	return true
}

func CPPrunner(exeFile string, scores []int, problemName string) string {
	verdictScore := 0
	var verdictText string
	onreject := onReject(problemName)
	time_limit := timeLimit(problemName)
	if includeSample(problemName) == "yes" {
		fmt.Println("Testing sample")
		sampledir := problemName + "/data/sample/"
		var samplescores []string
		for i := 1; i < 1000; i++ {
			infiledir := sampledir + strconv.Itoa(i) + ".in"
			outfiledir := sampledir + strconv.Itoa(i) + ".ans"
			if !checkFileExists(infiledir) {
				break
			}
			if !checkFileExists(outfiledir) {
				break
			}

			inputData, err := os.ReadFile(infiledir)
			if err != nil {
				commandfail("Failed reading input from " + infiledir)
				return ""
			}

			enableUnlimitedRecursion()

			solution := exec.Command("./" + exeFile)

			solution.Stdin = bytes.NewBufferString(string(inputData))

			var outBuffer bytes.Buffer
			solution.Stdout = &outBuffer
			solution.Stderr = os.Stderr

			err = solution.Start()

			timeout := time.Duration(time_limit) * time.Second

			solution.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
			pid := solution.Process.Pid

			time.Sleep(timeout)

			if isProcessRunning(pid) {
				if err := syscall.Kill(pid, syscall.SIGKILL); err != nil {
					fmt.Println("Kill error:", err)
				}
			}

			if err != nil {
				commandfail("Failed running " + exeFile + " " + err.Error())
				return ""
			}

			outData, err := os.ReadFile(outfiledir)
			if err != nil {
				commandfail("Failed reading output from " + outfiledir)
				return ""
			}

			scoredir := problemName + "/scorer/scoring.go"

			scoring := exec.Command("go", "run", scoredir, outBuffer.String(), string(outData))

			var scoreBuffer bytes.Buffer

			scoring.Stdin = os.Stdin
			scoring.Stderr = os.Stderr
			scoring.Stdout = &scoreBuffer

			err = scoring.Run()

			if err != nil {
				return "RTE"
			}

			samplescores = append(samplescores, strings.ReplaceAll(scoreBuffer.String(), "\n", ""))
		}

		if len(samplescores) == 0 {
			fmt.Println("No sample given")
		} else {
			gradedir := problemName + "/grader/grading.go"
			args := append([]string{"run", gradedir, "1"}, samplescores...)
			grading := exec.Command("go", args...)

			var infoBuffer bytes.Buffer
			grading.Stderr = os.Stderr
			grading.Stdin = os.Stdin
			grading.Stdout = &infoBuffer

			err := grading.Run()

			if err != nil {
				commandfail("Failed running grader/grading.go for sample")
				return ""
			}

			if strings.ReplaceAll(infoBuffer.String(), "\n", "") == "0" {
				warning("Solution " + exeFile + " is wrong on sample")
			}

		}

	}

	for group := 1; group <= len(scores); group++ {
		fmt.Println("Testing group " + strconv.Itoa(group))
		groupdir := problemName + "/data/secret/group" + strconv.Itoa(group) + "/"
		if !checkFileExists(groupdir) {
			commandfail("There are more scores than groups")
			break
		}
		var groupscores []string
		for i := 1; i < 1000; i++ {
			infiledir := groupdir + strconv.Itoa(i) + ".in"
			outfiledir := groupdir + strconv.Itoa(i) + ".ans"
			if !checkFileExists(infiledir) {
				break
			}
			if !checkFileExists(outfiledir) {
				break
			}

			fmt.Print("Testing file " + infiledir + ": ")

			inputData, err := os.ReadFile(infiledir)
			if err != nil {
				commandfail("Failed reading input from " + infiledir)
				return ""
			}

			enableUnlimitedRecursion()

			solution := exec.Command("./" + exeFile)

			solution.Stdin = bytes.NewBufferString(string(inputData))

			var outBuffer bytes.Buffer
			solution.Stdout = &outBuffer
			solution.Stderr = os.Stderr

			err = solution.Start()

			if err != nil {
				return "RTE"
			}

			timeout := time.Duration(time_limit) * time.Second

			solution.SysProcAttr = &syscall.SysProcAttr{Setpgid: true} // Run independently
			pid := solution.Process.Pid

			time.Sleep(timeout)

			if isProcessRunning(pid) {
				if err := syscall.Kill(pid, syscall.SIGKILL); err != nil {
					fmt.Println("Kill error:", err)
				}
			}

			outData, err := os.ReadFile(outfiledir)
			if err != nil {
				commandfail("Failed reading output from " + outfiledir)
				return ""
			}

			scoredir := problemName + "/scorer/scoring.go"

			scoring := exec.Command("go", "run", scoredir, outBuffer.String(), string(outData))

			var scoreBuffer bytes.Buffer

			scoring.Stdin = os.Stdin
			scoring.Stderr = os.Stderr
			scoring.Stdout = &scoreBuffer

			err = scoring.Run()

			if err != nil {
				commandfail("Failed running scoring.go")
				return ""
			}

			scorestr := strings.ReplaceAll(scoreBuffer.String(), "\n", "")

			fmt.Print(scorestr + "\n")

			groupscores = append(groupscores, scorestr)

			if (scorestr == "0" || scorestr == "-1") && onreject == "stop" {
				break
			}
		}
		gradedir := problemName + "/grader/grading.go"
		args := append([]string{"run", gradedir, strconv.Itoa(scores[group-1])}, groupscores...)
		grading := exec.Command("go", args...)

		var gradeBuffer bytes.Buffer
		grading.Stdin = os.Stdin
		grading.Stderr = os.Stderr
		grading.Stdout = &gradeBuffer

		err := grading.Run()

		if err != nil {
			commandfail("Failed running grader/grading.go for sample")
			return ""
		}

		grade, err := strconv.Atoi(strings.ReplaceAll(gradeBuffer.String(), "\n", ""))
		if err != nil {
			commandfail("Failed converting gradeBuffer.String()," + gradeBuffer.String() + " to integer")
			return ""
		}
		if grade == -1 {
			verdictText = "TLE"
		} else {
			verdictScore += grade
		}
	}

	fmt.Println("totalVerdict: " + strconv.Itoa(verdictScore))

	if verdictScore == 0 && verdictText == "TLE" {
		return "TLE"
	}

	return strconv.Itoa(verdictScore)
}

func checkSolutions(problemName string) {

	verdicts, _ := os.ReadDir(problemName + "/solutions/")
	if len(verdicts) == 0 {
		warning("No solution given")
	}

	var score_verdicts []int
	var text_verdicts []string

	for _, verdict := range verdicts {
		if verdict.IsDir() {
			if x, err := strconv.Atoi(verdict.Name()); err == nil {
				score_verdicts = append(score_verdicts, x)
			} else {
				text_verdicts = append(text_verdicts, verdict.Name())
			}
		} else {
			commandfail("A file is in the solution directory")
			return
		}
	}

	for _, s := range text_verdicts {
		items, _ := os.ReadDir(problemName + "/solutions/" + s)
		if len(items) == 0 {
			commandfail("No solutions in " + s)
			return
		}

		solutiondir := problemName + "/solutions/" + s + "/"

		for _, item := range items {
			filedir := solutiondir + item.Name()
			if item.IsDir() {
				commandfail("Directory inside a verdict directory.")
				return
			} else {
				fmt.Println("Testing " + item.Name())

				if item.Name()[len(item.Name())-4:] != ".cpp" {
					commandfail(item.Name() + " is not a C++ file")
					return
				}

				exeFile := item.Name()[:len(item.Name())-4] //UPDATE when supporting other languages than C++

				compileCPP(filedir, exeFile)

				groupscores := groupScores(problemName)

				if len(groupscores) == 0 {
					return
				}

				real_verdict := CPPrunner(exeFile, groupscores, problemName)

				fmt.Println(real_verdict)
				if real_verdict != s {
					warning("Solution " + filedir + " of verdict " + s + " didn't achieve the correct verdict. It achieved: " + real_verdict)
				}

			}
		}
	}

	for _, s := range score_verdicts {
		items, _ := os.ReadDir(problemName + "/solutions/" + strconv.Itoa(s))
		if len(items) == 0 {
			commandfail("No solutions in " + strconv.Itoa(s))
			return
		}

		solutiondir := problemName + "/solutions/" + strconv.Itoa(s) + "/"

		for _, item := range items {
			filedir := solutiondir + item.Name()
			if item.IsDir() {
				commandfail("Directory inside a verdict directory.")
				return
			} else {
				fmt.Println("Testing " + item.Name())

				if item.Name()[len(item.Name())-4:] != ".cpp" {
					commandfail(item.Name() + " is not a C++ file")
					return
				}

				exeFile := item.Name()[:len(item.Name())-4] //UPDATE when supporting other languages than C++

				finished := compileCPP(filedir, exeFile)

				if !finished {
					return
				}

				groupscores := groupScores(problemName)

				if len(groupscores) == 0 {
					return
				}

				real_verdict := CPPrunner(exeFile, groupscores, problemName)
				fmt.Println(real_verdict)

				if real_verdict != strconv.Itoa(s) {
					warning("Solution " + filedir + " of verdict " + strconv.Itoa(s) + " didn't achieve the correct verdict")
				}
			}
		}

	}
}

// Cleaning

func isExecutable(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return info.Mode()&0111 != 0 && filepath.Ext(path) == ""
}

func removeExecutables(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if isExecutable(path) {
			fmt.Println("Removing:", path)
			if err := os.Remove(path); err != nil {
				return err
			}
		}

		return nil
	})
}

func cleanUp(dir string) {
	err := removeExecutables(dir)
	if err != nil {
		commandfail("Clean-up error, removeExecutables failed: " + err.Error())
		return
	}
}

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verifies a Hundis problem",
	Long: `A command that verifies the following properties of a Hundis problem:
Its solutions get the amount of points they should
Its statements are compilable
Everything is configured correctly
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			commandfail("No file named")
			return
		}
		problemName := args[0]

		checkSolutions(problemName)

		fmt.Println("Verification done!")
		fmt.Println("Commencing clean up...")

		cleanUp(".")
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)
}
