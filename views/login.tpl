<!DOCTYPE HTML>
<html lang="zh-cn" style="background: url('http://s.willerce.com/a4.jpg') repeat center 0 fixed;">
<head>
    <meta charset="UTF-8">
    <title>餐库 › 请先登录</title>
    <link href="/css/style.css" rel="stylesheet" type="text/css"/>
    <script src="/js/jquery.js" type="text/javascript"></script>
</head>
<body>
    <div class="auth_form">
        <form method="post" class="session_form">
            <h1>餐库 › 登录</h1>
            <div class="item">
                <input type="text" id="name" name="account" class="input" placeholder="邮箱" title="你的帐号邮箱" required=""/>
            </div>
            <div class="item">
                <input type="password" id="password" name="password" class="input" placeholder="密码" title="你的帐号密码" required=""/>
            </div>
            <div class="item clearfix">
                <span><a href="/register" title="还没有餐库帐号？"> › 注册</a></span>&nbsp;&nbsp;
                <span><a href="/forget" title="忘记密码？"> › 忘记密码</a></span>
                <button type="submit" id="sub" class="btn">登录</button>
            </div>
        </form>
    </div>
</body>
</html>