package subpackage

import "flag"

func PrintConfig() string {
	return flag.Lookup("config-path").Value.(flag.Getter).Get().(string)
}
