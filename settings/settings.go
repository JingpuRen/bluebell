package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Init tip 使用viper来管理配置文件
func Init() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// tip : 这个路径中的./表示的是 D:\GoProjects\go_web 这个路径！！！这里一定要注意！！
	viper.AddConfigPath("./")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig() failed,err:%v\n", err)
		return err
	}

	// tip 支持配置热加载：也就是在运行时如果修改配置，我们是不需要重新运行代码的
	viper.WatchConfig()
	// tip 回调参数：当配置改变了的时候，就会触发下面的代码
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("Config file is modified ......")
	})
	return nil
}
