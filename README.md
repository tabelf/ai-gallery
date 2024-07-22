# ai-gallery
<hr/>

ai-gallery 是一项基于 Go-Zero + SD Plugin + Vite/React + Ant Design 的技术，完成管理 SD 绘画任务的前后端分离系统。

- [后端项目](https://github.com/tabelf/ai-gallery) <br/>
- [前端项目](https://github.com/tabelf/ai-gallery-ui) <br/>
- [SD 插件](https://github.com/tabelf/sd-webui-gen2gallery) <br/>

## 功能演示
<hr/>

[视频演示](https://www.bilibili.com/video/BV1mt8ue6E1Y)

## ✨ 功能
<hr/>

- SD 文生图、图生图绘画任务记录
- SD 文生图、图生图生成图片云端/本地存储
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

### 导入表结构
1. 在 mysql 中创建数据库 ai_gallery
```sql
create schema ai_gallery collate utf8mb4_0900_ai_ci;
```
2. 在 ai_gallery 库中执行表结构
```sql

create table core_task
(
    id              bigint auto_increment comment 'ID'
        primary key,
    create_time     datetime          not null comment '创建时间',
    update_time     datetime          not null comment '更新时间',
    delete_time     datetime          null comment '删除时间',
    category        varchar(20)       null comment '类型: txt2img, img2img',
    prompt          text              null comment '提示词',
    negative_prompt text              null comment '反向提示词',
    weight          float             null comment '长度',
    height          float             null comment '宽度',
    img_size        varchar(20)       null comment '尺寸',
    seed            varchar(20)       null comment '随机数',
    sampler_name    varchar(50)       null comment '采样器',
    steps           int               null comment '采样步数',
    cfg_scale       int               null comment '引导系数',
    batch_size      int               null comment '批次',
    total           int               null comment '总量',
    sd_model_name   varchar(100)      null comment '模型名称',
    sd_model_hash   varchar(50)       null comment '模型hash值',
    sd_vae_name     varchar(100)      null comment 'VAE名称',
    sd_vae_hash     varchar(50)       null comment 'VAE hash值',
    job_timestamp   datetime          null comment '任务时间',
    version         varchar(20)       null comment 'SD 版本号',
    grid_image_url  varchar(255)      null comment '网格图片地址',
    status          varchar(20)       null comment '状态：INIT, PROGRESS, COMPLETE, INTERRUPT',
    count           int     default 0 null comment '实际作品数量',
    has_excellent   tinyint default 0 null comment '是否存在优秀作品',
    ref_images      json              null comment '参考图片',
    store           varchar(20)       null comment '服务存储名称',
    author_id       int               null comment '作者ID',
    author_name     varchar(30)       null comment '作者名称',
    extra           text              null comment '扩展信息'
) comment 'ai任务表';

create index core_task_author_id_index on core_task (author_id);

create table core_task_detail
(
    id            bigint auto_increment comment 'ID'
        primary key,
    create_time   datetime     not null comment '创建时间',
    update_time   datetime     not null comment '更新时间',
    delete_time   datetime     null comment '删除时间',
    task_id       bigint       null comment '任务ID',
    image_url     varchar(255) null comment '图片地址',
    has_excellent tinyint      null comment '是否优秀作品'
) comment 'ai任务详情';

create index core_task_detail_task_id_index on core_task_detail (task_id);

create table core_account
(
    id          bigint auto_increment comment 'ID'
        primary key,
    create_time datetime     not null comment '创建时间',
    update_time datetime     not null comment '更新时间',
    delete_time datetime     null comment '删除时间',
    username    varchar(255) null comment '用户名',
    password    varchar(255) null comment '密码',
    nickname    varchar(30)  null comment '昵称',
    role        varchar(20)  null comment '角色: ADMIN, USER',
    status      int          null comment '状态: 0 禁用, 1 可用',
    constraint core_account_pk
        unique (username)
) comment '用户表';

create table core_setting
(
    id           bigint auto_increment comment 'ID'
        primary key,
    create_time  datetime     not null comment '创建时间',
    update_time  datetime     not null comment '更新时间',
    delete_time  datetime     null comment '删除时间',
    config_key   varchar(30)  null comment '配置名',
    config_value varchar(255) null comment '配置值',
    mark         varchar(255) null comment '说明',
    operate_id   int          null comment '操作人ID',
    operate_name varchar(30)  null comment '操作人名称',
    constraint core_setting_config_key_uindex
        unique (config_key)
) comment '系统配置';
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

# 下面的 root 修改为自己数据库的用户名, 12345678 为密码
# db:
#  url: root:12345678@(127.0.0.1:3306)/ai_gallery?charset=utf8mb4&parseTime=true&loc=Local&interpolateParams=true
#  driver: mysql
#  max_open_conns: 100
#  max_idle_conns: 30

# 编译项目
go build -o ai-gallery ./service/gallery.go

# 运行项目
./ai-gallery
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

### 前端登录
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
    alias /Users/stable-diffusion-webui/; # 修改成自己 stable-diffusion-webui 文件的路径
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
