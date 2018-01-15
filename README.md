# Cloudgo
## 基本功能
1.在访问“/”时显示“Hello world！”  
2.访问“/user/:name”时显示“Hello name！”其中name是访问时用户自定义的。  

对于web框架，我选择了beego，我认为begoo有很多优点  
比如：  
1.beego比较全面，可以实现的功能很多  
2.beego的教程很丰富，不仅有文字教程，还有视频教程发布在网上，让人学习起来感到轻松一些  
3.bee工具包使得项目的框架创建变得十分快捷  

## 安装beego创建项目：
1、安装beego:  
`go get github.com/astaxie/beego`  
2、安装开发工具bee：
`go get github.com/beego/bee`  
3、创建名为webgo的项目：
`bee new webgo`  
## 使用curl来测试  
curl -v http://loacalhost:9090/  和  curl -v http://loacalhost:9090/user/zhang
结果如下：<br>
![image](https://github.com/Tendernesszh/Cloudgo/blob/master/webgo/curl%E6%B5%8B%E8%AF%95.png)
#### ab测试（指令为：ab -n 1000 -c 100 http://localhost:9090/user/zhang）
其中-n代表执行的请求数量，-c代表并发请求个数。<br>

结果如下：<br>
![image](https://github.com/Tendernesszh/Cloudgo/blob/master/webgo/%E5%8E%8B%E5%8A%9B%E6%B5%8B%E8%AF%95.png)

可以看到在56ms内所有的请求被处理完成，且没有失败请求。
