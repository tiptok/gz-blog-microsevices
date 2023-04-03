create table posts
(
    id             bigint unsigned not null auto_increment primary key,
    uuid           varchar(36)     not null unique,
    user_id        bigint unsigned not null,
    title          varchar(255)    not null,
    content        text            not null,
    comments_count int unsigned    not null default 0,
    created_at     timestamp       not null default current_timestamp,
    updated_at     timestamp       not null default current_timestamp on update current_timestamp,
    deleted_at     timestamp       null,
    version    bigint       default 0,
    index (user_id)
);