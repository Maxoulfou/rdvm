/*
Copyright © 2022 BROCHIER MAXENCE maxence@brochier.xyz

*/
package config

import "os"

var (
	Filename     = "config.json"
	Directory, _ = os.Getwd()
	FilePathWin  = Directory + "\\" + Filename
	FilePathUnix = Directory + "/" + Filename
)
