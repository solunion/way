create table if not exists way.rules_rule_sets
(
  rule_id     uuid                                                  not null references way.rules on update cascade on delete cascade,
  rule_set_id uuid                                                  not null references way.rule_sets on update cascade on delete cascade,
  valid_from  timestamp(3) with time zone default CURRENT_TIMESTAMP not null,
  valid_to    timestamp(3) with time zone,
  created_at  timestamp(3) with time zone default CURRENT_TIMESTAMP not null,
  updated_at  timestamp(3) with time zone default CURRENT_TIMESTAMP not null,
  deleted_at  timestamp(3) with time zone,
  constraint rules_rule_sets_pk primary key (rule_id, rule_set_id)
);
