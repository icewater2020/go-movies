// Code generated by hero.
// source: E:\www\project\dexter\open\go-movies\views\hero\m3u8.html
// DO NOT EDIT!
package template

import (
	"bytes"

	"github.com/shiyanhui/hero"
)

func M3u8(show map[string]interface{}, buffer *bytes.Buffer) {
	buffer.WriteString(`<!DOCTYPE html> <!-- 用于声明、告诉浏览器当前是一个 HTML 文档 -->
`)
	buffer.WriteString(`
<html lang="zh-cn"> <!-- HTML 文档开始 | lang="zh-cn" 中文声明 -->
<head>
    <meta charset="UTF-8">
    <title>play</title>
    <meta name="referrer" content="never">

    <link href="//cdn.bootcss.com/video.js/7.6.0/alt/video-js-cdn.min.css" rel="stylesheet">
</head>
<body>

<video id=example-video  class="video-js vjs-default-skin" controls preload="none">
    <source
            src="`)
	hero.EscapeHTML(show["play_url"].(string), buffer)
	buffer.WriteString(`"
            type="application/x-mpegURL">
</video>


</body>

<script src="//cdn.bootcss.com/video.js/7.6.0/alt/video.core.min.js"></script>
<script src="//cdn.bootcss.com/videojs-contrib-hls/5.15.0/videojs-contrib-hls.min.js"></script>

<script>
    var player = videojs('example-video', { fluid: true });
    player.play();
</script>
</html>
`)

}
