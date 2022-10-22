CREATE TABLE transactions(
                      id INT NOT NULL  AUTO_INCREMENT,
                      user_id INT NOT NULL ,
                      total int,
                      PRIMARY KEY (id)
) ENGINE=InnoDB;