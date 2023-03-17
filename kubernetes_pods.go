package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/muesli/termenv"
)

func decoratePodLine(line string, term *termenv.Output) string {

	re := regexp.MustCompile(`\S+`)

	submatchall := re.FindAllString(line, -1)
	// for _, element := range submatchall {
	// 	fmt.Println(element)
	// }

	pod := term.Hyperlink(fmt.Sprintf("<bash:%s>", submatchall[0]), submatchall[0])
	del := term.Hyperlink(fmt.Sprintf("<bash:kubectl delete pod/%s -n ${NAMESPACE}>", submatchall[0]), submatchall[1])
	run := term.Hyperlink(fmt.Sprintf("<bash:kubectl exec -it %s -n ${NAMESPACE} -- /bin/bash>", submatchall[0]), submatchall[2])

	line = strings.Replace(line, submatchall[0], pod, 1)
	line = strings.Replace(line, submatchall[1], del, 1)
	line = strings.Replace(line, submatchall[2], run, 1)
	return line
}
