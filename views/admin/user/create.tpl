<div class="row">
{{template "../layout/left_menu.tpl" .}}

  <div class="col-sm-9 col-md-10 main" method="post">
    <h1 class="page-header">{{i18n .Lang "user.user_create"}}</h1>

    <div class="col-md-7  col-md-offset-1">
      <form class="form_ajax_submit" method="post" action="{{urlfor "admin/UserController.CreatePost"}}" redirect="{{urlfor "admin/UserController.Index"}}">
        <div class="form-group">
          <label for="inputEmail3" class="exampleInputFile">邮箱</label>
          <input type="email" name="Email" class="form-control" id="inputEmail3" placeholder="请输入邮件">
        </div>

        <div class="form-group">
          <label for="inputPassword3" class="exampleInputFile">姓名</label>
          <input type="nickname" name="Nickname" class="form-control" id="inputPassword3" placeholder="请输入会员名">
        </div>

        <div class="form-group">
          <label for="inputPassword3" class="exampleInputFile">密码</label>
          <input type="password" name="Password" class="form-control" id="inputPassword3" placeholder="请输入会员密码">
        </div>

        <div class="checkbox">
            <label>
              <input name="Isadmin" type="checkbox" value="1">设为管理员
            </label>
        </div>

        <input type="submit" class="btn btn-primary" value="{{i18n .Lang "public.save"}}">
        <input type="reset" class="btn btn-default" value="{{i18n .Lang "public.reset"}}">

      </form>
    </div>
  </div>
</div>

<script>
  $('form.form_ajax_submit').submit(function(){
    var form = $(this);
    var postUrl = form.attr("action");
    var redirect = form.attr("redirect");
    $.ajax({
        url:postUrl,
        type:'POST',
        async:true,
        dataType:'json',
        data:form.serialize(),
        beforeSend:function(XMLHttpRequest){
          form.find('.form-info').remove();
          form.prepend('<p class="form-info text-info">表单下在提交中....</p>');
        },
        complete:function(XMLHttpRequest){
          if(!form.find('.form-info').hasClass('text-danger'))
            form.find('.form-info').remove();
        },
        success:function(data){
          //成功跳转
          if(data.status == 200){
            window.location.href=redirect;
          }else{
            form.find('.form-info').html(data.info).removeClass('text-info').addClass('text-danger');
          }
        },
        error:function(XMLHttpRequest){
          form.find('.form-info').html("系统错误，请稍后再试！").removeClass('text-info').addClass('text-danger');
        }
    });
    return false;
  });


</script>