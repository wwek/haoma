package phone

import (
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func (p *Phone) Query_baidushoujiweishi() (pr Phone, err error) {
	qurl := "https://www.baidu.com/s?wd=" + p.PhoneNumber
	pr = *p
	pr.Index = 2
	timeout := time.Duration(6 * time.Second) //设置超时6秒
	client := http.Client{

		Timeout: timeout,
	}
	req, err := http.NewRequest("GET", qurl, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.110 Safari/537.36")
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
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	//	fmt.Println(doc)
	if err != nil {
		logs.Warn(err)
		return pr, err
	}

	pr.From = doc.Find(".op_fraudphone_word").Find("a").Text() //来源渠道
	sfcs := strings.Split(doc.Find(".op_fraudphone_addr.c-gap-right-small").Text(), " ")
	if len(sfcs) > 1 {
		pr.Location.Province = sfcs[0] //省份
		pr.Location.City = sfcs[1]     //城市
	}

	pr.Tag.TagName = doc.Find(".op_fraudphone_label").Text() //标记
	tagCnts := doc.Find(".op_fraudphone_word").Text()    //标记人数 位
	tagCntsRe := regexp.MustCompile(`\d+`)
	tagCnt, err := strconv.Atoi(tagCntsRe.FindString(tagCnts))
	pr.Tag.TagCnt = tagCnt

	return pr, err
}
