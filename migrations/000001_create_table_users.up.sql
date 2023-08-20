START TRANSACTION;
CREATE TABLE `users` ( 
  `uid_user` VARCHAR(36) NOT NULL,
  `created_at` TIMESTAMP NOT NULL,
  `updated_at` TIMESTAMP NOT NULL,
  `deleted_at` TIMESTAMP NULL,
  `email` VARCHAR(250) NOT NULL,
  `password` VARCHAR(250) NOT NULL,
  `username` VARCHAR(250) NOT NULL,
   PRIMARY KEY (`uid_user`)
);
COMMIT;
