create table way.applications
(
    id          uuid not null
                constraint applications_pk primary key,
    name        text not null,
    description text,
    tenant_id   uuid  not null
                constraint applications_tenants_fk references way.tenants,
    version     text,
    created_at  timestamptz default current_timestamp not null,
    updated_at  timestamptz default current_timestamp not null,
    deleted_at  timestamptz
);