/* SQLEditor (MySQL (2))*/


CREATE TABLE activity_process
(
    process_id INT,
    process_name INT,
    activity_id INT,
    PRIMARY KEY (process_id)
);

CREATE TABLE activity_task
(
    task_id INT,
    task_name INT,
    process_id INT,
    role VARCHAR(25),
    PRIMARY KEY (task_id)
);

CREATE TABLE code_activity_status
(
    code INT UNIQUE ,
    name VARCHAR(12) NOT NULL
);

CREATE TABLE code_authority
(
    authority_code INT,
    name VARCHAR(20) NOT NULL,
    PRIMARY KEY (authority_code)
);

CREATE TABLE activity
(
    activity_id INT,
    activity_name VARCHAR(128) NOT NULL,
    introduction VARCHAR(1024) NOT NULL,
    creator_id INTEGER NOT NULL,
    regulation VARCHAR(5) NOT NULL,
    roles VARCHAR(255) NOT NULL,
    authority_code INT,
    illustration VARCHAR(255) NOT NULL,
    current_process INT NOT NULL,
    status_code INT,
    PRIMARY KEY (activity_id)
);

CREATE TABLE system_config
(
    `key` VARCHAR(100),
    value VARCHAR(255) NOT NULL,
    PRIMARY KEY (`key`)
);

CREATE TABLE team
(
    team_id INT,
    activity_id INTEGER,
    team_name VARCHAR(64),
    PRIMARY KEY (team_id)
);

CREATE TABLE activity_task_record
(
    team_id INT,
    task_id INT,
    PRIMARY KEY (team_id,task_id)
);

CREATE TABLE user
(
    user_id INT,
    openid VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    avatar VARCHAR(255) NOT NULL,
    PRIMARY KEY (user_id)
);

CREATE TABLE team_join
(
    team_join_id INTEGER,
    user_id INT NOT NULL,
    team_id INT NOT NULL,
    role VARCHAR(24) NOT NULL,
    PRIMARY KEY (team_join_id)
);

ALTER TABLE activity_process ADD FOREIGN KEY activity_id_idxfk (activity_id) REFERENCES activity (activity_id);

ALTER TABLE activity_task ADD FOREIGN KEY process_id_idxfk (process_id) REFERENCES activity_process (process_id);

ALTER TABLE activity ADD FOREIGN KEY authority_code_idxfk (authority_code) REFERENCES code_authority (authority_code);

ALTER TABLE activity ADD FOREIGN KEY current_process_idxfk (current_process) REFERENCES activity_process (process_id);

ALTER TABLE activity ADD FOREIGN KEY status_code_idxfk (status_code) REFERENCES code_activity_status (code);

ALTER TABLE activity_task_record ADD FOREIGN KEY team_id_idxfk (team_id) REFERENCES team (team_id);

ALTER TABLE activity_task_record ADD FOREIGN KEY task_id_idxfk (task_id) REFERENCES activity_task (task_id);

ALTER TABLE team_join ADD FOREIGN KEY user_id_idxfk (user_id) REFERENCES user (user_id);

ALTER TABLE team_join ADD FOREIGN KEY team_id_idxfk_1 (team_id) REFERENCES team (team_id);
