set search_path = ums;

insert
into
  ums.transition
(name,from_state,transition,to_state)
values
('turnstile','locked','coin','unlocked'),
('turnstile','unlocked','push','locked');



insert
into
  ums.transition
(name,from_state,transition,to_state)
values
('door','opened','close','closing'),
('door','closed','open','opening'),
('door','opening','is_opened','opened'),
('door','closing','is_closed','closed');



insert
into
  ums.machine
(name,state)
values
('turnstile','bar');



insert
into
  ums.machine
(name,state)
values
('fork','opened');



insert
into
  ums.machine
(id,name,state)
values
(1,'door','opened'),
(2,'door','closed'),
(3,'turnstile','locked'),
(4,'turnstile','unlocked');



update ums.machine
set
  state = 'closing'
where
  id = 1;



update ums.machine
set
  state = 'closed'
where
  id = 1;



update ums.machine
set
  state = 'closing'
where
  id = 2;



select *
from
  ums.do_transition(2,'open');


