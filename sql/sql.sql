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
)
    comment '用户表';

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
)
    comment '系统配置';

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
)
    comment 'ai任务表';

create index core_task_author_id_index
    on core_task (author_id);

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
)
    comment 'ai任务详情';

create index core_task_detail_task_id_index
    on core_task_detail (task_id);
