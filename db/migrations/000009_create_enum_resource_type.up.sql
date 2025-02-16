do $$ begin
  create type resource_type as enum ('HTTP');
exception
  when duplicate_object then null;
end $$;
