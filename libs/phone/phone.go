package phone

import (
	"sort"
	"strings"
)

type Location struct {
	Province string `json:"province"` //省份
	City     string `json:"city"`     //城市
}

type Tag struct {
	TagName string `json:"tag_name"` //标记名称
	TagCnt  int    `json:"tag_cnt"`  //标记人数
}

type Phone struct {
	Index       int      `json:"index"`        //索引
	PhoneNumber string   `json:"phone_number"` //号码
	Sp          string   `json:"sp"`           //运营商
	Card        string   `json:"card"`         //卡类型
	From        string   `json:"from"`         //来源渠道
	Location    Location `json:"location"`     //归属地
	Tag         Tag      `json:"tag"`          //标记
}

type PhoneList []*Phone

func (p PhoneList) Len() int {
	return len(p)
}

func (p PhoneList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PhoneList) Less(i, j int) bool {
	return p[i].Index < p[j].Index
}

//新建查询号码
func New(phoneNumber string) (p *Phone) {
	p = new(Phone)
	p.PhoneNumber = strings.TrimSpace(phoneNumber)
	return
}

//查询号码 所有渠道的信息
func (p *Phone) QueryAll() (pr PhoneList, err error) {
	var pl PhoneList
	pchan := make(chan *Phone, 3)
	go func() {
		result360, err := p.Query_360shoujiweishi()
		if err != nil {
			pchan <- nil
		} else {
			pchan <- &result360
		}
	}()
	go func() {
		resultbaidu, err := p.Query_baidushoujiweishi()
		if err != nil {
			pchan <- nil
		} else {
			pchan <- &resultbaidu
		}
	}()
	go func() {
		resultsogou, err := p.Query_sogouhaomatong()
		if err != nil {
			//fmt.Println(resultsogou)
			pchan <- nil
		} else {
			pchan <- &resultsogou
		}
	}()
	for i := 0; i < 3; i++ {
		pone := <-pchan
		if pone != nil {
			pl = append(pl, pone)
		}
		//fmt.Println(pone)
	}
	sort.Sort(pl)
	//fmt.Println(PhoneList)
	pr = pl
	return
}

//查询号码 指定单独一个渠道的信息
func (p *Phone) QueryOne(from string) (pl PhoneList, err error) {
	var pr Phone
	switch from {
	case "360shoujiweishi":
		pr, _ = p.Query_360shoujiweishi()
	case "baidushoujiweishi":
		pr, _ = p.Query_baidushoujiweishi()
	case "sogouhaomatong":
		pr, _ = p.Query_sogouhaomatong()
	}
	pl = append(pl, &pr)
	return
}
