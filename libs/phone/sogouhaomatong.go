package phone

/*

搜狗搜索

*/
import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func (p *Phone) Query_sogouhaomatong() (pr Phone, err error) {
	qurl := "https://www.sogou.com/web?query=" + p.PhoneNumber
	pr = *p
	pr.Index = 3
	timeout := time.Duration(6 * time.Second) //设置超时6秒
	client := http.Client{

		Timeout: timeout,
	}
	req, err := http.NewRequest("GET", qurl, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.78 Safari/537.36")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return pr, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	//	fmt.Println(doc)
	if err != nil {
		return pr, err
	}

	from1 := doc.Find(".jzVrMsg").Find("a").Text()
	from2 := doc.Find("#sogou_vr_10001001_jzsource_0").Text()
	//来源渠道
	if "" != from1 {
		pr.From = from1
	}
	if "" != from2 {
		pr.From = from2
	}
	bodystr := string(body)

	startstr2 := "tpl491(491, \"10001001\", '', 0,\""
	start2 := strings.Index(bodystr, startstr2)
	endstr2 := "VR TYPE:10001001"
	end2 := strings.Index(bodystr, endstr2)
	if (start2 > 0) && (end2 > 0) {
		//如果找到关键字才继续
		//	fmt.Println(end)
		restr2 := strings.TrimSpace(strings.Replace(bodystr[(start2+31):(end2-78)], p.PhoneNumber, "", -1))
		//fmt.Println(restr2)
		restr2a := strings.Split(restr2, " ")
		if len(restr2a) > 1 {
			pr.Location.Province = restr2a[0] //省份
			pr.Location.City = restr2a[1]     //城市
		}
	}
	startstr := "queryphoneinfo"
	start := strings.Index(bodystr, startstr)
	endstr := "'.replace"
	end := strings.Index(bodystr, endstr)
	if start > 0 && end > 0 {
		//	fmt.Println(start)
		//	fmt.Println(end)
		restr := strings.Split(bodystr[(start+18):(end)], "：")
		if len(restr) > 1 {
			pr.Tag.TagName = restr[1] //标记
			tagCnt, _ := strconv.Atoi(restr[2])
			pr.Tag.TagCnt = tagCnt //标记人数 位
		}
	}

	return pr, err
}
