# ai-gallery
<hr/>

[ [English](README.md)  | Chinese ]

ai-gallery 是一个基于 Go-Zero + SD Plugin + Vite/React + Ant Design 技术的前后端分离系统，用于统一管理 SD 绘画任务。

- [后端项目](https://github.com/tabelf/ai-gallery) <br/>
- [前端项目](https://github.com/tabelf/ai-gallery-ui) <br/>
- [SD 插件](https://github.com/tabelf/sd-webui-gen2gallery) <br/>

## 功能演示
<hr/>

https://github.com/user-attachments/assets/db6ac661-84ca-47b8-934b-86f1f61a9578

[完整视频演示](https://www.bilibili.com/video/BV1mt8ue6E1Y)

## ✨ 功能
<hr/>

- 管理 SD 文生图、图生图绘画任务记录
- SD 生成图片云端/本地存储
- 多用户提交任务数据汇总
- 用户管理
- 系统设置

## 💼 准备工作

### 环境准备
```text
go sdk                   v1.21.0
node                     v21.6.2
npm                      v9.6.7
react                    v18.2.0
stable diffuison webui   v1.9.3
mysql                    v8.0.28
nginx                    v1.25.3
```

### 获取代码
```bash
# 获取插件代码
git clone https://github.com/tabelf/sd-webui-gen2gallery.git

# 获取后端代码
git clone https://github.com/tabelf/ai-gallery.git

# 获取前端代码
git clone https://github.com/tabelf/ai-gallery-ui.git
```

## 启动说明

### 服务端启动
```bash
# 进入到后端项目目录
cd ./ai-gallery

# 安装 go 项目依赖
go mod tidy

# 修改数据库配置
vi ./service/etc/application-dev.yaml

# root 修改为自己数据库的用户名
# 12345678 为密码
# ai_gallery 为数据库名称

# db:
#  url: root:12345678@(127.0.0.1:3306)/ai_gallery?charset=utf8mb4&parseTime=true&loc=Local&interpolateParams=true
#  driver: mysql
#  max_open_conns: 100
#  max_idle_conns: 30

# 生成表结构
go run cmd/main.go migrate --env dev

# 运行项目
go run cmd/main.go start --env dev
```

### 前端启动
```bash
# 安装依赖，如果执行慢需要配置镜像源
npm install

# 启动服务
vite dev

# 部署成功会显示地址
http://localhost:5173/

# 首先要进行登录
# 默认账号：admin 
# 密码：   1234567

# 到系统设置菜单栏中，配置图片要存储的位置
# 如果存储选择本地，需要配置下面的 nginx
# 如果存储选择腾讯云cos，需要配置存储地址，ID，key，记得打开跨域设置
# 如果存储选择阿里云oos，需要配置存储桶名称，存储地址，ID，key，记得打开跨域设置
```

### 登录账号
```
默认账号：admin 
密码：   1234567
```

### SD 插件
```
# 1. 把插件项目放到 stable diffusion webui 的 extensions 扩展目录下即可

# 2. 正常启动 webui
```

### nginx 配置
```
# 在 nginx.conf 配置文件中添加

location /upload/ {
    alias /Users/stable-diffusion-webui/; # 修改成自己 stable-diffusion-webui 文件的路径亘路径
    autoindex on;    

    add_header 'Access-Control-Allow-Origin' '*';
    add_header 'Access-Control-Allow-Methods' 'GET, POST, PUT, DELETE, HEAD, OPTIONS';
    add_header 'Access-Control-Allow-Headers' 'Content-Type, X-CSRF-Token, Authorization, AccessToken, Token, Cache-Control';
    add_header 'Access-Control-Allow-Credentials' 'true';

    if ($request_method = 'OPTIONS') {
        add_header 'Access-Control-Allow-Origin' '*';
        add_header 'Access-Control-Allow-Methods' 'GET, POST, PUT, DELETE, HEAD, OPTIONS';
        add_header 'Access-Control-Allow-Headers' 'Content-Type, X-CSRF-Token, Authorization, AccessToken, Token, Cache-Control';
        add_header 'Access-Control-Allow-Credentials' 'true';
        add_header 'Access-Control-Max-Age' 1728000;
        add_header 'Content-Type' 'text/plain charset=UTF-8';
        add_header 'Content-Length' 0;
        return 204;
     }
}
```

## 联系

<table>
   <tr>
    <td><img src="./wechat.png" width="180px"></td>
  </tr>
  <tr>
    <td>微信留言</td>
  </tr>
</table>
