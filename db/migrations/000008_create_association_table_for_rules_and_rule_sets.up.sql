create table way.rule_sets_rules
(
    id          uuid not null
                constraint rule_sets_rules_pk primary key,
    rule_set_id uuid not null
                constraint rule_sets_rules_rule_sets_fk references way.rule_sets,
    rule_id     uuid not null
                constraint rule_sets_rules_rules_fk references way.rules,
    valid_from  timestamp not null,
    valid_to    timestamp,
    created_at  timestamptz default current_timestamp not null,
    updated_at  timestamptz default current_timestamp not null,
    deleted_at  timestamptz
);