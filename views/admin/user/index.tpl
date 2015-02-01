<div class="row">
{{template "../layout/left_menu.tpl" .}}

<div class="col-sm-9 col-md-10 main">
  <h1 class="page-header">{{i18n .Lang "user.new_users"}}</h1>

  <div class="row placeholders">
    <div class="col-sm-10 col-md-7 main">
      <div class="col-xs-4 col-sm-2 placeholder">
        <img data-src="holder.js/200x200/auto/sky" class="img-responsive img-circle" alt="200x200" src="data:image/png;base64,{{.iamgeSource}}" data-holder-rendered="true">
        <span class="text-muted">舒志平{{.count}}</span>
      </div>
      <div class="col-xs-4 col-sm-2 placeholder">
        <img data-src="holder.js/200x200/auto/vine" class="img-responsive img-circle" alt="200x200" src="data:image/png;base64,{{.iamgeSource2}}" data-holder-rendered="true">
        <span class="text-muted">Something</span>
      </div>
      <div class="col-xs-4 col-sm-2 placeholder">
        <img data-src="holder.js/200x200/auto/sky" class="img-responsive img-circle" alt="200x200" src="data:image/png;base64,{{.iamgeSource3}}" data-holder-rendered="true">
        <span class="text-muted">Something</span>
      </div>
      <div class="col-xs-4 col-sm-2 placeholder">
        <img data-src="holder.js/200x200/auto/vine" class="img-responsive img-circle" alt="200x200" src="data:image/png;base64,{{.iamgeSource4}}" data-holder-rendered="true">
        <span class="text-muted">Something</span>
      </div>
      <div class="col-xs-4 col-sm-2 placeholder">
        <img data-src="holder.js/200x200/auto/vine" class="img-responsive img-circle" alt="200x200" src="data:image/png;base64,{{.iamgeSource5}}" data-holder-rendered="true">
        <span class="text-muted">Something</span>
      </div>
      <div class="col-xs-4 col-sm-2 placeholder">
        <img data-src="holder.js/200x200/auto/vine" class="img-responsive img-circle" alt="200x200" src="data:image/png;base64,{{.iamgeSource6}}" data-holder-rendered="true">
        <span class="text-muted">Something</span>
      </div>
    </div>
  </div>

  <h2 class="sub-header">{{i18n .Lang "user.user_list"}}</h2>
  <div class="table-responsive">
    <table class="table table-striped">
      <thead>
        <tr>
          <th>会员ID</th>
          <th>会员名</th>
          <th>会员邮箱</th>
          <th>上次登录日期</th>
          <th>下单次数</th>
          <th>下单金额</th>
          <th>管理员</th>
          <th>操作</th>
        </tr>
      </thead>
      <tbody>
        {{range $user := .rows}}
        <tr>
          <td>{{$user.Id}}</td>
          <td>{{$user.Nickname}}</td>
          <td>{{$user.Email}}</td>
          <td>{{$user.Lastlogin_data}}</td>
          <td>0.00</td>
          <td>0.00</td>
          <td>{{if eq $user.Isadmin 1}}是{{else}}否{{end}}</td>
          <td>
            <a href="{{urlfor "admin/UserController.Update" "id" $user.Id_str}}">修改</a>
            <a href="{{urlfor "admin/UserController.Delete" "id" $user.Id_str}}">删除</a>
          </td>
        </tr>
        {{end}}
      </tbody>
    </table>
    {{template "../../layout/pager.tpl" .}}
  </div>
</div>
</div>