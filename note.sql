
    CREATE TABLE `note` (
      `id` int(11) NOT NULL AUTO_INCREMENT,
      `name` varchar(50) NOT NULL,
      `addr` varchar(255) NOT NULL,
      `status` int(11) NOT NULL DEFAULT '1',
      PRIMARY KEY (`id`),
      KEY `status` (`status`) USING BTREE
    ) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8
