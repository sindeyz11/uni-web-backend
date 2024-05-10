CREATE TABLE languages
(
    id   int(10) unsigned NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO languages (id, name)
VALUES (1, 'Pascal'),
       (2, 'C'),
       (3, 'C++'),
       (4, 'JavaScript'),
       (5, 'PHP'),
       (6, 'Python'),
       (7, 'Java'),
       (8, 'Haskel'),
       (9, 'Clojure'),
       (10, 'Prolog'),
       (11, 'Scala');


CREATE TABLE form
(
    id        int(10) unsigned NOT NULL AUTO_INCREMENT,
    fio       varchar(255) NOT NULL,
    phone     varchar(11)  NOT NULL,
    email     varchar(255) NOT NULL,
    birthday  bigint       NOT NULL,
    gender    varchar(6)   NOT NULL,
    biography text         NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE forms_languages
(
    id      int(10) unsigned NOT NULL AUTO_INCREMENT,
    id_form int NOT NULL,
    id_lang int NOT NULL,
    PRIMARY KEY (id)
);