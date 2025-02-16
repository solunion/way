create table if not exists way.rules
(
  id          uuid                                                  not null
    constraint rules_pk primary key,
  name        text                                                  not null,
  description text,
  type        rule_type                                             not null,
  value       jsonb                                                 not null,
  tenant_id   uuid                                                  references way.tenants on update cascade on delete set null,
  created_at  timestamp(3) with time zone default CURRENT_TIMESTAMP not null,
  updated_at  timestamp(3) with time zone default CURRENT_TIMESTAMP not null,
  deleted_at  timestamp(3) with time zone
);
