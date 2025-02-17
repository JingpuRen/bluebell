package snowflake

import (
	"time"

	"github.com/spf13/viper"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func Init() (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", viper.GetString("snowflake.startTime"))
	if err != nil {
		return
	}
	snowflake.Epoch = st.UnixNano() / 1000000

	node, err = snowflake.NewNode(viper.GetInt64("snowflake.machineID"))

	return
}

func GenID() int64 {
	return node.Generate().Int64()
}
