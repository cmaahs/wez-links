package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/muesli/termenv"
)

func decorateNamespaceLine(line string, term *termenv.Output) string {

	re := regexp.MustCompile(`\S+`)

	submatchall := re.FindAllString(line, -1)
	// for _, element := range submatchall {
	// 	fmt.Println(element)
	// }

	namespace := term.Hyperlink(fmt.Sprintf("<bash:NAMESPACE=%s>", submatchall[0]), submatchall[0])

	line = strings.Replace(line, submatchall[0], namespace, 1)
	return line
}
