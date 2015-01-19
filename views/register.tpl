<!DOCTYPE HTML>
<html lang="zh-cn" style="background: url('http://s.willerce.com/a4.jpg') repeat center 0 fixed;">
<head>
    <meta charset="UTF-8">
    <title>餐库 › 注册</title>
    <link href="/css/style.css" rel="stylesheet" type="text/css"/>
    <script src="/js/jquery.js" type="text/javascript"></script>
</head>
<body>
<div class="auth_form">
    <form class="session_form" method="post" action="/register">
        <h1>餐库 › 注册</h1>
        <div class="item">
            <input type="text" id="email" name="email" class="input" placeholder="邮箱" title="你的常用邮箱" required=""/>
        </div>
        <div class="item">
            <input type="text" id="nickname" name="nickname" class="input" placeholder="姓名" title="用中文怎么称呼你" required="" maxlength="4"/>
        </div>
        <div class="item">
            <input type="password" id="password" name="password" class="input" placeholder="密码" title="设置一个密码" required="" maxlength="16"/>
        </div>
        <div class="item cleafix">
            <span><a href="/login" title="已有餐库帐号？"> › 登录</a></span>&nbsp;&nbsp;
            <span><a href="/forget" title="忘记密码？"> › 忘记密码</a></span>
            <button type="submit" class="btn">注册</button>
        </div>
    </form>
</div>
</body>
</html>