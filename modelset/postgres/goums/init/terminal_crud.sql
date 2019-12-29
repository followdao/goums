insert
into
  ums.terminal
(serial_number,active_code)
values
($1,$2)
returning id;

update ums.terminal
set
  active_status=$1
  , active_date = $2
  , max_active_session = $3
  , service_status = $4
where
  id = $5;
