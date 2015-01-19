<div class="row row-offcanvas row-offcanvas-right">

  <div class="col-xs-12 col-sm-9">
    <p class="pull-right visible-xs">
      <button type="button" class="btn btn-primary btn-xs" data-toggle="offcanvas">Toggle nav</button>
    </p>

  <div class="row">
    <div id="container" class="col-xs-12 col-sm-9" style="height: 400px;width: 100%;"></div>
  </div>

    <div class="row">
      <div class="col-xs-6 col-lg-4">
        <h3>食其家(广中西路店)</h3>
        <p><img src="/static/img/shiqijia.png" class="img-responsive" alt="Responsive image"></p>
        <a class="btn btn-default" href="#" role="button">查看详情 &raquo;</a>
      </div>

      <div class="col-xs-6 col-lg-4">
        <h3>味千拉面(大宁店)</h3>
        <p><img src="/static/img/weiqian.png" class="img-responsive" alt="Responsive image"></p>
        <a class="btn btn-default" href="#" role="button">查看详情 &raquo;</a>
      </div>

      <div class="col-xs-6 col-lg-4">
        <h3>汉堡王(大宁店)</h3>
        <p><img src="/static/img/hanbaowang.png" class="img-responsive" alt="Responsive image"></p>
        <a class="btn btn-default" href="#" role="button">查看详情 &raquo;</a>
      </div>

      <div class="col-xs-6 col-lg-4">
        <h3>麻辣诱惑(大宁店)</h3>
        <p><img src="/static/img/malayouhuo.jpg" class="img-responsive" alt="Responsive image"></p>
        <a class="btn btn-default" href="#" role="button">查看详情 &raquo;</a>
      </div>

      <div class="col-xs-6 col-lg-4">
        <h3>宏港烧腊便当</h3>
        <p><img src="/static/img/gangshaola.png" class="img-responsive" alt="Responsive image"></p>
        <a class="btn btn-default" href="#" role="button">查看详情 &raquo;</a>
      </div>

      <div class="col-xs-6 col-lg-4">
        <h3>众湘汇</h3>
        <p><img src="/static/img/zhongxianghui.png" class="img-responsive" alt="Responsive image"></p>
        <a class="btn btn-default" href="#" role="button">查看详情 &raquo;</a>
      </div>

    </div><!--/row-->
  </div><!--/.col-xs-12.col-sm-9-->

  <div class="col-xs-6 col-sm-3 sidebar-offcanvas" id="sidebar" role="navigation">
    <div class="list-group">
      <a href="#" class="list-group-item active">Link</a>
      <a href="#" class="list-group-item">Link</a>
      <a href="#" class="list-group-item">Link</a>
      <a href="#" class="list-group-item">Link</a>
      <a href="#" class="list-group-item">Link</a>
      <a href="#" class="list-group-item">Link</a>
      <a href="#" class="list-group-item">Link</a>
      <a href="#" class="list-group-item">Link</a>
      <a href="#" class="list-group-item">Link</a>
      <a href="#" class="list-group-item">Link</a>
    </div>
  </div><!--/.sidebar-offcanvas-->
</div><!--/row-->


<script src="/static/js/highchart/highcharts.js"></script>
<script type="text/javascript" src="/static/js/highchart/themes/grid.js"></script>
<script type="text/javascript">
  $(function () {
    $('#container').highcharts({
        chart: {
            type: 'bar'
        },
        title: {
            text: '昨日排行'
        },
        xAxis: {
            categories: ['福荣祥港式烧腊','宏港烧腊便当','万州小厨','味千拉面','汉堡王']
        },
        yAxis: {
            min: 0,
            title: {
                text: 'Total fruit consumption'
            }
        },
        legend: {
           // backgroundColor: '#eeeeee',
            reversed: true
        },
        plotOptions: {
            series: {
                stacking: 'normal'
            }
        },
        series: [{
            name: 'John',
            data: [15, 8, 4, 7, 2]
        }]
    });
  });
</script>