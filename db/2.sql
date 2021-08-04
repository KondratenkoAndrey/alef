create table company_info (
    id uuid default gen_random_uuid() not null constraint company_info_pk primary key,
    short_name varchar(64) not null,
    full_name varchar(256) not null,
    phone varchar(16) not null,
    email varchar(64) not null
);
comment on table company_info is 'Информация о компании';
comment on column company_info.id is 'Идентификатор компании';
comment on column company_info.short_name is 'Краткое наименование';
comment on column company_info.full_name is 'Полное наименование';
comment on column company_info.phone is 'Телефон';
comment on column company_info.email is 'Электронная почта';
alter table company_info owner to postgres;
create unique index company_info_id_uindex on company_info (id);
create unique index company_info_short_name_uindex on company_info (short_name);

--Версия БД
insert into db_version (version) values (2);