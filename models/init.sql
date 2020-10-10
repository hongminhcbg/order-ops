DROP TABLE IF EXISTS `persons`;
CREATE TABLE `persons` (
    id INT(11) AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(64),
    age INT(11),
    address VARCHAR(128),
    is_married TINYINT(1)
);