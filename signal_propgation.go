package main

import (
  "log"
  "os"
  "os/signal"
  "syscall"
)


// All goroutines (well, every channel registered) will receive a sighup
func main () {
  log.Printf("pid: %d", os.Getpid())
  outer := make(chan os.Signal)
  signal.Notify(outer, syscall.SIGINT)

  for i := 0; i < 10; i++ {
    n := i
    go func () {
      c := make(chan os.Signal)
      signal.Notify(c, syscall.SIGINT)
      <-c
      log.Printf("Received signal on goroutine %d", n)
    }()
  }

  <-outer
  log.Printf("Received signal (main)")
}

func DetachCommand_AvoidSignal() {

  cmd := exec.Command("memcached")
  // Detach to a different process group, so this daemon does not
  // get killed by a signal
  cmd.SysProcAttr = &syscall.SysProcAttr {
    Setpgid: true,
  }
  go cmd.Run()

  // So we need to kill it ourselves!
  cmd.Process.Signal(syscall.SIGTERM)
}