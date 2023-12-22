package main

import "os"

func main() {
	var rmdirs []func()
	dirs := []string{}
	for i := 0; i < len(dirs); i++ {
		j := i // NOTE: necessary for the user in later use
		os.MkdirAll(dirs[i], 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dirs[j])
		})
	}
}
