/*
 Navicat Premium Data Transfer

 Source Server         : 虚拟机192.168.10.47
 Source Server Type    : MySQL
 Source Server Version : 50728
 Source Host           : 192.168.10.47:3306
 Source Schema         : oldme

 Target Server Type    : MySQL
 Target Server Version : 50728
 File Encoding         : 65001

 Date: 25/11/2022 17:29:22
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `nickname` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '昵称',
  `avatar` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '头像，base64',
  `register` datetime(0) NOT NULL COMMENT '注册时间',
  `salt` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '随机盐值',
  `last_login` datetime(0) NOT NULL COMMENT '最后登录时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username唯一`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin
-- ----------------------------
INSERT INTO `admin` VALUES (1, 'admin', '7e4f9f6bce4319e8dce45dd0dca30790', '灯火消逝的码头', 'data:image/png;base64,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', '2022-11-16 23:38:21', 'taSzTJ70v3CFbmQvraPYngl7ojybFBHS', '2022-11-16 23:38:21');

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `group_id` int(11) NOT NULL COMMENT '分组id',
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
  `author` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '作者',
  `thumb` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图片地址',
  `key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '关键词，依英文逗号隔开',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '简介',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '内容',
  `order` smallint(5) NOT NULL DEFAULT 0 COMMENT '排序，越大越靠前',
  `ontop` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否置顶',
  `onshow` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否显示',
  `hist` mediumint(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '点击数',
  `post` mediumint(8) NOT NULL COMMENT '评论数',
  `created_at` datetime(0) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NOT NULL COMMENT '更新时间',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  `lasted_at` datetime(0) NULL DEFAULT NULL COMMENT '最后浏览时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for article_group
-- ----------------------------
DROP TABLE IF EXISTS `article_group`;
CREATE TABLE `article_group`  (
  `id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `onshow` tinyint(1) UNSIGNED NOT NULL COMMENT '是否显示',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
