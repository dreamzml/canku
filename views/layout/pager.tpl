{{if gt .pager.PageNums 1}}
<nav class="text-center">
<ul class="pagination">
    {{if .pager.HasPrev}}
        <li><a href="{{.pager.PageLinkFirst}}">{{i18n .Lang "pager.first_page"}}</a></li>
        <li><a href="{{.pager.PageLinkPrev}}">&lt;</a></li>
    {{else}}
        <li class="disabled"><a>{{i18n .Lang "pager.first_page"}}</a></li>
        <li class="disabled"><a>&lt;</a></li>
    {{end}}
    {{range $index, $page := .pager.Pages}}
        <li{{if $.pager.IsActive .}} class="active"{{end}}>
            <a href="{{$.pager.PageLink $page}}">{{$page}}</a>
        </li>
    {{end}}
    {{if .pager.HasNext}}
        <li><a href="{{.pager.PageLinkNext}}">&gt;</a></li>
        <li><a href="{{.pager.PageLinkLast}}">{{i18n .Lang "pager.last_page"}}</a></li>
    {{else}}
        <li class="disabled"><a>&gt;</a></li>
        <li class="disabled"><a>{{i18n .Lang "pager.last_page"}}</a></li>
    {{end}}
</ul>
</nav>
{{end}}