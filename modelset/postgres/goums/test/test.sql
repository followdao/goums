--aaaaa
select
  id
  , serial_number
  , active_status
  , active_date
  , max_active_session
  , access_role
  , service_status
  , service_expiration
from
  ums.active('11112222','bbbbbbbb','com.apple.liveTV');

select
  id
  , serial_number
  , active_status
  , active_date
  , max_active_session
  , access_role
  , service_status
  , service_expiration
from
  ums.terminal;

