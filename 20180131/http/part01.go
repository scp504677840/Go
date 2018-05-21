package main

import (
	"net/http"
	"log"
	"fmt"
	"reflect"
	"io/ioutil"
	url2 "net/url"
	"encoding/json"
	"strings"
	"math/rand"
	"time"
	"math"
	"os"
	"io"
)

/*
productCondition.keywords = keyword
productCondition.start = 0
productCondition.limit = 10
productCondition.region = "KO"
productCondition.minPrice = 0.01
productCondition.status = "0"
 */
func main() {

	//搜索产品
	searchProduct()

	//搜索店铺产品
	//products := searchShopProducts()

	//productNames := []string{}
	//以下是未上传的
	//productNames = append(productNames, `{"zh":"蓝色保湿补水果冻面膜 10片/盒 保湿滋润补水 BANOBAGI 巴诺巴奇"}`)
	//productNames = append(productNames, `{"zh":"绿色舒缓镇定果冻面膜 10片/盒 BANOBAGI 巴诺巴奇"}`)
	//productNames = append(productNames, `{"zh":"橘色维他基因美白果冻面膜 10片/盒 BANOBAGI 巴诺巴奇"}`)
	//productNames = append(productNames, `{"zh":"红色弹力紧致果冻面膜 10片/盒 BANOBAGI 巴诺巴奇"}`)
	//productNames = append(productNames, `{"zh":"皮肤科锡纸面膜 铁盒装 10片/盒 BANOBAGI 巴诺巴奇"}`)
	//productNames = append(productNames, `{"zh":"茶秘水润亮白焕活面膜贴 5片/盒 VIVLAS 唯兰颂"}`)
	//productNames = append(productNames, `{"zh":"茶秘保湿面膜 5片/盒 VIVLAS 唯兰颂"}`)
	//productNames = append(productNames, `{"zh":"魔法精华双重浓缩精华 50ml SUM37°呼吸 苏秘"}`)
	//productNames = append(productNames, `{"zh":"魔法精华BB气垫 送替换装 SUM37°呼吸 苏秘"}`)
	/*for _, v := range productNames {
		saveProduct(v)
	}*/
}

func saveProductImage(imgUrl string, photoName string) {
	file, err := os.Create(fmt.Sprintf("./%s", photoName))
	if err != nil {
		fmt.Println("文件创建失败！！！")
		return
	}
	resp, err := http.Get(imgUrl)
	if err != nil {
		fmt.Println("请求图片失败！！！")
		return
	}

	written, err := io.Copy(file, resp.Body)

	if err != nil {
		fmt.Println("图片写入文件失败！！！")
		return
	}
	fmt.Println(written)
}

func saveProduct(productName string) {
	var url = "http://192.168.10.80:8080/reabuy/shopManager/productManager/saveProduct.do"
	client := &http.Client{}

	//随机数
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	product := &SaveProduct{}
	product.CategoryID = 457
	product.CategoryPath = "LV1_453,LV2_454,LV3_457"
	product.Currency = "CURR_KRW"
	product.DeliveryProp = "1.2"
	product.Discount = 0.8
	product.Keywords = "浓缩 咖啡 唤肤 修复 精华 面膜"
	product.Price = math.Trunc(r.Float64() * 1000000)
	product.ProductName = productName
	product.PublishDate = 1517470425890
	product.RegionPath = "LV0_Region_1434610561508_KO/LV1_Region_1435023095132/LV2_Region_1438050788522"
	product.ShopID = 2025
	product.Stock = 10

	bytes, err := json.Marshal(product)
	if err != nil {
		log.Panicln("json编码错误！！！", err)
		return
	}

	fmt.Println(string(bytes))

	resp, err := client.Post(url, "application/json", strings.NewReader(string(bytes)))

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Panicln("请求体失败！！！", err)
		return
	}

	fmt.Println(string(body))

}

/*
上架产品
 */
func shelves(productID int) {
	var url = "http://192.168.10.80:8080/reabuy/shopManager/productManager/modStatus.do"
	client := &http.Client{}

	values := url2.Values{}
	values.Set("productID", string(productID))
	values.Set("changeStatus", "0")
	resp, err := client.PostForm(url, values)

	if err != nil {
		log.Panicln("请求出错！！！", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Panicln("请求体失败！！！", err)
		return
	}

	fmt.Println(string(body))
}

/*
搜索店铺后台产品
 */
func searchShopProducts() (products []Product) {
	var url = "http://192.168.10.80:8080/reabuy/shopManager/productManager/queryProList.do?_dc=1517473820075&productIDs=&keywords=&startDate=&endDate=&productName=&status=&shopID=2025&page=1&start=0&limit=100"
	client := &http.Client{}

	resp, err := client.Get(url)

	if err != nil {
		log.Panicln("请求出错！！！", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Panicln("请求体失败！！！", err)
		return
	}

	fmt.Println(string(body))

	result := &SearchProductResult{}

	err = json.Unmarshal(body, result)

	if err != nil {
		log.Panicln("json解析失败！！！", err)
		return
	}

	products = result.Root
	return
}




