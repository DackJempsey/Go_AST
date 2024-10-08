package compler

import (
	"io/ioutil"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"go_ast/aster"
)

func RunReturn(srcFile string) string {

	fmt.Println("Sending To Aster: ", srcFile)
	// srcFile is path to file

	goodCode, tmpDir := aster.Fix(srcFile)
	defer os.RemoveAll(tmpDir)

	fmt.Println("Running exec")

	cmd := exec.Command("go", "run", goodCode)

	stdoutStderr, _ := cmd.CombinedOutput()

	fmt.Printf("OUT: %s\n", stdoutStderr)
	fmt.Println("Finished Running")
	return string(stdoutStderr)
}

func Comp(user_input string) string {
	importsFile := filepath.Join("./go_templates", "imports")
	mainFile := filepath.Join("./go_templates", "main")

	import_data, _ := os.ReadFile(importsFile)
	main_data, _ := os.ReadFile(mainFile)

	tmpDir, _ := ioutil.TempDir("", "go-run")
	// Consider saving some compiled code for later analysis
	defer os.RemoveAll(tmpDir)
	srcFile := filepath.Join(tmpDir, "main.go")
	file, _ := os.OpenFile(srcFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	file.WriteString(string(import_data))
	file.WriteString(string(main_data))

	file.WriteString(user_input)

	std := RunReturn(srcFile)

	return std
}
