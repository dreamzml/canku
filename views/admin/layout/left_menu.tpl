<div class="col-sm-3 col-md-2 sidebar">
  <ul class="nav nav-sidebar">
    <li><a href="">店家管理</a>
        <ul class="nav">
            <li class="text-right">
                <a  {{if eq .subCur "shop_index"}}class="text-success"{{end}} 
                    href="{{urlfor "admin/ShopController.Index"}}">店家列表</a>
            </li>
            <li class="text-right"><a href="#overview-mobile">等待审核</a>
            </li>
            <li class="text-right"><a href="#overview-type-links">黑名单</a>
            </li>
            <li class="text-right"><a href="#overview-normalize">被学报的</a>
            </li>
            <li class="text-right"><a href="#overview-container">增加店铺</a>
            </li>
        </ul>
    </li>
    <li><a href="">会员管理</a>
        <ul class="nav">
            <li class="text-right">
                <a {{if eq .subCur "user_index"}}class="text-success"{{end}} 
                    href="{{urlfor "admin/UserController.Index"}}">会员列表</a>
            </li>
            <li class="text-right">
                <a {{if eq .subCur "user_create"}}class="text-success"{{end}} 
                    href="{{urlfor "admin/UserController.Create"}}">添加会员</a>
            </li>
        </ul>
    </li>
    <li><a href="">订单管理</a>
        <ul class="nav">
            <li class="text-right"><a class="text-primary" href="#overview-doctype">订单列表</a>
            </li>
            <li class="text-right"><a href="#overview-mobile">创建订单</a>
            </li>
            <li class="text-right"><a href="#overview-mobile">店家订单</a>
            </li>
            <li class="text-right"><a href="#overview-mobile">会员点单</a>
            </li>
        </ul>
    </li>
  </ul>
</div>