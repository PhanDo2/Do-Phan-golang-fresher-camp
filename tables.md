
    CREATE TABLE `users` (
      `Id` int NOT NULL AUTO_INCREMENT,
      `Name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
      `Address` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
      `Phone` char(10) COLLATE utf8_unicode_ci NOT NULL,
      `Gmail` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
      PRIMARY KEY (`Id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_unicode_ci

    CREATE TABLE `retaurents` (
      `Id` int NOT NULL AUTO_INCREMENT,
      `Name` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
      `Address` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
      `Phone` varchar(10) COLLATE utf8_unicode_ci NOT NULL,
      PRIMARY KEY (`Id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_unicode_ci


    CREATE TABLE `categories` (
      `Id` int NOT NULL AUTO_INCREMENT,
      `Name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
      `RestaurantID` int NOT NULL,
      PRIMARY KEY (`Id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_unicode_ci


    CREATE TABLE `products` (
      `Id` int NOT NULL AUTO_INCREMENT,
      `CategoryID` int NOT NULL,
      `Name` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
      `Conten` longtext COLLATE utf8_unicode_ci NOT NULL,
      `Price` double NOT NULL,
      `Discount` double DEFAULT NULL,
      `Hot` int DEFAULT NULL,
      `Image` json NOT NULL,
      PRIMARY KEY (`Id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_unicode_ci
