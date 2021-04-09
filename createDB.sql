-- psql -U postgres -f createDB.sql -b

create database assessment_manager;

\c assessment_manager;

create schema assessment;

-- Справочник должностей
create table assessment.t_ref_positions(
    pk_id           serial not null primary key, --Первичный ключ должности
    c_position      varchar not null             --Должность
);
-- Таблица сотрудников
create table assessment.t_employees(
    pk_id           serial not null primary key, --Первичный ключ сотрудника
    fk_position     int not null references assessment.t_ref_positions(pk_id), --Внешний ключ должности
    c_lastname      varchar not null,       --Фамилия сотрудника
    c_firstname     varchar not null,       --Имя сотрудника
    c_middlename    varchar,                --Отчесво сотрудника
    c_phone_number  varchar,                --Номер телефона сотрудника
    c_email         varchar,                --Адрес электронной почты сотрудника
    c_birthdate     varchar not null,       --Дата рождения сотрудника
    c_gender        varchar not null        --Пол сотрудника
);
-- Таблица кандидатов
create table assessment.t_candidates(
    pk_id           serial  not null primary key, --Первичный ключ кандидата
    c_lastname      varchar not null,   --Фамилия кандидата  
    c_firstname     varchar not null,   --Имя кандидата              
    c_middlename    varchar,            --Отчество кандидата  
    c_phone_number  varchar,            --Номер телефона кандидата
    c_email         varchar,            --Адрес электронной почты кандидата
    c_birthdate     varchar not null,   --Дата рождения кандидата
    c_gender        varchar not null    --Пол кандидата
);
-- Справочник состояний ассессментов
create table assessment.t_ref_state_assessments(
    pk_id               serial  not null primary key,   --Первичный ключ сотояния ассессмента
    c_state             varchar not null                --Состояние ассессмента
);
-- Таблица ассессментов
create table assessment.t_assessments(
    pk_id           serial  not null primary key,   --Первичный ключ ассессмента
    fk_state        int     not null references assessment.t_ref_state_assessments(pk_id),--Внешний ключ ассессмента указывающий на состояние
    fk_disposer     int     not null references assessment.t_employees(pk_id),  --Внешний ключ распорядителя
    c_date          varchar not null,   --Дата проведения ассессмента
    c_title         varchar not null    --Название мероприятия
);
-- Таблица связи сотрудников и ассессментов
create table assessment.toc_employees_assessments(
    pk_id           serial  not null primary key,   --Первичный ключ связи сотрудников и ассессментов
    fk_employee     int not null references assessment.t_employees(pk_id),     --Внешний ключ сотрудника 
    fk_assessment   int not null references assessment.t_assessments(pk_id)    --Внешний ключ ассессмента
);
-- Таблица связей кандидатов и ассессментов
create table assessment.toc_candidates_assessments(
    pk_id               serial not null primary key,    --Первичный ключ связи сотрудников
    fk_candidate        int not null references assessment.t_candidates(pk_id),            --Внешний ключ кандидата
    fk_assessment       int not null references assessment.t_assessments(pk_id)            --Внешний ключ ассессмента
);