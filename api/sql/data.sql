INSERT INTO users (name, nick, email, password)
VALUES
("User 1", "user_1", "user1@gmail.com", "$2a$10$YY./uzbA1QzKeokgiXIUOOqlKfwHwhQnt4K5b0xgmqROKzvIBtpOi"),
("User 2", "user_2", "user2@gmail.com", "$2a$10$YY./uzbA1QzKeokgiXIUOOqlKfwHwhQnt4K5b0xgmqROKzvIBtpOi"),
("User 3", "user_3", "user3@gmail.com", "$2a$10$YY./uzbA1QzKeokgiXIUOOqlKfwHwhQnt4K5b0xgmqROKzvIBtpOi"),
("User 4", "user_4", "user4@gmail.com", "$2a$10$YY./uzbA1QzKeokgiXIUOOqlKfwHwhQnt4K5b0xgmqROKzvIBtpOi"),
("User 5", "user_5", "user5@gmail.com", "$2a$10$YY./uzbA1QzKeokgiXIUOOqlKfwHwhQnt4K5b0xgmqROKzvIBtpOi");

INSERT INTO followers (user_id, follower_id)
VALUES
(1, 2),
(3, 1),
(1, 3),
(1, 4),
(2, 4),
(4, 3),
(3, 4),
(5, 2);