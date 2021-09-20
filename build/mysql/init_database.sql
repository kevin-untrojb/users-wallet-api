DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
                            `id` int unsigned NOT NULL AUTO_INCREMENT,
                            `first_name` varchar(100) NOT NULL,
                            `last_name` varchar(100) NOT NULL,
                            `alias` varchar(100) NOT NULL,
                            `email` varchar(255) NOT NULL,
                            `date_created` timestamp NOT NULL,
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;


INSERT INTO `user` (`id`,`first_name`, `last_name`, `alias`, `email`,`date_created`) VALUES ()
INSERT INTO `user` (`id`,`first_name`, `last_name`, `alias`, `email`,`date_created`) VALUES ()
INSERT INTO `user` (`id`,`first_name`, `last_name`, `alias`, `email`,`date_created`) VALUES ()
INSERT INTO `user` (`id`,`first_name`, `last_name`, `alias`, `email`,`date_created`) VALUES ()
INSERT INTO `user` (`id`,`first_name`, `last_name`, `alias`, `email`,`date_created`) VALUES ()
INSERT INTO `user` (`id`,`first_name`, `last_name`, `alias`, `email`,`date_created`) VALUES ()



DROP TABLE IF EXISTS `currency`;
CREATE TABLE `currency` (
                          `id` int unsigned NOT NULL AUTO_INCREMENT,
                          `name` varchar(255) NOT NULL,
                          `exponent` int NOT NULL,
                          `current_balance` varchar(255) NOT NULL,
                          PRIMARY KEY (`id`),
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `wallet`;
CREATE TABLE `wallet` (
                        `id` int unsigned NOT NULL AUTO_INCREMENT,
                        `user_id` int NOT NULL,
                        `currency_id` int NOT NULL,
                        `current_balance` varchar(255) NOT NULL,
                        `date_created` timestamp NOT NULL,
                        PRIMARY KEY (`id`),
                        FOREIGN KEY (`user_id`) REFERENCES user(`id`),
                        FOREIGN KEY (`currency_id`) REFERENCES currency(`id`)

) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `transaction`;
CREATE TABLE `transaction` (
                          `id` int unsigned NOT NULL AUTO_INCREMENT,
                          `wallet_id` int NOT NULL,
                          `transaction_type` varchar(255) NOT NULL,
                          `amount` varchar(255) NOT NULL,
                          `date_created` timestamp NOT NULL,
                          PRIMARY KEY (`id`),
                          FOREIGN KEY (`wallet_id`) REFERENCES wallet(`id`),
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;