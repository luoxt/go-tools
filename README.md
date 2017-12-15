# go-tools
golang工具包

## packfile-for-mac
### 执行方法：
  $ ./packfile-for-mac <br>
  请输入根路径:<br>
  /Users/luoxiantao/WebServer/zy.ttyun.com/<br>
  请输入文件列表和'end'并按回车结束输入：<br>
  addons/cy163_customerservice/ImageCrop.class.php<br>
  addons/cy163_customerservice/emoji/emoji.css<br>
  addons/cy163_customerservice/emoji/emoji.php<br>
  addons/cy163_customerservice/emoji/emoji.png<br>
  addons/cy163_customerservice/icon-custom.jpg<br>
  end<br>

### 输出状态:
  打包文件列表为：<br>
  【序号】 1<br>
  源文件： addons/cy163_customerservice/ImageCrop.class.php<br>
  新文件： /Users/luoxt/Workspace/golang/src/go-tools/bin/zh_201712152153/addons/cy163_customerservice/ImageCrop.class.php<br>
  成功：√ 大小： 7291<br>
  【序号】 2<br>
  源文件： addons/cy163_customerservice/emoji/emoji.css<br>
  新文件： /Users/luoxt/Workspace/golang/src/go-tools/bin/zh_201712152153/addons/cy163_customerservice/emoji/emoji.css<br>
  成功：√ 大小： 78391<br>
  【序号】 3<br>
  源文件： addons/cy163_customerservice/emoji/emoji.php<br>
  新文件： /Users/luoxt/Workspace/golang/src/go-tools/bin/zh_201712152153/addons/cy163_customerservice/emoji/emoji.php<br>
  成功：√ 大小： 516141<br>
  .....<br>

### 最终结果：
  $ ls<br>
  packfile-for-mac	zh_201712152153<br>



