package phone

import (
	"fmt"
	"github.com/antchfx/xquery/html"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (p *Phone) Query_360shoujiweishi() (pr Phone, err error) {
	qurl := "https://www.so.com/s?q=" + p.PhoneNumber
	pr = *p
	pr.Index = 1
	timeout := time.Duration(3 * time.Second) //设置超时3秒
	client := http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("GET", qurl, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return pr, err
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	bodystr := string(body)
	//fmt.Println(bodystr)

	// 判断是否有返回电话相关的信息
	pmarkstr := `mohe-mobileInfoContent`
	pmark := strings.Index(bodystr, pmarkstr)
	//fmt.Println(pmark)

	//模式判断
	// 360搜索结果目前发现有2个模式
	// 模式1 没有标记的情况 例如用 051288178411 cx.shouji.360.cn 关键字 class="mh-search"
	// 模式2 有标记的情况 例如用 053266114000 关键字 class="mohe-tips"
	//pmodekey1 := `class="mh-search"`
	//pmodemark1 := strings.Index(bodystr,pmodekey1)
	//pmodekey2 := `class="mohe-tips"`
	//pmodemark2 := strings.Index(bodystr,pmodekey2)
	//
	//fmt.Println(pmodemark1, pmodemark2)

	if pmark > 0 {
		//如果找到关键字才继续
		root, err := htmlquery.Parse(strings.NewReader(bodystr))
		if err != nil {
			return pr, err
		}
		node := htmlquery.FindOne(root, "//*[@class='mohe-mobileInfoContent']")
		pr.From = htmlquery.InnerText(htmlquery.FindOne(node, "//*[@class='mohe-sjws']")) //来源渠道
		pandc := htmlquery.InnerText(htmlquery.FindOne(node, "//div[1]/span[2]"))         //省份和城市在一起的需要切割
		pandcs := strings.TrimSpace(pandc)
		re1, _ := regexp.Compile(`\n`)
		pandcs = re1.ReplaceAllString(pandcs, ":")
		pandcArr := strings.Split(pandcs, ":")
		if len(pandcArr) > 1 {
			pr.Location.Province = strings.TrimSpace(pandcArr[0]) //省份
			pr.Location.City = strings.TrimSpace(pandcArr[1])     //城市
		}

		pr.Tag.TagName = strings.TrimSpace(htmlquery.InnerText(htmlquery.FindOne(node, "//*[@class='mohe-ph-mark']"))) //标记
		tagCnts := htmlquery.InnerText(htmlquery.FindOne(node, "//*[@class='mohe-tips']/span[2]/b"))
		if tagCnts != "" {
			tagCnt, err1 := strconv.Atoi(tagCnts)
			if err1 != nil {
				tagCnt = 0
			}
			pr.Tag.TagCnt = tagCnt //标记人数 位

		}

	}

	return pr, err
}
