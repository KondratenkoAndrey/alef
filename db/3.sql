create table page (
    id uuid default gen_random_uuid() not null constraint page_pk primary key,
    parent_id uuid constraint page_page_fk references page on update cascade on delete restrict,
    url varchar(255) not null,
    name varchar(50),
    title varchar(100),
    description varchar(255)
);
comment on table page is 'Таблица страниц сайта';
comment on column page.id is 'Идентификатор страницы';
comment on column page.parent_id is 'Идентификатор родительской страницы';
comment on column page.url is 'Относительный адрес';
comment on column page.name is 'Имя страницы';
comment on column page.title is 'Заголовок страницы';
comment on column page.description is 'Краткое описание';
alter table page owner to postgres;
create unique index page_url_uindex on page (url);
create index page_parent_id_index on page (parent_id);

--Версия БД
insert into db_version (version) values (3);