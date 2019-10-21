//chinese weather search
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	simplejson "github.com/bitly/go-simplejson"
)

//中国天气预报结构体
type WeatherChina struct {
	City    string "city"    //城市名称
	Cityid  string "cityid"  //城市编号
	Temp    string "temp"    //温度
	WD      string "WD"      //风向
	WS      string "WS"      //风速
	SD      string "SD"      //湿度
	WSE     string "WSE"     //
	Time    string "time"    //时间
	IsRadar string "isRadar" //
	Radar   string "Radar"   //
	Njd     string "njd"
	Qy      string "qy"
}

func (wc WeatherChina) printWeather() {
	fmt.Println("-------------------------")
	fmt.Println("City    :", wc.City)
	fmt.Println("CityId  :", wc.Cityid)
	fmt.Println("Temp    :", wc.Temp)
	fmt.Println("WD      :", wc.WD)
	fmt.Println("WS      :", wc.WS)
	fmt.Println("SD      :", wc.SD)
	fmt.Println("WSE     :", wc.WSE)
	fmt.Println("Time    :", wc.Time)
	fmt.Println("IsReadar:", wc.IsRadar)
	fmt.Println("Radar   :", wc.Radar)
	fmt.Println("Njd     :", wc.Njd)
	fmt.Println("Qy      :", wc.Qy)
	fmt.Println("-------------------------")
}

//将json转换成struct
func ResolveWeatherJson(weatherJson string, wc *WeatherChina) {
	js, err := simplejson.NewJson([]byte(weatherJson))
	if err != nil {
		fmt.Println("NewJson err")
	}
	wi := js.Get("weatherinfo")
	wc.City = wi.Get("city").MustString()
	wc.Cityid = wi.Get("cityid").MustString()
	wc.Temp = wi.Get("temp").MustString()
	wc.WD = wi.Get("WD").MustString()
	wc.WS = wi.Get("WS").MustString()
	wc.SD = wi.Get("SD").MustString()
	wc.WSE = wi.Get("WSE").MustString()
	wc.Time = wi.Get("time").MustString()
	wc.IsRadar = wi.Get("isRadar").MustString()
	wc.Radar = wi.Get("radar").MustString()
	wc.Njd = wi.Get("njd").MustString()
	wc.Qy = wi.Get("qy").MustString()
}

//get weather json data with url
func getChinaWeather(url string) string {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("http Get err")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("resp>Body readAll err")
	}
	jsonBody := string(body)
	return jsonBody
}

//get url with paramater cityid
//cityid search url:http://www.cnblogs.com/wangjingblogs/p/3192953.html
func getWeatherUrlByCityid(cityid string) string {
	return "http://www.weather.com.cn/data/sk/" + cityid + ".html"
}

//主函数
func main() {
	zoncode := "101090101"
	if len(os.Args) == 1 {
	} else if len(os.Args) == 2 {
		name := strings.TrimSpace(os.Args[1])
		zoncode = getCityIdByName(name)
		if zoncode == "" {
			fmt.Println("未找到城市[", name, "]")
			return
		}
	} else {
		fmt.Println("参数错误")
		return
	}
	url := getWeatherUrlByCityid(zoncode)
	fmt.Println(url)
	jsonBody := getChinaWeather(url)
	var wc WeatherChina
	ResolveWeatherJson(jsonBody, &wc)
	wc.printWeather()
}