package compler

import (
    "io/ioutil"
    "log"
    "os"
    "os/exec"
    "path/filepath"
    "fmt"
    // "time"
)

func Comp(code string) {
    // src := `
    // package main
    // import (
    // "fmt"
    // "time"
    // )
    // func main() {
    //     time.Sleep(10 * time.Second)
    //     fmt.Println("Hello from dynamically compiled code!")
    // }
    // `
    // src := code

    tmpDir, _ := ioutil.TempDir("", "go-run")
    // Consider saving some compiled code for later analysis
    defer os.RemoveAll(tmpDir)

    srcFile := filepath.Join(tmpDir, "main.go")
    ioutil.WriteFile(srcFile, []byte(code), 0644)
    fmt.Println("Running exec")
    
    cmd := exec.Command("go", "run", srcFile)

    stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("OUT: %s\n", stdoutStderr)
    fmt.Println("Finished Running")
}
