create table content (
    id uuid default gen_random_uuid() not null constraint content_pk primary key,
    page_id uuid not null constraint content_page_fk references page on update cascade on delete restrict,
    tag varchar(100) not null,
    data text not null
);
comment on table content is 'Таблица контента';
comment on column content.id is 'Идентификатор контента';
comment on column content.page_id is 'Страница';
comment on column content.tag is 'Контент';
comment on column content.data is 'Тэг контента';
alter table content owner to postgres;
create index content_page_index on content (page_id);
create unique index content_page_tag_uindex on content (page_id, tag);

--Версия БД
insert into db_version (version) values (4);