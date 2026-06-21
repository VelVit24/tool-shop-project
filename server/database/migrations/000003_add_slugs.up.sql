alter table products add column slug text not null unique;
alter table categories add column slug text not null unique;
