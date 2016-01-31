package main

import (
	"fmt"
	"strconv"
"strings"
)

func main()  {
	r := strings.NewReplacer(
		"yyyy"  ,   strconv.Itoa(1394),
		"yyy"   ,   strconv.Itoa(1394),
		"yy"    ,   strconv.Itoa(1394)[2:],
		"y"     ,   strconv.Itoa(1394),
		"MMM"   ,   "فروردین",
		"MMD"   ,   "حمل",
		"MM"    ,   fmt.Sprintf("%02d", 2),
		"M"     ,   strconv.Itoa(2),
	)

	fmt.Println(r.Replace("yyyy yy y MMM MMD MM M yyY"))
}
