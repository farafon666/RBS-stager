-- psql -U postgres -f insertData.sql -b

\c assessment_manager;

--Справочник должностей
INSERT INTO assessment.t_ref_positions(c_position)
	VALUES ('Программист');
INSERT INTO assessment.t_ref_positions(c_position)
	VALUES ('Бухгалтер');
INSERT INTO assessment.t_ref_positions(c_position)
	VALUES ('Системный администратор');

-- Таблица сотрудников
INSERT INTO assessment.t_employees(
	fk_position, c_lastname, c_firstname, c_middlename, c_phone_number, c_email, c_birthdate, c_gender)
	VALUES (1, 'Иванов', 'Иван', 'Иванович', '6486468', 'safas@mail.ru', '12.06.1960', 'Мужчина');
INSERT INTO assessment.t_employees(
	fk_position, c_lastname, c_firstname, c_middlename, c_phone_number, c_email, c_birthdate, c_gender)
	VALUES (3, 'Фарафонов', 'Виталий', 'Геннадьевич', '4684684', 'saffsaasas@mail.ru', '10.04.2000', 'Мужчина');
INSERT INTO assessment.t_employees(
	fk_position, c_lastname, c_firstname, c_middlename, c_phone_number, c_email, c_birthdate, c_gender)
	VALUES (2, 'Радионов', 'Егор', 'Игоревич', '4654161', 'sdvfsg@mail.ru', '20.11.1987', 'Мужчина');
INSERT INTO assessment.t_employees(
	fk_position, c_lastname, c_firstname, c_middlename, c_phone_number, c_email, c_birthdate, c_gender)
	VALUES (1, 'Машкова', 'Екатерина', 'Владимировна', '13543351', 'dfhdsa@mail.ru', '10.03.1998', 'Женщина');

-- Справочник состояний кандидатов

-- Таблица кандидатов
INSERT INTO assessment.t_candidates(
    c_lastname, c_firstname, c_middlename, c_phone_number, c_email, c_birthdate, c_gender)
    VALUES ('Андреев', 'Георгий', 'Витальевич', '7344556', 'georg@mail.ru', '10.02.1996', 'Мужчина');
INSERT INTO assessment.t_candidates(
    c_lastname, c_firstname, c_middlename, c_phone_number, c_email, c_birthdate, c_gender)
    VALUES ('Пономарева', 'Ксения', 'Николаевна', '45464534', 'ksks@mail.ru', '24.05.1995', 'Женщина');
INSERT INTO assessment.t_candidates(
    c_lastname, c_firstname, c_middlename, c_phone_number, c_email, c_birthdate, c_gender)
    VALUES ('Корхова', 'Дарья', 'Викторовна', '7863452', 'devil@mail.ru', '01.04.2001', 'Женщина');
INSERT INTO assessment.t_candidates(
    c_lastname, c_firstname, c_middlename, c_phone_number, c_email, c_birthdate, c_gender)
    VALUES ('Саяпин', 'Илья', 'Сергеевич', '12134454', 'ilya@mail.ru', '18.11.2002', 'Мужчина');

-- Справочник состояний ассессментов
INSERT INTO assessment.t_ref_state_assessments(c_state)
	VALUES ('Предстоит');
INSERT INTO assessment.t_ref_state_assessments(c_state)
	VALUES ('Прошёл');
INSERT INTO assessment.t_ref_state_assessments(c_state)
	VALUES ('Заархивирован');

-- Таблица ассессментов
INSERT INTO assessment.t_assessments(fk_state, fk_disposer, c_date, c_title)
	VALUES (1, 1, '2021-04-19 10:00', 'Ассессмент');
INSERT INTO assessment.t_assessments(fk_state, fk_disposer, c_date, c_title)
	VALUES (2, 3, '2021-04-19 10:00', 'Ассессмент');
INSERT INTO assessment.t_assessments(fk_state, fk_disposer, c_date, c_title)
	VALUES (2, 2, '2021-04-06 14:39', 'Ассессмент');
	
-- Таблица связи сотрудников и ассессментов

-- Таблица связей кандидатов и ассессментов
