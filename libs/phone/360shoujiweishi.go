package phone

import (
	"fmt"
	"github.com/antchfx/xquery/html"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (p *Phone) Query_360shoujiweishi() (pr Phone, err error) {
	qurl := "https://www.so.com/s?q=" + p.PhoneNumber
	pr = *p
	pr.Index = 1
	cj, _ := cookiejar.New(nil)
	timeout := time.Duration(6 * time.Second) //设置超时6秒
	client := http.Client{
		Timeout: timeout,
		Jar: cj,
	}
	req, err := http.NewRequest("GET", qurl, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36")
	//req.Header.Del("Cookie")
	resp, err := client.Do(req)
	if err != nil {
		logs.Warn(err)
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
	pmarkstr := `mohe-mobilecheck`
	pmark := strings.Index(bodystr, pmarkstr)
	//fmt.Println(pmark)

	//模式判断
	// 360搜索结果目前发现有3个模式
	// 模式1 没有标记的情况 例如用 051288178411 cx.shouji.360.cn 关键字 class="mh-search"
	// 模式2 有标记的情况 例如用 15555555555 关键字 class="mh-hy-tips" 360黄页
	// 模式3 有标记的情况 例如用 053266114000 关键字 class="mohe-tips"
	//  id="mohe-biu_kefudianhua"
	pmodekey1 := `class="mh-search"`
	pmodemark1 := strings.Index(bodystr,pmodekey1)
	pmodekey2 := `class="mh-hy-tips"`
	pmodemark2 := strings.Index(bodystr,pmodekey2)
	pmodekey3 := `class="mohe-tips"`
	pmodemark3 := strings.Index(bodystr,pmodekey3)
	//
	//fmt.Println(pmodemark1, pmodemark2)

	if pmark > 0 {
		//如果找到关键字才继续
		root, err := htmlquery.Parse(strings.NewReader(bodystr))
		if err != nil {
			logs.Warn(err)
			return pr, err
		}
		node := htmlquery.FindOne(root, "//*[@id='mohe-mobilecheck']")
		//logs.Info(bodystr)

		// 模式1
		if pmodemark1 > 0 {
			pr.From = htmlquery.InnerText(htmlquery.FindOne(node, "//*[@class='src']/p[2]/a")) //来源渠道
			pandc := htmlquery.InnerText(htmlquery.FindOne(node, "//*[@class='mh-detail']"))         //省份和城市在一起的需要切割
			pandcs := strings.TrimSpace(pandc)
			re1, _ := regexp.Compile(`\n`)
			pandcs = re1.ReplaceAllString(pandcs, ":")
			pandcArr := strings.Split(pandcs, ":")
			pandcArrc := []string{}
			for _,v := range pandcArr {
				v = strings.TrimSpace(v)
				v = strings.Replace(v," ","",-1)
				v = strings.Replace(v,"	","",-1)
				v = strings.Replace(v,"\n","",-1)
				if v != "" {
					fmt.Println(v)
					pandcArrc = append(pandcArrc,v)
				}
			}
			//fmt.Println(pandcArrc)
			if len(pandcArrc) > 3 {
				pr.Location.Province = strings.TrimSpace(pandcArrc[1]) //省份
				pr.Location.City = strings.TrimSpace(pandcArrc[2])     //城市
				pr.Sp = strings.TrimSpace(pandcArrc[3])
			}
			return pr, err
		}

		//模式2
		if pmodemark2 > 0 {
			pr.From = htmlquery.InnerText(htmlquery.FindOne(node, "//*[@class='mohe-tips mh-ws-hy']//a")) //来源渠道
			pandc := htmlquery.InnerText(htmlquery.FindOne(node, "//*[@class='mohe-mobileInfoContent']//span[2]"))         //省份和城市在一起的需要切割
			pandcs := strings.TrimSpace(pandc)
			re1, _ := regexp.Compile(`\n`)
			pandcs = re1.ReplaceAllString(pandcs, ":")
			pandcArr := strings.Split(pandcs, ":")
			if len(pandcArr) > 1 {
				pr.Location.Province = strings.TrimSpace(pandcArr[0]) //省份
				pr.Location.City = strings.TrimSpace(pandcArr[1])     //城市
			}
			if len(pandcArr) > 2 {
				pr.Sp = strings.TrimSpace(pandcArr[2]) //运营商
			}

			pr.Tag.TagName = strings.TrimSpace(htmlquery.InnerText(htmlquery.FindOne(node, "//*[@class='mohe-ph-mark']"))) //标记
			tagCnts := htmlquery.InnerText(htmlquery.FindOne(node, "//*[@class='mohe-tips mh-ws-hy']/span[2]/b"))
			if tagCnts != "" {
				tagCnt, err1 := strconv.Atoi(tagCnts)
				if err1 != nil {
					tagCnt = 0
				}
				pr.Tag.TagCnt = tagCnt //标记人数 位

			}
			return pr, err
		}

		//模式3
		if pmodemark3 > 0 {
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
			return pr, err
		}

	}

	return pr, err
}
