create table if not exists way.tenants
(
    id          uuid default uuid_generate_v4() not null
                constraint tenants_pk primary key,
    name        text not null,
    description text,
    created_at  timestamp not null,
    updated_at  timestamp not null,
    deleted_at  timestamp
);