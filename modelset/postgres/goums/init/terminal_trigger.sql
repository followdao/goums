create or replace function ums.terminal_change_notify
(
) returns trigger
as $$
declare
begin
  if tg_op = 'UPDATE'
    then
      new.serial_number = old.serial_number;
      if new.active_status <> old.active_status or new.service_status <> old.service_status
        then
          perform pg_notify('terninal_notify',
                            json_build_object('id',old.id,'serial',old.serial_number,'active',new.active_status,'role',
                                              new.access_role,'status',new.service_status,'expiration',
                                              new.service_expiration,'tg_op',tg_op) :: text);
      end if;
  end if;
  return new;
end;
$$ language plpgsql;


drop trigger if exists terminal_status_notify on ums.terminal;

create trigger terminal_status_notify
  after update or insert or delete
  on ums.terminal
  for each row
execute procedure  ums.notify_trigger(); -- ums.terminal_change_notify();
