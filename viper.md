# 配置文件学习
一个项目总会有很多东西需要写到配置文件中，避免修改的时候修改代码，或者某些敏感信息不写入代码中。

go 语言中比较合适的管理配置文件的库如viper，下面我们就好好学习以下viper
```go
package main

import (
	"flag"
	"fmt"
    "github.com/spf13/viper"
)

var (
	config = flag.String("config", "config", "配置文件名称，默认 config")
)

func main() {
	// 配置文件名称
	viper.SetConfigName(*config)
	// 查找文件,可以配置多个路径
	viper.AddConfigPath("./")
	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: % \n",err))
	}
	// 监控文件变化
	viper.WatchConfig()
    // 读取配置内容
	environment := viper.GetBool("security.enabled")
	fmt.Println("security.enabled: ",environment)


}
```
以上即列出基于「viper」对配置文件的处理，这里重点解释以下监控文件变化：
先来追踪一下源码实现
```go
func WatchConfig() { v.WatchConfig() }

func (v *Viper) WatchConfig() {
	initWG := sync.WaitGroup{}
	initWG.Add(1)
	go func() {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}
		defer watcher.Close()
		// we have to watch the entire directory to pick up renames/atomic saves in a cross-platform way
		filename, err := v.getConfigFile()
		if err != nil {
			log.Printf("error: %v\n", err)
			initWG.Done()
			return
		}

		configFile := filepath.Clean(filename)
		configDir, _ := filepath.Split(configFile)
		realConfigFile, _ := filepath.EvalSymlinks(filename)

		eventsWG := sync.WaitGroup{}
		eventsWG.Add(1)
		go func() {
			for {
				select {
				case event, ok := <-watcher.Events:
					if !ok { // 'Events' channel is closed
						eventsWG.Done()
						return
					}
					currentConfigFile, _ := filepath.EvalSymlinks(filename)
					// we only care about the config file with the following cases:
					// 1 - if the config file was modified or created
					// 2 - if the real path to the config file changed (eg: k8s ConfigMap replacement)
					const writeOrCreateMask = fsnotify.Write | fsnotify.Create
					if (filepath.Clean(event.Name) == configFile &&
						event.Op&writeOrCreateMask != 0) ||
						(currentConfigFile != "" && currentConfigFile != realConfigFile) {
						realConfigFile = currentConfigFile
						err := v.ReadInConfig()
						if err != nil {
							log.Printf("error reading config file: %v\n", err)
						}
						if v.onConfigChange != nil {
							v.onConfigChange(event)
						}
					} else if filepath.Clean(event.Name) == configFile &&
						event.Op&fsnotify.Remove&fsnotify.Remove != 0 {
						eventsWG.Done()
						return
					}

				case err, ok := <-watcher.Errors:
					if ok { // 'Errors' channel is not closed
						log.Printf("watcher error: %v\n", err)
					}
					eventsWG.Done()
					return
				}
			}
		}()
		watcher.Add(configDir)
		initWG.Done()   // done initializing the watch in this go routine, so the parent routine can move on...
		eventsWG.Wait() // now, 的的wait for event loop to end in this go-routine...
	}()
	initWG.Wait() // make sure that the go routine above fully ended before returning
}

```