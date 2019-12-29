create trigger apk_type_change_notify
  after update or insert or delete
  on ums.apktype
  for each row
execute procedure ums.notify_trigger();
