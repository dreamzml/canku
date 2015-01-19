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
        <h2>Heading</h2>
        <p>Donec id elit non mi porta gravida at eget metus. Fusce dapibus, tellus ac cursus commodo, tortor mauris condimentum nibh, ut fermentum massa justo sit amet risus. Etiam porta sem malesuada magna mollis euismod. Donec sed odio dui. </p>
        <p><a class="btn btn-default" href="#" role="button">View details &raquo;</a></p>
      </div><!--/.col-xs-6.col-lg-4-->
      <div class="col-xs-6 col-lg-4">
        <h2>Heading</h2>
        <p>Donec id elit non mi porta gravida at eget metus. Fusce dapibus, tellus ac cursus commodo, tortor mauris condimentum nibh, ut fermentum massa justo sit amet risus. Etiam porta sem malesuada magna mollis euismod. Donec sed odio dui. </p>
        <p><a class="btn btn-default" href="#" role="button">View details &raquo;</a></p>
      </div><!--/.col-xs-6.col-lg-4-->
      <div class="col-xs-6 col-lg-4">
        <h2>Heading</h2>
        <p>Donec id elit non mi porta gravida at eget metus. Fusce dapibus, tellus ac cursus commodo, tortor mauris condimentum nibh, ut fermentum massa justo sit amet risus. Etiam porta sem malesuada magna mollis euismod. Donec sed odio dui. </p>
        <p><a class="btn btn-default" href="#" role="button">View details &raquo;</a></p>
      </div><!--/.col-xs-6.col-lg-4-->
      <div class="col-xs-6 col-lg-4">
        <h2>Heading</h2>
        <p>Donec id elit non mi porta gravida at eget metus. Fusce dapibus, tellus ac cursus commodo, tortor mauris condimentum nibh, ut fermentum massa justo sit amet risus. Etiam porta sem malesuada magna mollis euismod. Donec sed odio dui. </p>
        <p><a class="btn btn-default" href="#" role="button">View details &raquo;</a></p>
      </div><!--/.col-xs-6.col-lg-4-->
      <div class="col-xs-6 col-lg-4">
        <h2>Heading</h2>
        <p>Donec id elit non mi porta gravida at eget metus. Fusce dapibus, tellus ac cursus commodo, tortor mauris condimentum nibh, ut fermentum massa justo sit amet risus. Etiam porta sem malesuada magna mollis euismod. Donec sed odio dui. </p>
        <p><a class="btn btn-default" href="#" role="button">View details &raquo;</a></p>
      </div><!--/.col-xs-6.col-lg-4-->
      <div class="col-xs-6 col-lg-4">
        <h2>Heading</h2>
        <p>Donec id elit non mi porta gravida at eget metus. Fusce dapibus, tellus ac cursus commodo, tortor mauris condimentum nibh, ut fermentum massa justo sit amet risus. Etiam porta sem malesuada magna mollis euismod. Donec sed odio dui. </p>
        <p><a class="btn btn-default" href="#" role="button">View details &raquo;</a></p>
      </div><!--/.col-xs-6.col-lg-4-->
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