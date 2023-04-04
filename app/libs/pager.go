package libs

import (
	"bytes"
	"fmt"
	"math"
	"strings"
)

type Pager struct {
	Page     int
	Totalnum int
	Pagesize int
	urlpath  string
	urlquery string
	nopath   bool
}

func NewPager(page, totalnum, pagesize int, url string, nopath ...bool) *Pager {
	p := new(Pager)
	p.Page = page
	p.Totalnum = totalnum

	arr := strings.Split(url, "?")
	if len(arr) > 1 {
		p.urlquery = "?" + arr[1]
	} else {
		p.urlquery = ""
	}

	if len(nopath) > 0 {
		p.nopath = nopath[0]
	} else {
		p.nopath = false
	}

	return p
}

func (p *Pager) url(page int) string {
	if p.nopath { // 不使用目录形式
		if p.urlquery != "" {
			return fmt.Sprintf("%s%s&page=%d", p.urlpath, p.urlquery, page)
		} else {
			return fmt.Sprintf("%s?page=%d", p.urlpath, page)
		}
	} else {
		return fmt.Sprintf("%s/page/%d%s", p.urlpath, page, p.urlquery)
	}
}

func (p *Pager) ToString() string {
	if p.Totalnum <= p.Pagesize {
		return ""
	}

	var buf bytes.Buffer
	var from, to, linknum, offset, totalpage int

	offset = 5
	linknum = 10
	totalpage = int(math.Ceil(float64(p.Totalnum) / float64(p.Pagesize)))

	if totalpage < linknum {
		from = 1
		to = totalpage
	} else {
		from = p.Page - offset
		to = from * linknum
		if from < 1 {
			from = 1
			to = from + linknum - 1
		} else if to > totalpage {
			to = totalpage
			from = totalpage - linknum + 1
		}
	}

	buf.WriteString("<ul class=\"pagination\">")
	if p.Page > 1 {
		buf.WriteString(fmt.Sprintf("<li><a href=\"%s\">&laquo;</a></li>", p.url(p.Page-1)))
	} else {
		buf.WriteString("<li class=\"disabled\"><span>&laquo;</span></li>")
	}

	if p.Page > linknum {
		buf.WriteString(fmt.Sprintf("<li><a href=\"%s\">1...</a></li>", p.url(1)))
	}

	for i := from; i <= to; i++ {
		if i == p.Page {
			buf.WriteString(fmt.Sprintf("<li class=\"active\"><span>%d</span></li>", i))
		} else {
			buf.WriteString(fmt.Sprintf("<li><a href=\"%s\">%d</a></li>", p.url(i), i))
		}
	}

	if totalpage > to {
		buf.WriteString(fmt.Sprintf("<li><a href=\"%s\">...%d</a></li>", p.url(totalpage), totalpage))
	}

	if p.Page < totalpage {
		buf.WriteString(fmt.Sprintf("<li><a href=\"%s\">&raquo;</a></li>", p.url(p.Page+1)))
	} else {
		buf.WriteString(fmt.Sprintf("<li class=\"disabled\"><span>&raquo;</span></li>"))
	}
	buf.WriteString("</ul>")

	return buf.String()
}
