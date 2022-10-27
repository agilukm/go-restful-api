CREATE TABLE workspace_members(
                           id INT NOT NULL  AUTO_INCREMENT,
                           workspace_id INT NOT NULL,
                           user_id VARCHAR(255) NOT NULL ,
                           type VARCHAR(255),
                           PRIMARY KEY (id),
                           FOREIGN KEY (workspace_id) REFERENCES workspaces(id) on delete cascade on update cascade
) ENGINE=InnoDB;

CREATE UNIQUE INDEX idx_workspace_members_workspace_id_user_id
    ON workspace_members (workspace_id, user_id);