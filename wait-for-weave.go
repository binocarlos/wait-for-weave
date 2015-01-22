package main

import (
  "os"
  "os/exec"
  "fmt"
  "github.com/zettio/weave/net"
)

const WEAVE_INTERFACE_NAME = "ethwe"
const WAIT_FOR_SECONDS = 10
const QUIT_IMMEDIATELY_VAR = "WAIT_FOR_WEAVE_QUIT"

/*

  get the cli arguments as an array and remove the
  first element leaving us with the intended entrypoint

  then run that entrypoint attaching stdin, stdout and stderr
  
*/
func runEntryPoint() {
  entryPoint := make([]string, len(os.Args))
  copy(entryPoint, os.Args)
  // remove blocker from args[0]
  entryPoint = append(entryPoint[:0], entryPoint[1:]...)
  // get and remove the actual command from args[0]
  commandString := entryPoint[0]
  entryPoint = append(entryPoint[:0], entryPoint[1:]...)
  cmd := exec.Command(commandString, entryPoint...)
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  cmd.Run()
}

/*

  If the ethwe interface is not found after 10 seconds then write to stderr and exit with non-zero code
  
*/
func main() {

  /*
  
    this is so we can run an initial version of this container making its volumes available
    to other containers using --volumes-from
    
  */
  quitFlag := os.Getenv(QUIT_IMMEDIATELY_VAR)

  if quitFlag == "yes" {
    fmt.Fprintln(os.Stdout, "exiting without waiting because WAIT_FOR_WEAVE_QUIT == yes")
    os.Exit(0)
  }

  if _, err := net.EnsureInterface(WEAVE_INTERFACE_NAME, WAIT_FOR_SECONDS); err != nil {
    a := fmt.Sprint("interface ", WEAVE_INTERFACE_NAME, " not found after ", WAIT_FOR_SECONDS, " seconds")
    fmt.Fprintln(os.Stderr, a)
    os.Exit(1)
  } else {
    runEntryPoint()
  }
}