create table way.profiles
(
    id          uuid default uuid_generate_v4() not null
                constraint profiles_pk primary key,
    name        text not null,
    description text,
    tenant_id   uuid not null
                constraint profiles_tenants_fk references way.tenants,
    created_at  timestamp not null,
    updated_at  timestamp not null,
    deleted_at  timestamp
);