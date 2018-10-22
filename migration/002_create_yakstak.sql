create table yakstak (
  id bigint primary key,
  public_id text not null,
  name text not null default 'Unnamed',
  insert_time timestamptz not null default now(),
  last_update_time timestamptz not null default now()
);
select set_default_to_next_duid_block('yakstak', 'id', 'yakstak_id_seq');

create index on yakstak (public_id);

grant select, insert, delete, update on table yakstak to {{.app_user}};

---- create above / drop below ----

drop table yakstak;
