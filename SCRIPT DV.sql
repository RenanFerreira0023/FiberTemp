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




SELECT id , users_agent_id , channel_name , dt_create_channel    FROM channels WHERE users_agent_id = 1 AND dt_create_channel BETWEEN 2021-05-22 01:18:42 AND 2025-05-22 01:12:42  LIMIT 0,30;






