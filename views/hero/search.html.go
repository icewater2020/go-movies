// Code generated by hero.
// source: E:\www\project\dexter\open\go-movies\views\hero\search.html
// DO NOT EDIT!
package template

import (
	"bytes"
	"go_movies/services"
	"time"

	"github.com/shiyanhui/hero"
)

func Search(show map[string]interface{}, buffer *bytes.Buffer) {
	buffer.WriteString(`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head profile="http://gmpg.org/xfn/11">
    <meta charset="UTF-8">
    <meta http-equiv="Content-Type" content="text/html">
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <title>GoMovies-电影</title>
    <meta name="keywords" content="GoMovies-电影">
    <meta name="description" content="GoMovies-电影">
    <meta property="wb:webmaster" content="bec25808">
    <meta name="referrer" content="never">
    <link rel="canonical" href="/static/css/css">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, user-scalable=0, minimum-scale=1.0, maximum-scale=1.0">
    <link rel="stylesheet" type="text/css" href="/static/css/kube.css">
    <link rel="stylesheet" type="text/css" href="/static/css/reset.css">
    <link rel="stylesheet" type="text/css" href="/static/css/style.css">
    <link rel="stylesheet" type="text/css" href="//cdn.bootcss.com/font-awesome/4.7.0/css/font-awesome.min.css">
    <link rel="shortcut icon" href="/static/image/favicon.ico" type="image/x-icon" />
    <script>var killIE6ImgUrl = "/skin/66scc/images";</script>
    <!--<link rel="stylesheet" type="text/css" href="/static/css/lightbox.css" />-->
    <!--<script type='text/javascript' src='/static/js/lightbox.min.js'></script>-->
    <!--[if lt IE 9]>
    <script src="/static/js/html5.js"></script>
    <![endif]-->
</head>
`)
	buffer.WriteString(`

<body class="custom-background">

`)
	buffer.WriteString(`
    <div id="head" class="row">
        <div class="container row">
            <div class="row">
                <div id="topbar">
                    <ul id="toolbar" class="menu">

                    </ul>
                </div>
                <div id="rss">
                    <ul>
                        <li><a href="javascript:;"  title="最新更新文字版">不要点击</a>  </li>
                    </ul>
                </div>
            </div>
        </div>
        <div class="clear"></div>
        <div class="container">
            <div id="blogname" class="third"> <a href="/" title="6v电影-新版">
                    <h1>
                        go-movies      </h1>
                    <img src="/static/image/logo_.png" alt="6v电影-新版"></a> </div>

        </div>
        <div class="clear"></div>
    </div>`)
	buffer.WriteString(`<div class="mainmenus container" id="nav_b">
        <div class="mainmenu">
            <div class="topnav">
                <ul id="menus">

                    <li class="menu-item
                    `)
	if show["nav_link"] == "/" {
		hero.EscapeHTML("current_page_item", buffer)
	}
	buffer.WriteString(` ">
                    <a href="/">首页</a>
                    </li>


                    `)
	for _, category := range show["categories"].([]map[string]interface{}) {
		buffer.WriteString(`
                        <li class="menu-item
                        `)
		if category["link"].(string) == show["nav_link"] {
			hero.EscapeHTML("current_page_item", buffer)
		}
		buffer.WriteString(` ">
                    <a href="/?cate=`)
		hero.EscapeHTML(category["link"].(string), buffer)
		buffer.WriteString(` ">`)
		hero.EscapeHTML(category["name"].(string), buffer)
		buffer.WriteString(`</a>
                    </li>
                    `)
	}
	buffer.WriteString(`

                    <li class="menu-item
                    `)
	if show["nav_link"] == "/about" {
		hero.EscapeHTML("current_page_item", buffer)
	}
	buffer.WriteString(` ">
                    <a href="/about"> 关于 </a>
                    </li>

                </ul>
                <div id="select_menu">
                    <select onChange="document.location.href=this.options[this.selectedIndex].value;" id="select-menu-nav">
                        <option value="#">导航菜单</option>
                        `)
	for _, category := range show["categories"].([]map[string]interface{}) {
		buffer.WriteString(`
                        <option   value="/?cate=`)
		hero.EscapeHTML(category["link"].(string), buffer)
		buffer.WriteString(` ">`)
		hero.EscapeHTML(category["name"].(string), buffer)
		buffer.WriteString(`</option>
                        `)
	}
	buffer.WriteString(`

                        <option   value="about"> 关于 </option>
                    </select>
                </div>
            </div>
        </div>
        <div class="clear"></div>
    </div>


<div class="search_m box row">
    <div class="search_site" style="width: 100%;">
        <form method="get" name="searchform" id="searchform" action="/search">
            <input type="submit" value="" id="searchsubmit" class="button">
            <label><span>请输入搜索内容，仅支持电影名称</span>
                <input type="text" class="search-s" name="q" required>
            </label>
        </form>
    </div>
</div>
`)
	buffer.WriteString(`


<div class="container">

    <div class="subsidiary row box">
        <div class="bulletin fourfifth"> <span class="sixth">站点公告：</span>
            <marquee class="fivesixth" direction="left" onmouseout="start();" onmouseover="stop();" scrollamount="2" scrolldelay="15;">
                <script src="/static/js/thea2.js"></script>
            </marquee>
        </div>
        <div class="bdshare_small fifth">
        </div>
    </div>

    `)
	buffer.WriteString(`
    <div id="sidebar">


        <div class="search box row">
            <div class="search_site">
                <form method="get" name="searchform" id="searchform" action="/search">
                    <input type="submit" value="" id="searchsubmit" class="button">
                    <label><span>请输入搜索内容，仅支持电影名称</span>
                        <input type="text" class="search-s" name="q" required>
                    </label>
                </form>
            </div>
        </div>

        <div class="widget box row">
            <div id="tab-title">
                <div class="tab">
                    <ul id="tabnav">
                        <li class="selected">最新电影</li>
                        <li class="">最新剧集</li>
                        <li class="">站长推荐</li>
                    </ul>
                </div>
                <div class="clear"></div>
            </div>
            <div id="tab-content">
                <ul class="">
                    `)
	for _, new_film := range show["new_film"].([]services.MovieListStruct) {
		buffer.WriteString(`
                    <li>
                        <a href='/movie?link=`)
		hero.EscapeHTML(new_film.Link, buffer)
		buffer.WriteString(`'>
                            `)
		hero.EscapeHTML(new_film.Name, buffer)
		buffer.WriteString(`
                        </a>
                    </li>
                    `)
	}
	buffer.WriteString(`

                </ul>
                <ul class="hide">
                    `)
	for _, new_tv := range show["new_tv"].([]services.MovieListStruct) {
		buffer.WriteString(`
                    <li>
                        <a href='/movie?link=`)
		hero.EscapeHTML(new_tv.Link, buffer)
		buffer.WriteString(`'>
                            `)
		hero.EscapeHTML(new_tv.Name, buffer)
		buffer.WriteString(`
                        </a>
                    </li>
                    `)
	}
	buffer.WriteString(`
                </ul>
                <ul class="hide">
                    <li>
                        <a href='/movie?link=/?m=vod-detail-id-26855.html'>
                            表姐，你好嘢！
                        </a>
                    </li>

                    <li>
                        <a href='/movie?link=/?m=vod-detail-id-24088.html'>
                            维多利亚一号
                        </a>
                    </li>

                    <li>
                        <a href='/movie?link=/?m=vod-detail-id-14876.html'>
                            爱人同志
                        </a>
                    </li>

                    <li>
                        <a href='/movie?link=/?m=vod-detail-id-26856.html'>
                            表姐，你好嘢！2
                        </a>
                    </li>

                    <li>
                        <a href='/movie?link=/?m=vod-detail-id-26857.html'>
                            表姐，你好嘢！3之大人驾到
                        </a>
                    </li>

                    <li>
                        <a href='/movie?link=/?m=vod-detail-id-26858.html'>
                            表姐，你好嘢！4之情不自禁
                        </a>
                    </li>

                    <li>
                        <a href='/movie?link=/?m=vod-detail-id-14334'>
                            绝命毒师第一季完结
                        </a>
                    </li>
                    <li>
                        <a href='/movie?link=/?m=vod-detail-id-14335.html'>
                            绝命毒师第二季完结
                        </a>
                    </li>
                    <li>
                        <a href='/movie?link=/?m=vod-detail-id-14336.html'>
                            绝命毒师第三季完结
                        </a>
                    </li>
                    <li>
                        <a href='/movie?link=/?m=vod-detail-id-14338.html'>
                            绝命毒师第四季完结
                        </a>
                    </li>
                    <li>
                        <a href='/movie?link=/?m=vod-detail-id-14339.html'>
                            绝命毒师第五季完结
                        </a>
                    </li>

                </ul>
            </div>
        </div>


    </div>
`)
	buffer.WriteString(`
    <div class="mainleft">
        <ul id="post_container" class="masonry clearfix">

            `)
	for _, movieLists := range show["movieLists"].([]services.MovieListStruct) {
		buffer.WriteString(`

            <li class="post box row fixed-hight">
                <div class="post_hover">
                    <div class="thumbnail">
                        <a href="/movie?link=`)
		hero.EscapeHTML(movieLists.Link, buffer)
		buffer.WriteString(`" class="zoom" rel="bookmark" title="`)
		hero.EscapeHTML(movieLists.Name, buffer)
		buffer.WriteString(`">
                            <img
                                src="/static/image/golang.png"
                                data-url="`)
		hero.EscapeHTML(movieLists.Cover, buffer)
		buffer.WriteString(`"
                                onerror="this.src='/static/image/golang.png'"
                                alt="`)
		hero.EscapeHTML(movieLists.Name, buffer)
		buffer.WriteString(`" >
                        </a>
                    </div>
                    <div class="article">
                        <h2><a href="/movie?link=`)
		hero.EscapeHTML(movieLists.Link, buffer)
		buffer.WriteString(`" rel="bookmark" title="`)
		hero.EscapeHTML(movieLists.Name, buffer)
		buffer.WriteString(`">`)
		hero.EscapeHTML(movieLists.Name, buffer)
		buffer.WriteString(`</a></h2>
                        <div class="entry_post">
                            <p>

                               主演： `)
		hero.EscapeHTML(movieLists.Starring, buffer)
		buffer.WriteString(`                            </p>
                        </div>

                        <div class="arrow-catpanel-top">&nbsp;</div>
                    </div>
                    <div class="info">
                        <span class="info_date info_ico">`)
		hero.EscapeHTML(movieLists.UpdatedAt, buffer)
		buffer.WriteString(`</span>

                        <span class="info_category info_ico">`)
		hero.EscapeHTML(movieLists.Category, buffer)
		buffer.WriteString(`</span>
                    </div>
                </div>
            </li>

            `)
	}
	buffer.WriteString(`

        </ul>
        <div class="clear"></div>
    </div>


`)
	buffer.WriteString(`

	</div>

<div class="clear"></div>

</body>


`)
	buffer.WriteString(`
    <div class="clear"></div>
    <div id="footer">
        <div class="footnav container">
            <ul id="footnav" class="menu">
            </ul>
        </div>
        <div class="footnav container">
            <ul id="friendlink" class="menu">
            </ul>
        </div>
        <div class="copyright">
            <p> Copyright &copy; 2019- `)
	hero.EscapeHTML(time.Now().Format("2006"), buffer)
	buffer.WriteString(`
                <a href="">
                    <strong> GoMovies </strong>
                </a>
            </p>

            <p> 本站所有视频均由程序自动采集而来，版权归原创者所有，如果侵犯了你的权益，请通知我删除侵权内容，谢谢合作。 </p>

            <p class="author"> power by
                <a href="https://hzz.cool" target="_blank" rel="external">
                    hzz.cool
                </a>
            </p>
        </div>
    </div>

    <!--gototop-->
    <div id="tbox"> <a id="gotop" href="javascript:void(0)"></a> </div>


    <script type='text/javascript' src='/static/js/jquery.min-3.8.1.js'></script>
    <script type='text/javascript' src='/static/js/lets-kill-ie6-3.8.1.js'></script>
    <script src="/static/org/layer-v3.1.1/layer/layer.js"></script>

    <script type="text/javascript" src="/static/js/jquery.masonry.js"></script>
    <script type="text/javascript" src="/static/js/loostrive.js"></script>

    <script language="javascript" type="text/javascript">
        (function() {
            var oDiv = document.getElementById("nav_b");
            var H = 0, iE6;
            var Y = oDiv;
            while (Y) {
                H += Y.offsetTop;
                Y = Y.offsetParent;
            };
            iE6 = window.ActiveXObject && !window.XMLHttpRequest;
            if (!iE6) {
                window.onscroll = function() {
                    var s = document.body.scrollTop || document.documentElement.scrollTop;
                    if (s > H) {
                        oDiv.className = "mainmenus container nav_b";
                        if (iE6) {
                            oDiv.style.top = (s - H) + "px";
                        }
                    } else {
                        oDiv.className = "mainmenus container";
                    }
                }
            }
        })();
    </script>


    <script>
        function replaceImg(){
            $("img").each(function () {
                let realImgUrl = $(this).attr("data-url");
                if ( realImgUrl !== "" )
                {
                    $(this).attr("src",$(this).attr("data-url"))
                }
            });
        }
        setTimeout(replaceImg, 1000);
    </script>

    <style>
        .search_m{display: none;}
        @media only screen and (max-width: 640px){
            .search_m{max-width: 360px !important;margin: 0 auto;margin-bottom: 10px;display: block;}
        }
    </style>
`)
	buffer.WriteString(`



</html>
`)

}
