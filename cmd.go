// +build windows

/*
	osver.go
	-John Taylor
	May-2-2020

	Display the Windows OS version with the version being compatible with
	the Windows Nano Server Docker image

	MIT LICENSE: https://github.com/jftuga/osver/blob/master/LICENSE
	Nano Server Docker: https://hub.docker.com/_/microsoft-windows-nanoserver
*/

package main

import (
	"fmt"
	"log"

	"golang.org/x/sys/windows/registry"
)

func main() {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	maj, _, err := k.GetIntegerValue("CurrentMajorVersionNumber")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("CurrentMajorVersionNumber: %d\n", maj)

	min, _, err := k.GetIntegerValue("CurrentMinorVersionNumber")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("CurrentMinorVersionNumber: %d\n", min)

	cb, _, err := k.GetStringValue("CurrentBuild")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("CurrentVersion: %s\n", cb)

	ubr, _, err := k.GetIntegerValue("UBR")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("UBR: %d\n", ubr)

	fullVersion := fmt.Sprintf("%d.%d.%s.%d", maj, min, cb, ubr)
	fmt.Println(fullVersion)
}
