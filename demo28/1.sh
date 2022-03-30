#!/bin/bash

echo "开始ffmpeg批量图片合成mp4";

docker run -v /Users/lvxin/go/src/github.com/lvxin0315/7788demo/demo28/images:/images \
jrottenberg/ffmpeg:latest \
-f image2 -i /images/_%09d.png  -vcodec libx264  /images/final.mp4