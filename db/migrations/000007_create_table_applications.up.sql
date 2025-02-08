create table if not exists way.applications
(
  id          uuid                        default uuid_generate_v4() not null
    constraint applications_pk primary key,
  name        text                                                   not null,
  description text,
  tenant_id   uuid                                                   references way.tenants on update cascade on delete set null,
  created_at  timestamp(3) with time zone default CURRENT_TIMESTAMP  not null,
  updated_at  timestamp(3) with time zone default CURRENT_TIMESTAMP  not null,
  deleted_at  timestamp(3) with time zone
);

