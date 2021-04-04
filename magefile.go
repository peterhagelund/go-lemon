// +build mage

package main

import "github.com/magefile/mage/sh"

func Clean() error {
	return sh.RunV("go", "clean", "./pkg/logger")
}

func Build() error {
	return sh.RunV("go", "build", "./pkg/logger")
}
