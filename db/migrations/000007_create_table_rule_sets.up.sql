create table way.rule_sets
(
    id          uuid not null
                constraint rule_sets_pk primary key,
    name        text not null,
    description text,
    tenant_id   uuid not null
                constraint rule_sets_tenants_fk references way.tenants,
    created_at  timestamp not null,
    updated_at  timestamp not null,
    deleted_at  timestamp
);