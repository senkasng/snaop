package util

import (
	"github.com/Unknwon/goconfig"
	//"github.com/labstack/echo/middleware"
	//"fmt"
)
var cfg *goconfig.ConfigFile
var err error
func init() {
	cfg,err = goconfig.LoadConfigFile("./config/config.ini")
	if err != nil {
		panic(err)
	}
}

func ParseGlobal() (global map[string]string) {
	global,err := cfg.GetSection("global")
	if err != nil {
		panic(err)
	}
	return global
}

func ParseObject()(map[string]map[string]string) {
	var object map[string]map[string]string
	//object = make(map[string]map[string]string)
	object = make(map[string]map[string]string)
	obj := cfg.GetSectionList()
	if err != nil {
		panic(err)
	}
	//fmt.Println(obj[1])
	for i :=1;i<len(obj);i++{
		object[obj[i]] = make(map[string]string)
		object[obj[i]],err = cfg.GetSection(obj[i])
		if err != nil {
			panic(err)
		}
	}
	/*for _,v := range object["t_cvs_kafka"] {
		fmt.Println(v)
	}*/
	return object
}

