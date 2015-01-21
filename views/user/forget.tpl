
<!DOCTYPE HTML>
<html lang="zh-cn" style="background: url('http://s.willerce.com/a4.jpg') repeat center 0 fixed;">
<head>
  <meta charset="UTF-8">
  <title>餐库 › 重置密码</title>
  <link href="/css/style.css" rel="stylesheet" type="text/css"/>
  <script src="/js/jquery.js" type="text/javascript"></script>
</head>
<body>
<div class="auth_form">
  <form class="session_form" method="post" action="/sendmail">
    <h1>餐库 › 重置密码</h1>
    <div class="item">
      <input type="text" id="email" name="email" class="input" placeholder="帐户邮箱" title="你的帐户邮箱" required=""/>
    </div>
    <div class="item cleafix">
      <span><a href="/login" title="已有餐库帐号？"> › 登录</a></span>&nbsp;&nbsp;
      <span><a href="/register" title="注册一个"> › 注册 </a></span>
      <button type="submit" class="btn">重置密码</button>
    </div>
    
  </form>
</div>
</body>
</html>