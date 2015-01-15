package main

import (
  "os"
  "os/exec"
  "io/ioutil"
  "time"
)

const WEAVE_CARRIER_PATH = "/sys/class/net/ethwe/carrier"

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

  check to see if the weave network is ready
  we do this by checking the path:

  /sys/class/net/ethwe/carrier

  to see if it contains "1"

  
*/
func isWeaveReady() bool {
  if _, err := os.Stat(WEAVE_CARRIER_PATH); os.IsNotExist(err) {
    return false
  } else {
    bytes, err := ioutil.ReadFile(WEAVE_CARRIER_PATH)
    if err != nil {
      return false
    }
    if(string(bytes[0]) == "1"){
      return true
    } else {
      return false
    }
  }
}

/*

  run isWeaveReady continously until it returns true

  wait for 1 second inbetween attempts
  
*/
func waitForWeave(){
  for !isWeaveReady() {
    time.Sleep(1 * time.Second)
  }
  return
}

func main() {
  waitForWeave()
  runEntryPoint()
}