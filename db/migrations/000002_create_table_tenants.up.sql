create table if not exists way.tenants
(
  id          uuid                        default uuid_generate_v4() not null
    constraint tenants_pk primary key,
  name        text                                                   not null,
  description text,
  created_at   timestamp(3) with time zone default CURRENT_TIMESTAMP  not null,
  updated_at   timestamp(3) with time zone default CURRENT_TIMESTAMP  not null,
  deleted_at   timestamp(3) with time zone
);
