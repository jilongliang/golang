package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	httpUrl = "https://j.autohome.com.cn/pcplatform/common/getBrandInfo"
)

//处理错误函数
func checkErrs(err error) {
	if err != nil {
		panic(err)
	}
}

/**
* Jsons结构体
*/
type Jsons struct {
	Message string `json:"message"`
	Result  Result `json:"result"`
}
/**
 * Result结构体
 */
type Result struct {
	A []Brand `json:"A"`
	B []Brand `json:"B"`
	C []Brand `json:"C"`
	D []Brand `json:"D"`
	E []Brand `json:"E"`
	F []Brand `json:"F"`
	G []Brand `json:"G"`
	H []Brand `json:"H"`
	I []Brand `json:"I"`
	J []Brand `json:"J"`
	K []Brand `json:"K"`
	L []Brand `json:"L"`
	M []Brand `json:"M"`
	N []Brand `json:"N"`
	O []Brand `json:"O"`
	P []Brand `json:"P"`
	Q []Brand `json:"Q"`
	R []Brand `json:"R"`
	S []Brand `json:"S"`
	T []Brand `json:"T"`
	U []Brand `json:"U"`
	V []Brand `json:"V"`
	W []Brand `json:"W"`
	X []Brand `json:"X"`
	Y []Brand `json:"Y"`
	Z []Brand `json:"Z"`
}

/**
 * Brand品牌结构体
 */
type Brand struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Bfirstletter string `json:"bfirstletter"`
	Logo         string `json:"logo"`
}

func brands() {
	//读取Api数据
	resp, err := http.Get(httpUrl)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	//json转化成map
	b := Jsons{}
	json.Unmarshal([]byte(data), &b)
	fmt.Println(b)
}

func brands1() {
	//读取Api数据
	resp, err := http.Get(httpUrl)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	//json转化成map
	m := make(map[string]interface{})
	json.Unmarshal([]byte(data), &m)
	fmt.Println(m["result"].(map[string]interface{})["A"].([]interface{})[3].(map[string]interface{})["name"])
}

/**
 * 对象化处理JSON值
 */
func brandsVo() {
	var jsonsRet Jsons
	//读取Api数据
	resp, err := http.Get(httpUrl)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(data), &jsonsRet)

	//===============判断是否存在个值再去读取处理key====================
	resultMakeMap := make(map[string]interface{})
	json.Unmarshal([]byte(data), &resultMakeMap)

	resultMap := resultMakeMap["result"].(map[string]interface{})

	//判断是否存A的值
	if _, ok := resultMap["A"]; ok {
		resultObj := jsonsRet.Result
		//遍历的值
		for a := range resultObj.A {
			fmt.Println("车的名称:", resultObj.A[a].Name, "\t,车的Logo:", resultObj.A[a].Logo)
		}
	}

}

/*func atoz() [...]string {
	//arrys := [...] string{"A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"}
}*/

func main() {
	//brands()
	//brands1()
	brandsVo()
}
