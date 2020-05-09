/* SQLEditor (MySQL (2))*/


CREATE TABLE activity_process
(
    process_id INT NOT NULL AUTO_INCREMENT,
    process_name VARCHAR(50) NOT NULL,
    activity_id INT NOT NULL,
    start_time DATETIME NOT NULL,
    PRIMARY KEY (process_id)
);

CREATE TABLE activity_task
(
    task_id INT NOT NULL AUTO_INCREMENT,
    task_name VARCHAR(50) NOT NULL,
    process_id INT NOT NULL,
    role VARCHAR(25) NOT NULL,
    PRIMARY KEY (task_id)
);

CREATE TABLE activity_team
(
    team_id INT NOT NULL AUTO_INCREMENT,
    activity_id INTEGER NOT NULL,
    team_name VARCHAR(64) NOT NULL,
    PRIMARY KEY (team_id)
);

CREATE TABLE code_activity_status
(
    code INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(12) NOT NULL,
    PRIMARY KEY (code)
);

CREATE TABLE activity_task_record
(
    team_id INT,
    task_id INT,
    PRIMARY KEY (team_id,task_id)
);

CREATE TABLE code_authority
(
    authority_code INT AUTO_INCREMENT,
    name VARCHAR(20) NOT NULL,
    PRIMARY KEY (authority_code)
);

CREATE TABLE activity
(
    activity_id INT AUTO_INCREMENT,
    activity_name VARCHAR(128) NOT NULL,
    introduction VARCHAR(1024) NOT NULL,
    creator_id INTEGER NOT NULL,
    regulation VARCHAR(5) NOT NULL,
    roles VARCHAR(255) NOT NULL,
    authority_code INT NOT NULL,
    illustration VARCHAR(255) NOT NULL,
    current_process_id INT,
    status_code INT NOT NULL,
    PRIMARY KEY (activity_id)
);

CREATE TABLE system_config
(
    `key` VARCHAR(100) NOT NULL,
    value VARCHAR(255) NOT NULL,
    PRIMARY KEY (`key`)
);

CREATE TABLE user
(
    user_id INT NOT NULL AUTO_INCREMENT,
    openid VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    avatar VARCHAR(255) NOT NULL,
    PRIMARY KEY (user_id)
);

CREATE TABLE activity_join
(
    team_join_id INTEGER NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL,
    team_id INT NOT NULL,
    role VARCHAR(24) NOT NULL,
    activity_id INT NOT NULL,
    PRIMARY KEY (team_join_id)
);

ALTER TABLE activity_process ADD FOREIGN KEY activity_id_idxfk (activity_id) REFERENCES activity (activity_id);

ALTER TABLE activity_task ADD FOREIGN KEY process_id_idxfk (process_id) REFERENCES activity_process (process_id);

ALTER TABLE activity_task_record ADD FOREIGN KEY team_id_idxfk (team_id) REFERENCES activity_team (team_id);

ALTER TABLE activity_task_record ADD FOREIGN KEY task_id_idxfk (task_id) REFERENCES activity_task (task_id);

ALTER TABLE activity ADD FOREIGN KEY authority_code_idxfk (authority_code) REFERENCES code_authority (authority_code);

ALTER TABLE activity ADD FOREIGN KEY current_process_id_idxfk (current_process_id) REFERENCES activity_process (process_id);

ALTER TABLE activity ADD FOREIGN KEY status_code_idxfk (status_code) REFERENCES code_activity_status (code);

ALTER TABLE activity_join ADD FOREIGN KEY user_id_idxfk (user_id) REFERENCES user (user_id);

ALTER TABLE activity_join ADD FOREIGN KEY team_id_idxfk_1 (team_id) REFERENCES activity_team (team_id);

ALTER TABLE activity_join ADD FOREIGN KEY activity_id_idxfk_1 (activity_id) REFERENCES activity (activity_id);
