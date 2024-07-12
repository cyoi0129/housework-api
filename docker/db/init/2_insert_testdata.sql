DELETE FROM users;
DELETE FROM masters;
DELETE FROM tasks;

INSERT INTO users (name, email, password)
VALUES ('テストユーザー','tester@test.com','9bba5c53a0545e0c80184b946153c9f58387e3bd1d4ee35740f29ac2e718b019');

INSERT INTO masters (userID, name, category, point)
VALUES (1,'洗濯','家事',20),(1,'寝かしつけ','育児',50);

INSERT INTO tasks (userID, masterID, person, date)
VALUES (1, 1, 'パパ','2024-11-22'),(1, 2, 'ママ','2024-12-12');

SELECT * FROM users;
SELECT * FROM masters;
SELECT * FROM tasks;