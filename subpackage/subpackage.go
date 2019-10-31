package subpackage

import (
	"fmt"

	"github.com/spf13/viper"
)

func PrintConfig() {
	fmt.Println("dbpath ", viper.GetString("dbPath"))
	fmt.Println("isDebug ", viper.GetBool("isDebug"))
}
