
drop function if exists active(bigint
                              ,text
                              ,text
);
create or replace function ums.active
(
  active_code text
, serial_number text
, apk_type text
) returns setof ums.terminal
  language plpgsql
as $$
declare
  usr record;
begin
  -- 检查试用码是否存在
  apk_type = trim(apk_type);
  serial_number = trim(serial_number);
  active_code = trim(active_code);
  execute format(
      'select * from  ums.terminal where active_code=$1 and serial_number=$2') into usr using active_code::text, serial_number::text;
  if usr is null -- if record not exists, return err
    then
      raise exception 'not found';
    else -- 记录存在
      if usr.active_status is not true -- check this sn been active  or not
        then
          update ums.terminal
          set
            active_status = true
            , active_date = now()
            , service_expiration = now() + interval '1 year'
          where
            id = usr.id;
      end if;
      -- return
      return query execute format('select * from  ums.terminal where id=$1') using usr.id :: bigint;
  end if;
end
$$;
