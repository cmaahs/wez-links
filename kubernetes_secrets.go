package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/muesli/termenv"
)

func decorateSecretLine(line string, term *termenv.Output) string {

	re := regexp.MustCompile(`\S+`)

	submatchall := re.FindAllString(line, -1)
	// for _, element := range submatchall {
	// 	fmt.Println(element)
	// }

	secret := term.Hyperlink(fmt.Sprintf("<bash:%s>", submatchall[0]), submatchall[0])
	view := term.Hyperlink(fmt.Sprintf("<bash:kubectl view-secret %s -a -n ${NAMESPACE}>", submatchall[0]), submatchall[1])

	line = strings.Replace(line, submatchall[0], secret, 1)
	line = strings.Replace(line, submatchall[1], view, 1)
	return line
}
