create table if not exists way.rule_sets
(
  id         uuid                                                  not null
    constraint rule_sets_pk primary key,
  name       text                                                  not null,
  tenant_id  uuid                                                  references way.tenants on update cascade on delete set null,
  parent_id  uuid                                                  references way.rule_sets on update cascade on delete set null,
  created_at timestamp(3) with time zone default CURRENT_TIMESTAMP not null,
  updated_at timestamp(3) with time zone                           not null,
  deleted_at timestamp(3) with time zone
);
