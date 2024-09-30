package compler

import (
    "io/ioutil"
    // "log"
    "os"
    "os/exec"
    "path/filepath"
    "fmt"
    // "time"
    // "go_ast/aster"
)

func Comp(code string) {

    tmpDir, _ := ioutil.TempDir("", "go-run")
    // Consider saving some compiled code for later analysis
    defer os.RemoveAll(tmpDir)

    srcFile := filepath.Join(tmpDir, "main.go")
    ioutil.WriteFile(srcFile, []byte(code), 0644)
    fmt.Println("Running exec")
    
    cmd := exec.Command("go", "run", srcFile)

    stdoutStderr, _ := cmd.CombinedOutput()

	fmt.Printf("OUT: %s\n", stdoutStderr)
    fmt.Println("Finished Running")
}
