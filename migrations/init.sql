CREATE DATABASE IF NOT EXISTS auth_api;

create table users
(
    id              int auto_increment,
    email           varchar(256)                          not null,
    canonical_email varchar(256)                          not null,
    status          smallint    default 0                 not null,
    password        varchar(1024)                         null,
    salt            varchar(32) default ''                not null,
    created_at      datetime    default CURRENT_TIMESTAMP not null,
    updated_at      datetime    default NULL              null on update CURRENT_TIMESTAMP,
    constraint users_pk
        primary key (id),
    constraint users_email_key
        unique (canonical_email)
);

create index users_status_index
    on users (status);
