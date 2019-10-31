// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io/ioutil"
// 	"os"

// 	toml "github.com/pelletier/go-toml"
// )

// var (
// 	baseConfigPath = "./baseConfig.toml"
// 	nodeTempPath   = "./"
// )

// func main() {
// 	// configByte, err := ioutil.ReadFile(baseConfigPath)
// 	// if err != nil {
// 	// 	panic("base config file is not found")
// 	// }
// 	// baseConfig, errLoad := toml.Load(string(configByte))
// 	// if errLoad != nil {
// 	// 	panic("base config file is not found")
// 	// }
// 	// baseConfig.Set("apiRPCPort", "4000")

// 	// f, osErr := os.Create("./configModified.toml")
// 	// if osErr != nil {
// 	// 	panic("failed to creata a new file")
// 	// }
// 	// defer f.Close()

// 	// f.Sync()
// 	// writer := bufio.NewWriter(f)
// 	// baseConfig.WriteTo(writer)
// 	// writer.Flush()

// 	configByte, err := ioutil.ReadFile(baseConfigPath)
// 	if err != nil {
// 		panic(fmt.Sprint("base config file is not found ", baseConfigPath))
// 	}
// 	baseConfig, errLoad := toml.Load(string(configByte))
// 	if errLoad != nil {
// 		panic(fmt.Sprint("failed to load baseConfig file ", baseConfigPath))
// 	}
// 	// baseConfig.Set("peerPort", 123)
// 	baseConfig.Set("apiRPCPort", 123)
// 	baseConfig.Set("myAddress", "adf")
// 	// baseConfig.Set("wellknownPeers", []string{"123"})
// 	baseConfig.Set("ownerAccountAddress", "zvzc")

// 	newConfigFilePath := fmt.Sprintf("%s/config.toml", nodeTempPath)
// 	f, osErr := os.Create(newConfigFilePath)
// 	if osErr != nil {
// 		panic(fmt.Sprint("failed to creata a new file ", osErr))
// 	}
// 	defer f.Close()

// 	_ = f.Sync()
// 	writer := bufio.NewWriter(f)
// 	_, errWrite := baseConfig.WriteTo(writer)
// 	if errWrite != nil {
// 		panic(fmt.Sprintf("failed to write to the new config file %s\n", newConfigFilePath))
// 	}
// 	writer.Flush()
// }

// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// )

// func check(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }

// func main() {

// 	// To start, here's how to dump a string (or just
// 	// bytes) into a file.
// 	d1 := []byte("hello\ngo\n")
// 	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
// 	check(err)

// 	// For more granular writes, open a file for writing.
// 	f, err := os.Create("./tmp/dat2")
// 	check(err)

// 	// It's idiomatic to defer a `Close` immediately
// 	// after opening a file.
// 	defer f.Close()

// 	// You can `Write` byte slices as you'd expect.
// 	d2 := []byte{115, 111, 109, 101, 10}
// 	n2, err := f.Write(d2)
// 	check(err)
// 	fmt.Printf("wrote %d bytes\n", n2)

// 	// A `WriteString` is also available.
// 	n3, err := f.WriteString("writes\n")
// 	fmt.Printf("wrote %d bytes\n", n3)

// 	// Issue a `Sync` to flush writes to stable storage.
// 	f.Sync()

// 	// `bufio` provides buffered writers in addition
// 	// to the buffered readers we saw earlier.
// 	w := bufio.NewWriter(f)
// 	n4, err := w.WriteString("buffered\ntesting double enter\n")
// 	fmt.Printf("wrote %d bytes\n", n4)

// 	n5, err := w.WriteString("test after buffer\n")
// 	fmt.Printf("wrote %d bytes\n", n5)

// 	// Use `Flush` to ensure all buffered operations have
// 	// been applied to the underlying writer.
// 	w.Flush()

// }

// package main

// import (
// 	"log"
// 	"net/http"
// 	"os"
// 	"time"

// 	"github.com/prometheus/client_golang/prometheus"
// 	"github.com/prometheus/client_golang/prometheus/promhttp"
// 	"github.com/shirou/gopsutil/mem"
// )

// func main() {
// 	// init log service
// 	logger := log.New(os.Stdout, "[Memory]", log.Lshortfile|log.Ldate|log.Ltime)

// 	// init http handler
// 	http.Handle("/metrics", promhttp.Handler())

// 	// init container
// 	memoryPercent := prometheus.NewGaugeVec(prometheus.GaugeOpts{
// 		Name: "memory_percent",
// 		Help: "memory use percent",
// 	},
// 		[]string{"percent"},
// 	)
// 	prometheus.MustRegister(memoryPercent)

// 	go func() {
// 		logger.Println("ListenAndServe at:0.0.0.0:2121")
// 		err := http.ListenAndServe("localhost:2121", nil)

// 		if err != nil {
// 			logger.Fatal("ListenAndServe: ", err)
// 		}
// 	}()

// 	// collection memory use percent
// 	for {
// 		logger.Println("start collect memory used percent!")
// 		v, err := mem.VirtualMemory()
// 		if err != nil {
// 			logger.Println("get memory use percent error:%s", err)
// 		}
// 		usedPercent := v.UsedPercent
// 		logger.Println("get memory use percent:", usedPercent)
// 		memoryPercent.WithLabelValues("usedMemory").Set(usedPercent)
// 		time.Sleep(time.Second * 2)
// 	}
// }

package main

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/zoobc/test/subpackage"
)

var ConfigFile = "./config.toml"

func init() {
	viper.SetConfigFile(ConfigFile)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	viper.Set("isDebug", true)
}

func main() {
	subpackage.PrintConfig()
}
