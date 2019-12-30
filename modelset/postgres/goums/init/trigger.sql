drop function if exists notify_trigger(
) cascade;

create or replace function ums.notify_trigger
(
) returns trigger
as $trigger$
declare
  rec record;
  payload text;
  data text;
begin
  -- Set record row depending on operation
  case tg_op
    when 'INSERT' then
      rec := new;
    when 'UPDATE' then
      rec := new;
    when 'DELETE' then
      rec := old;
    else
      raise exception 'Unknown TG_OP: "%". Should not occur!', tg_op;
    end case;
-- build data from record
  select
    row_to_json(rec):: text
  into data;
  -- Build the payload
  payload :=
      json_build_object('timestamp',CURRENT_TIMESTAMP,'operation',tg_op,'schema',tg_table_schema,'table',tg_table_name,
                        'data',data);
  -- Notify the channel
  perform pg_notify('ums_notify',payload);
  return rec;
end;
$trigger$ language plpgsql;
-- trigger on ums.terminal
drop trigger if exists terminal_change_notify on ums.terminal;

create trigger terminal_change_notify
  after update or insert or delete
  on ums.terminal
  for each row
execute procedure ums.notify_trigger();
-- trigger on ums.apktype table
drop trigger if exists apk_type_change_notify on ums.apktype;
create trigger apk_type_change_notify
  after update or insert or delete
  on ums.apktype
  for each row
execute procedure ums.notify_trigger();
