/* SQLEditor (MySQL (2))*/

CREATE TABLE code_activity_status
(
    code INTEGER UNIQUE ,
    name VARCHAR(12) NOT NULL
);

CREATE TABLE code_authority
(
    authority_code INTEGER,
    name VARCHAR(20) NOT NULL,
    PRIMARY KEY (authority_code)
);

CREATE TABLE system_config
(
    `key` VARCHAR(100),
    value VARCHAR(255) NOT NULL,
    PRIMARY KEY (`key`)
);

CREATE TABLE activity
(
    activity_id INTEGER,
    activity_name VARCHAR(128) NOT NULL,
    introduction VARCHAR(1024) NOT NULL,
    creator_id INTEGER NOT NULL,
    regulation VARCHAR(5) NOT NULL,
    roles VARCHAR(255) NOT NULL,
    authority_code INTEGER,
    illustration VARCHAR(255) NOT NULL,
    process VARCHAR(1024) NOT NULL,
    current_process VARCHAR(24) NOT NULL,
    status_code INTEGER,
    PRIMARY KEY (activity_id)
);

CREATE TABLE team
(
    team_id INTEGER,
    activity_id INTEGER,
    team_name VARCHAR(64),
    PRIMARY KEY (team_id)
);

CREATE TABLE user
(
    user_id INTEGER,
    openid VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    avatar VARCHAR(255) NOT NULL,
    PRIMARY KEY (user_id)
);

CREATE TABLE team_join
(
    team_join_id INTEGER,
    user_id INTEGER NOT NULL,
    team_id INTEGER NOT NULL,
    role VARCHAR(24) NOT NULL,
    PRIMARY KEY (team_join_id)
);

ALTER TABLE activity ADD FOREIGN KEY authority_code_idxfk (authority_code) REFERENCES code_authority (authority_code);

ALTER TABLE activity ADD FOREIGN KEY status_code_idxfk (status_code) REFERENCES code_activity_status (code);

ALTER TABLE team_join ADD FOREIGN KEY user_id_idxfk (user_id) REFERENCES user (user_id);

ALTER TABLE team_join ADD FOREIGN KEY team_id_idxfk (team_id) REFERENCES team (team_id);
