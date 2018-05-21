package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"GLab/part01/serverer/user"
	"time"
)

func main() {
	http.HandleFunc("/product", product)
	http.HandleFunc("/reabuy/webMobile/buyerLogin.do", login)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("网络服务开启失败！！！")
	}
}

func product(w http.ResponseWriter, r *http.Request) {
	var str = `{
	"root": {
		"productID": 164953,
		"shopID": 2024,
		"productName": "兰芝口红",
		"keywords": "口红",
		"mainPic": ["1516871978427.png", "1516872079084.png"],
		"descPic": ["1516872087862.jpg", "1516872090067.jpg"],
		"prodProp": [{
			"price": 8500,
			"name": "狂炫酷吧吊炸天",
			"pic": ""
		}, {
			"price": 9100,
			"name": "低调奢华有内涵",
			"pic": ""
		}],
		"productRate": 0.00636,
		"exchangeCurrency": "人民币",
		"price": 8000.0,
		"discount": 1.0,
		"currency": "韩币",
		"stock": 5,
		"prodDesc": "<div class=\"pd\"><div style=\"font-size: 13px;\"><font face=\"helvetica, arial, verdana, sans-serif\"><span style=\"font-size: 13px;\">产品名称 : laneige/兰芝水润淡彩护唇膏</span></font></div><div style=\"font-size: 13px;\"><font face=\"helvetica, arial, verdana, sans-serif\"><span style=\"font-size: 13px;\">规格类型 : 正常规格</span></font></div><div style=\"font-size: 13px;\"><font face=\"helvetica, arial, verdana, sans-serif\"><span style=\"font-size: 13px;\">化妆品净含量 : 3.5g</span></font></div><div style=\"font-size: 13px;\"><font face=\"helvetica, arial, verdana, sans-serif\"><span style=\"font-size: 13px;\">品牌 : Laneige/兰芝</span></font></div><div style=\"font-size: 13px;\"><font face=\"helvetica, arial, verdana, sans-serif\"><span style=\"font-size: 13px;\">产地 : 韩国</span></font></div><div style=\"font-size: 13px;\"><font face=\"helvetica, arial, verdana, sans-serif\"><span style=\"font-size: 13px;\">功效 ：润唇，保湿</span></font></div><div style=\"font-size: 13px;\"><font face=\"helvetica, arial, verdana, sans-serif\"><span style=\"font-size: 13px;\">分类：润唇膏，啫喱</span></font></div><div><font size=\"2\" face=\"helvetica, arial, verdana, sans-serif\"><br></font></div><div><font size=\"2\" face=\"helvetica, arial, verdana, sans-serif\">产品详情</font></div><div style=\"font-size: 13px;\"><font face=\"helvetica, arial, verdana, sans-serif\"><span style=\"font-size: 13px;\">1. 越涂抹越保湿，水滴染色色彩的水彩画唇釉</span></font></div><div style=\"font-size: 13px;\"><font face=\"helvetica, arial, verdana, sans-serif\"><span style=\"font-size: 13px;\">2. 保湿力升级</span></font></div><div style=\"font-size: 13px;\"><font face=\"helvetica, arial, verdana, sans-serif\"><span style=\"font-size: 13px;\">3. 柔和清爽的使用感</span></font></div><div style=\"font-size: 13px;\"><font face=\"helvetica, arial, verdana, sans-serif\"><span style=\"font-size: 13px;\">4. 呈现最适合自己的专属色彩</span></font></div></div>",
		"shopName": "兰芝",
		"shopDomain": "test01",
		"shopStatus": "0",
		"shopDev": {
			"devID": "40288a85613035a601613056f0150000",
			"shopID": 2024,
			"startQuan": 1.0,
			"startPrice": 3000.0,
			"addQuan": 0.2,
			"addPrice": 50.0,
			"freePrice": 30000.0,
			"regionOnly": null
		}
	},
	"resMsg": "0"
}`
	w.Write([]byte(str))
}

func login(resp http.ResponseWriter, req *http.Request) {
	time.Sleep(time.Second * 5)
	json.NewEncoder(resp).Encode(user.LoginResult{"1", user.Buyer{}})
}
