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


