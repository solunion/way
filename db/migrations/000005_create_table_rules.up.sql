create table way.rules
(
    id          uuid default uuid_generate_v4() not null
                constraint rules_pk primary key,
    name        text not null,
    description text,
    tenant_id   uuid not null
                constraint rules_tenants_fk references way.tenants,
    type        text not null
                constraint rules_types CHECK ( type in ('API') ),
    data        jsonb not null,
    created_at  timestamptz default current_timestamp not null,
    updated_at  timestamptz default current_timestamp not null,
    deleted_at  timestamptz
);