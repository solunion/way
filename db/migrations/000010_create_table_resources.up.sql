create table if not exists way.resources
(
  id           uuid                        default uuid_generate_v4() not null
    constraint resources_pk primary key,
  name         text                                                   not null,
  type         resource_type                                          not null,
  component_id uuid                                                   not null references way.components on update cascade on delete restrict,
  tenant_id    uuid                                                   references way.tenants on update cascade on delete set null,
  created_at   timestamp(3) with time zone default CURRENT_TIMESTAMP  not null,
  updated_at   timestamp(3) with time zone default CURRENT_TIMESTAMP  not null,
  deleted_at   timestamp(3) with time zone
);
