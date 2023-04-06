package main

import "strings"

func inputType(line string) string {

	noSpaces := strings.ReplaceAll(line, " ", "")
	switch noSpaces {
	case "NAMEREADYSTATUSRESTARTSAGE":
		return "pod_list"
	case "NAMETYPEDATAAGE":
		return "secret_list"
	case "NAMESTATUSAGE":
		return "namespace_list"
	}

	return ""
}
