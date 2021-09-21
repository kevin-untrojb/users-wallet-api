DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
                            `id` int unsigned NOT NULL AUTO_INCREMENT,
                            `first_name` varchar(100) NOT NULL,
                            `last_name` varchar(100) NOT NULL,
                            `alias` varchar(100) NOT NULL,
                            `email` varchar(255) NOT NULL,
                            `date_created` timestamp NOT NULL,
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;


INSERT INTO `user` (`id`,`first_name`, `last_name`, `alias`, `email`,`date_created`) VALUES (1,'roberto','robertino','tito','tito12@htmail.com',UTC_TIMESTAMP());
INSERT INTO `user` (`id`,`first_name`, `last_name`, `alias`, `email`,`date_created`) VALUES (2,'carlos','carlitos','carlitos','carlitos12@htmail.com',UTC_TIMESTAMP());
INSERT INTO `user` (`id`,`first_name`, `last_name`, `alias`, `email`,`date_created`) VALUES (3,'test','test','test','test@htmail.com',UTC_TIMESTAMP());




DROP TABLE IF EXISTS `currency`;
CREATE TABLE `currency` (
                            `id` int unsigned NOT NULL AUTO_INCREMENT,
                            `name` varchar(255) NOT NULL,
                            `exponent` int NOT NULL,
                            `date_created` timestamp NOT NULL,
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
INSERT INTO `currency` (`id`,`name`, `exponent`,`date_created`) VALUES (1,'ARS',2,UTC_TIMESTAMP());
INSERT INTO `currency` (`id`,`name`, `exponent`,`date_created`) VALUES (2,'BTC',8,UTC_TIMESTAMP());
INSERT INTO `currency` (`id`,`name`, `exponent`,`date_created`) VALUES (3,'USDT',2,UTC_TIMESTAMP());


DROP TABLE IF EXISTS `wallet`;
CREATE TABLE `wallet` (
                          `id` int unsigned NOT NULL AUTO_INCREMENT,
                          `user_id` int unsigned NOT NULL,
                          `currency_id` int unsigned NOT NULL,
                          `current_balance` varchar(255) NOT NULL,
                          `date_created` timestamp NOT NULL,
                          PRIMARY KEY (`id`),
                          FOREIGN KEY (`user_id`) REFERENCES user(`id`),
                          FOREIGN KEY (`currency_id`) REFERENCES currency(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;


INSERT INTO `wallet` (`id`,`user_id`, `currency_id`, `current_balance`,`date_created`) VALUES (1,1,1,'1000.00',UTC_TIMESTAMP());
INSERT INTO `wallet` (`id`,`user_id`, `currency_id`, `current_balance`,`date_created`) VALUES (2,1,2,'0.00023124',UTC_TIMESTAMP());
INSERT INTO `wallet` (`id`,`user_id`, `currency_id`, `current_balance`,`date_created`) VALUES (3,1,3,'34.24',UTC_TIMESTAMP());

INSERT INTO `wallet` (`id`,`user_id`, `currency_id`, `current_balance`,`date_created`) VALUES (4,3,1,'0.32',UTC_TIMESTAMP());
INSERT INTO `wallet` (`id`,`user_id`, `currency_id`, `current_balance`,`date_created`) VALUES (5,3,2,'0.00000110',UTC_TIMESTAMP());
INSERT INTO `wallet` (`id`,`user_id`, `currency_id`, `current_balance`,`date_created`) VALUES (6,3,3,'0.00',UTC_TIMESTAMP());



DROP TABLE IF EXISTS `transaction`;
CREATE TABLE `transaction` (
                               `id` int unsigned NOT NULL AUTO_INCREMENT,
                               `wallet_id` int unsigned NOT NULL,
                               `transaction_type` varchar(255) NOT NULL,
                               `amount` varchar(255) NOT NULL,
                               `date_created` timestamp NOT NULL,
                               PRIMARY KEY (`id`),
                               FOREIGN KEY (`wallet_id`) REFERENCES wallet(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

INSERT INTO `transaction` (`id`,`wallet_id`, `transaction_type`, `amount`,`date_created`) VALUES (1,1,'deposit','500.00',UTC_TIMESTAMP());
INSERT INTO `transaction` (`id`,`wallet_id`, `transaction_type`, `amount`,`date_created`)  VALUES (2,1,'deposit','600.00',UTC_TIMESTAMP());
INSERT INTO `transaction` (`id`,`wallet_id`, `transaction_type`, `amount`,`date_created`)  VALUES (3,1,'extraction','100.00',UTC_TIMESTAMP());
