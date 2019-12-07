create table if not exists recv_log (
  id serial primary key,
  mdn bigint not null,
  receive_mdn bigint not null,
  external_id varchar(36) not null,
  message varchar(1600) not null,
  from_city varchar(36) null,
  from_state char(2) null,
  from_zip varchar(10) null,
  from_country char(2) null,
  created timestamp not null default current_timestamp,
  constraint uc_external_id unique (external_id)
);