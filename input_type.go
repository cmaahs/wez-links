package main

func inputType(line string) string {
	switch line {
	case "NAME                                                   READY   STATUS    RESTARTS   AGE":
		return "pod_list"
	case "NAME                                               TYPE                                  DATA   AGE":
		return "secret_list"
	}
	return ""
}
