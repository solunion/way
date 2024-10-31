create table if not exists way.tenants
(
    id          uuid default uuid_generate_v4() not null
                constraint tenants_pk primary key,
    name        text not null,
    description text,
    created_at  timestamptz default current_timestamp not null,
    updated_at  timestamptz default current_timestamp not null,
    deleted_at  timestamptz
);