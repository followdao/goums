-- create database in pg 11+
-- drop database if exists goums;
create database goums;
-- create schema
drop schema ums cascade;
create schema ums;

-- create extension if not exists "uuid-ossp" with schema queue;
-- apk type
drop table if exists ums.apktype;
create table if not exists ums.apktype
(
id int generated always as identity (cache 100)
  primary key,
apk_name varchar(64),
apk_type varchar(64) not null,
active_status boolean default true not null
);

comment on column ums.apktype.apk_type is 'in android project gradle setting applicationId, it is the apk_type';

insert
into
  ums.apktype
(apk_name,apk_type,active_status)
values
('appleTV','com.apple.liveTV',true),
('appleVOD','com.apple.movieHD',false);

drop table if exists ums.terminal;
create table if not exists ums.terminal
(
id bigint not null
  primary key default queue.id_generator(),
serial_number varchar(64) not null,
active_code varchar(64) not null,
active_status boolean default false not null,
active_date timestamp default now(),
max_active_session bigint default 1,
access_role varchar(32) default 'tvbox'::character varying not null,
service_status smallint default 0,
service_expiration timestamp default (now() + ((31)::double precision * '1 day'::interval))
);

comment on column ums.terminal.service_status is 'service status define:
0 default
1 actived
2 suspend
3 disabled
4 deleted';

insert
into
  ums.terminal
(serial_number,active_code,access_role)
values
('aaaaaaaa','12345678','tvbox'),
('bbbbbbbb','11112222','phone');


