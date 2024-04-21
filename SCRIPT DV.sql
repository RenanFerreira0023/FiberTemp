CREATE DATABASE rds_db;

use rds_db;


select * from users_agent;


UPDATE users_agent
SET password_agent = '55a5e9e78207b4df8699d60886fa070079463547b095d1a05bc719bb4e6cd251'
WHERE id = 1;


SELECT id, first_name, second_name, email FROM users_receptor ;

select * from users_receptor;

select * from channels;



# 		criar select lista 

SELECT * FROM permission WHERE user_receptor_id = 4 AND channel_id = 4;

#  4

(SELECT id FROM permission WHERE channel_id = 4);



SELECT id, first_name, second_name, email FROM users_receptor WHERE id >= 4;


SELECT id, first_name, second_name, email FROM users_receptor WHERE id IN (SELECT id FROM permission WHERE channel_id != 4);

SELECT  users_receptor.id,  users_receptor.first_name,  users_receptor.second_name,  users_receptor.email,   (SELECT channel_id  FROM permission  WHERE permission.user_receptor_id = users_receptor.id) 
		 AS channel_id FROM      
			 users_receptor  WHERE      
              users_receptor.id IN 
									 (SELECT user_receptor_id FROM permission WHERE channel_id = 4);


select * from users_receptor;

select * from permission;

DELETE FROM permission  WHERE user_receptor_id = 3 AND channel_id = 4;


select * from permission;
select * from channels;
	
    
 use rds_db;   
    
SELECT ur.id,  ur.first_name,  ur.second_name,  ur.email
FROM users_receptor AS ur
LEFT JOIN Permission AS p ON ur.ID = p.user_receptor_id
WHERE p.user_receptor_id IS NULL;



SELECT ur.id, ur.first_name, ur.second_name, ur.email FROM users_receptor AS ur LEFT JOIN Permission AS p ON ur.ID = p.user_receptor_id WHERE p.user_receptor_id IS NULL OR (p.channel_id IS NOT NULL AND p.channel_id != 4);


    
    
select * from Permission ;



SELECT  users_receptor.id,  users_receptor.first_name,  users_receptor.second_name,  users_receptor.email,   permission.channel_id 
FROM    users_receptor 
JOIN    permission ON users_receptor.id = permission.user_receptor_id
WHERE   permission.channel_id = 3;    

SELECT ur.id, ur.first_name, ur.second_name, ur.email 
FROM users_receptor AS ur 
LEFT JOIN Permission AS p 
ON ur.ID = p.user_receptor_id 
WHERE p.user_receptor_id IS NULL 
OR (p.channel_id IS NOT NULL AND p.channel_id != 4);


select* from permission where channel_id != 4;




SELECT id , users_agent_id , channel_name , dt_create_channel    
FROM channels 
WHERE users_agent_id = 1 
AND dt_create_channel 
BETWEEN '2021-05-22 01:18:42' AND '2025-05-22 01:12:42'  
LIMIT 0,30;


select * from permission;



SELECT  c.id,  c.users_agent_id,  c.channel_name,  c.dt_create_channel,
(SELECT COUNT(*) FROM Permission p WHERE p.channel_id = c.id) AS total_receptor_copy
FROM  channels c
WHERE  c.users_agent_id = 1 
AND c.dt_create_channel BETWEEN '2021-05-22 01:18:42' AND '2025-05-22 01:12:42'
LIMIT 0,30;



select * from channels;

select * from all_copy;
use rds_db;
select * from users_agent;

SELECT channel_name,dt_create_channel 
FROM channels
WHERE id = 5;




SELECT  c.channel_name,  c.dt_create_channel,
(SELECT COUNT(*) FROM all_copy ac WHERE ac.channel_id = c.id) AS count_channel
FROM  channels c
WHERE  c.id = 5;


 use rds_db;
select * from users_agent; 

SELECT password_agent FROM users_agent WHERE email = 'admin@rdstrader.com';


UPDATE users_agent
SET password_agent = IF(email = 'admin@rdstrader.com', NULL, password_agent)
WHERE email = 'admin@rdstrader.com' AND id = 1;


UPDATE users_agent
SET password_agent = '55a5e9e78207b4df8699d60886fa070079463547b095d1a05bc719bb4e6cd251'
WHERE id = 1;


UPDATE users_agent
SET password_agent = NULL
WHERE id = (
    SELECT id
    FROM users_agent
    WHERE email = 'admin@rdstrader.com'
)
AND password_agent IS NOT NULL;


UPDATE users_agent
SET password_agent = NULL
WHERE email = 'admin@rdstrader.com'
AND password_agent IS NOT NULL
AND id = (
    SELECT id
    FROM users_agent
    WHERE email = 'admin@rdstrader.com'
);
SELECT id, first_name, email
FROM users_agent
WHERE email = 'admin@rdstrader.com';


use rds_db;
select * from channels;




SELECT id, users_agent_id, channel_name, dt_create_channel 
FROM channels 
WHERE users_agent_id = 2 AND dt_create_channel 
BETWEEN '2021-05-22 01:12:42' AND '2025-05-22 01:18:42' LIMIT 0, 30;

SELECT id, users_agent_id, channel_name, dt_create_channel FROM channels WHERE users_agent_id = 2 AND dt_create_channel BETWEEN '2021-05-22 01:18:42' AND '2025-05-22 01:12:42' LIMIT 0, 30;


SELECT ur.id, ur.first_name, ur.second_name, ur.email 
FROM users_receptor AS ur LEFT JOIN Permission AS p ON ur.ID = p.user_receptor_id 
WHERE p.user_receptor_id IS NULL OR (p.channel_id IS NOT NULL AND p.channel_id != 16);



select * from Permission;

select * from users_receptor;
select * from users_agent;


#SELECT  ur.id, ur.first_name, ur.second_name, ur.email 
SELECT  *
FROM users_receptor AS ur
WHERE ur.id NOT IN (SELECT user_receptor_id FROM Permission WHERE channel_id = 16 ) 
AND agent_id = (SELECT ua.id FROM users_agent AS ua WHERE (ua.id = ur.agent_id));



AND agent_id = (SELECT ua.id FROM users_agent AS ua WHERE (ua.id = ur.agent_id));




