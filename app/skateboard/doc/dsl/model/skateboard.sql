CREATE TABLE `shop`  (
                         `id` int(0) NOT NULL,
                         `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
                         `introduction` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                         `address` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                         `context` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                         `created_at` bigint(0) NOT NULL,
                         `updated_at` bigint(0) NOT NULL,
                         `deleted_at` bigint(0) NOT NULL,
                         version    bigint       default 0,
                         PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;
