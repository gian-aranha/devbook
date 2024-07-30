INSERT INTO users (name, nick, email, password)
VALUES
("User 1", "user_1", "user1@gmail.com", "$2a$10$YY./uzbA1QzKeokgiXIUOOqlKfwHwhQnt4K5b0xgmqROKzvIBtpOi"),
("User 2", "user_2", "user2@gmail.com", "$2a$10$YY./uzbA1QzKeokgiXIUOOqlKfwHwhQnt4K5b0xgmqROKzvIBtpOi"),
("User 3", "user_3", "user3@gmail.com", "$2a$10$YY./uzbA1QzKeokgiXIUOOqlKfwHwhQnt4K5b0xgmqROKzvIBtpOi");

INSERT INTO followers (user_id, follower_id)
VALUES
(1, 2),
(3, 1),
(1, 3);

INSERT INTO posts (title, content, author_id) 
VALUES
("User 1 Post", "This is the post of user 1", 1),
("User 2 Post", "This is the post of user 2", 2),
("User 3 Post", "This is the post of user 3", 3);