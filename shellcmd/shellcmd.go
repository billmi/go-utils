package shellcmd

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

/**
  Base on linux sys to exec shell
  @Author Bill

  demo :
  		ExecCommand("cat /home/ubuntu/123.txt")
*/
func ExecCommand(strCommand string) string {
	cmd := exec.Command("/bin/bash", "-c", strCommand)
	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		fmt.Println("Execute failed when Start:" + err.Error())
		return ""
	}
	out_bytes, _ := ioutil.ReadAll(stdout)
	stdout.Close()
	if err := cmd.Wait(); err != nil {
		fmt.Println("Execute failed when Wait:" + err.Error())
		return ""
	}
	return strings.TrimSpace(string(out_bytes))
}