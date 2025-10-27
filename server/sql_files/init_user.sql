USE halloween_code;
DROP TABLE user;
CREATE TABLE user (
		user_id INT AUTO_INCREMENT PRIMARY KEY,
		username varchar(50),
		email varchar(50),
		hashedpassword varchar(255),
		firstname varchar(50),
		surname varchar(50)
	);
    
