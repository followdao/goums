drop trigger if exists terminal_change_notify on ums.terminal;

create trigger terminal_change_notify
  after update or insert or delete
  on ums.terminal
  for each row
execute procedure ums.notify_trigger();
