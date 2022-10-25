CREATE TABLE workspaces(
                         id INT NOT NULL  AUTO_INCREMENT,
                         name VARCHAR(255) NOT NULL ,
                         user_id VARCHAR(255) NOT NULL ,
                         token VARCHAR(255) default NULL,
                         token_expired_at DATETIME default NULL,
                         created_at DATETIME default CURRENT_TIMESTAMP,
                         updated_at DATETIME default CURRENT_TIMESTAMP on UPDATE CURRENT_TIMESTAMP,
                         deleted_at DATETIME default NULL,
                         PRIMARY KEY (id)
) ENGINE=InnoDB;