create database if not exists sms;

use sms;

create table if not exists recv_log (
  id bigint not null auto_increment,
  mdn bigint not null,
  receive_mdn bigint not null,
  external_id varchar(36) character set ascii not null,
  message varchar(1600) character set utf8mb4 collate utf8mb4_unicode_ci not null,
  from_city varchar(36) character set ascii null,
  from_state char(2) character set ascii null,
  from_zip varchar(10) character set ascii null,
  from_country char(2) character set ascii null,
  created datetime not null default current_timestamp,
  primary key (id),
  constraint uc_external_id unique (external_id),
  index mdn (mdn),
  index created_mdn_idx (mdn,created desc),
  index (mdn,receive_mdn,created desc)
);