/* 
* @Author: dreamzml
* @Date:   2015-02-01 05:44:56
* @Last Modified by:   zoro
* @Last Modified time: 2015-02-01 05:48:11
*/

package lib

import (
    "math"
    "net/http"
    "net/url"
    "strconv"
)

type Pager struct {
    Request     *http.Request
    PerPageNums int
    MaxPages    int

    nums      int64
    pageRange []int
    pageNums  int
    page      int
}

func (p *Pager) PageNums() int {
    if p.pageNums != 0 {
        return p.pageNums
    }
    pageNums := math.Ceil(float64(p.nums) / float64(p.PerPageNums))
    if p.MaxPages > 0 {
        pageNums = math.Min(pageNums, float64(p.MaxPages))
    }
    p.pageNums = int(pageNums)
    return p.pageNums
}

func (p *Pager) Nums() int64 {
    return p.nums
}

func (p *Pager) SetNums(nums interface{}) {
    p.nums, _ = ToInt64(nums)
}

func (p *Pager) Page() int {
    if p.page != 0 {
        return p.page
    }
    if p.Request.Form == nil {
        p.Request.ParseForm()
    }
    p.page, _ = strconv.Atoi(p.Request.Form.Get("p"))
    if p.page > p.PageNums() {
        p.page = p.PageNums()
    }
    if p.page <= 0 {
        p.page = 1
    }
    return p.page
}

func (p *Pager) Pages() []int {
    if p.pageRange == nil && p.nums > 0 {
        var pages []int
        pageNums := p.PageNums()
        page := p.Page()
        switch {
        case page >= pageNums-4 && pageNums > 9:
            start := pageNums - 9 + 1
            pages = make([]int, 9)
            for i, _ := range pages {
                pages[i] = start + i
            }
        case page >= 5 && pageNums > 9:
            start := page - 5 + 1
            pages = make([]int, int(math.Min(9, float64(page+4+1))))
            for i, _ := range pages {
                pages[i] = start + i
            }
        default:
            pages = make([]int, int(math.Min(9, float64(pageNums))))
            for i, _ := range pages {
                pages[i] = i + 1
            }
        }
        p.pageRange = pages
    }
    return p.pageRange
}

func (p *Pager) PageLink(page int) string {
    link, _ := url.ParseRequestURI(p.Request.RequestURI)
    values := link.Query()
    if page == 1 {
        values.Del("p")
    } else {
        values.Set("p", strconv.Itoa(page))
    }
    link.RawQuery = values.Encode()
    return link.String()
}

func (p *Pager) PageLinkPrev() (link string) {
    if p.HasPrev() {
        link = p.PageLink(p.Page() - 1)
    }
    return
}

func (p *Pager) PageLinkNext() (link string) {
    if p.HasNext() {
        link = p.PageLink(p.Page() + 1)
    }
    return
}

func (p *Pager) PageLinkFirst() (link string) {
    return p.PageLink(1)
}

func (p *Pager) PageLinkLast() (link string) {
    return p.PageLink(p.PageNums())
}

func (p *Pager) HasPrev() bool {
    return p.Page() > 1
}

func (p *Pager) HasNext() bool {
    return p.Page() < p.PageNums()
}

func (p *Pager) IsActive(page int) bool {
    return p.Page() == page
}

func (p *Pager) Offset() int {
    return (p.Page() - 1) * p.PerPageNums
}

func (p *Pager) HasPages() bool {
    return p.PageNums() > 1
}

func NewPager(req *http.Request, per int, nums interface{}) *Pager {
    p := Pager{}
    p.Request = req
    if per <= 0 {
        per = 10
    }
    p.PerPageNums = per
    p.SetNums(nums)
    return &p
}