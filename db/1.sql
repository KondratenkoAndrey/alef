create table db_version (
    version int not null,
    applied_at timestamptz(0) not null default now()
);

--Версия БД
insert into db_version (version) values (2);