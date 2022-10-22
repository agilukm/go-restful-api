CREATE TABLE products(
    id INT NOT NULL  AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL ,
    price integer NOT NULL default 0,
    PRIMARY KEY (id)
) ENGINE=InnoDB;