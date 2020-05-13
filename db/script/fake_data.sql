SET FOREIGN_KEY_CHECKS = 0;
truncate table activity_task_record;
truncate table activity_task;
truncate table activity_process;
truncate table user;
truncate table activity;
truncate table code_authority;
truncate table code_activity_status;
truncate table activity_team;
insert into code_activity_status (code, name) values (1, '未开始'), (2, '进行中'), (3, '已结束');

insert into code_authority (authority_code, name) values (1, '公开'), (2, '受邀请');


INSERT INTO user (user_id, openid, username, avatar) VALUES (1, 'testopenidid', 'test_user', '-');

insert into activity (activity_id, activity_name, introduction, creator_id, regulation, roles, authority_code, illustration, current_process_id, status_code)
values (1, '吃屎大赛', '活动介绍', 1, '-', '吃屎头领,吃屎队员', 1, '-', null, 1);

insert into activity_process (process_id, process_name, activity_id) values (1, '粑粑生产', 1);

insert into activity_team (activity_id, team_name, slogan) values (1, '老八冠军队', 'Very Good');

SET FOREIGN_KEY_CHECKS = 1;