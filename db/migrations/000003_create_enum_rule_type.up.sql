do $$ begin
  create type rule_type as enum ('HTTP');
exception
  when duplicate_object then null;
end $$;
