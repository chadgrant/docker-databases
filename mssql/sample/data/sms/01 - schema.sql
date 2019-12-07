create table recv_log (
  id int identity(1,1) not null,
  mdn bigint not null,
  receive_mdn bigint not null,
  external_id varchar(36) not null,
  message varchar(1600) not null,
  from_city varchar(36) null,
  from_state char(2) null,
  from_zip varchar(10) null,
  from_country char(2) null,
  created datetime not null default getutcdate(),
  constraint [PK_id] primary key clustered ([id] asc),
)