drop schema if exists queue cascade;
create schema if not exists queue;
create sequence queue.global_id_sequence;

create or replace function queue.id_generator
(
  out result bigint
)
as $$
declare
  our_epoch bigint := 1314220021721;
  seq_id bigint;
  now_millis bigint;
  -- the id of this DB shard, must be set for each
  -- schema shard you have - you could pass this as a parameter too
  shard_id int := 1;
begin
  select
    nextval('queue.global_id_sequence') % 1024
  into seq_id;
  select
    floor(extract(epoch from clock_timestamp()) * 1000)
  into now_millis;
  result := (now_millis - our_epoch) << 23;
  result := result | (shard_id << 10);
  result := result | (seq_id);
end;
$$ language PLPGSQL;
-- explain (buffers ,analyse )
select queue.id_generator();


