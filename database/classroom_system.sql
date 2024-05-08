/*
 Navicat Premium Data Transfer

 Source Server         : con
 Source Server Type    : MySQL
 Source Server Version : 80033
 Source Host           : localhost:3306
 Source Schema         : classroom_system

 Target Server Type    : MySQL
 Target Server Version : 80033
 File Encoding         : 65001

 Date: 08/05/2024 10:56:22
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for authcodes
-- ----------------------------
DROP TABLE IF EXISTS `authcodes`;
CREATE TABLE `authcodes`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `teacher_card` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `grade` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `college` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `major` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `class` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `course_id` int(0) NULL DEFAULT NULL,
  `authcode` int(0) NULL DEFAULT NULL,
  `create_time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of authcodes
-- ----------------------------
INSERT INTO `authcodes` VALUES (1, 'gdxz003', '22', '计算机学院', '软件技术', '4班', 1, 15693, '2024-05-07 19:50:09');
INSERT INTO `authcodes` VALUES (2, 'gdxz003', '22', '计算机学院', '软件技术', '4班', 7, 52621, '2024-05-07 20:16:23');
INSERT INTO `authcodes` VALUES (3, 'gdxz003', '22', '计算机学院', '软件技术', '4班', 7, 38452, '2024-05-08 09:56:27');
INSERT INTO `authcodes` VALUES (4, 'gdxz003', '22', '计算机学院', '软件技术', '4班', 4, 90028, '2024-05-08 09:58:17');

-- ----------------------------
-- Table structure for classes
-- ----------------------------
DROP TABLE IF EXISTS `classes`;
CREATE TABLE `classes`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `grade` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `college` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `major` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of classes
-- ----------------------------
INSERT INTO `classes` VALUES (1, '22', '计算机学院', '软件技术', '1班');
INSERT INTO `classes` VALUES (2, '22', '计算机学院', '软件技术', '2班');
INSERT INTO `classes` VALUES (3, '22', '计算机学院', '软件技术', '3班');
INSERT INTO `classes` VALUES (4, '22', '计算机学院', '软件技术', '4班');
INSERT INTO `classes` VALUES (5, '23', '计算机学院', '软件技术', '1班');
INSERT INTO `classes` VALUES (6, '23', '计算机学院', '软件技术', '2班');

-- ----------------------------
-- Table structure for colleges
-- ----------------------------
DROP TABLE IF EXISTS `colleges`;
CREATE TABLE `colleges`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `campus` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of colleges
-- ----------------------------
INSERT INTO `colleges` VALUES (1, '计算机学院', '白云校区');
INSERT INTO `colleges` VALUES (2, '财经商贸学院', '白云校区');
INSERT INTO `colleges` VALUES (3, '文化艺术学院', '白云校区');
INSERT INTO `colleges` VALUES (4, '公共管理学院', '花都校区');
INSERT INTO `colleges` VALUES (5, '电子信息学院', '花都校区');
INSERT INTO `colleges` VALUES (6, '外国语学院', '花都校区');

-- ----------------------------
-- Table structure for courses
-- ----------------------------
DROP TABLE IF EXISTS `courses`;
CREATE TABLE `courses`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `teacher_card` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `grade` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `college` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `major` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `class` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `course` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `start_time` datetime(0) NULL DEFAULT NULL,
  `end_time` datetime(0) NULL DEFAULT NULL,
  `classroom` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of courses
-- ----------------------------
INSERT INTO `courses` VALUES (1, 'gdxz003', '23', '计算机学院', '软件技术', '2班', 'Java程序设计', '2024-03-05 08:10:00', '2024-03-05 23:59:59', '8-204');
INSERT INTO `courses` VALUES (2, 'gdxz003', '23', '计算机学院', '软件技术', '2班', '数据结构与算法', '2024-03-08 08:10:00', '2024-03-08 23:59:59', '8-204');
INSERT INTO `courses` VALUES (3, 'gdxz003', '22', '计算机学院', '软件技术', '4班', 'Web应用开发', '2024-03-04 08:00:00', '2024-03-04 23:59:59', '8-205');
INSERT INTO `courses` VALUES (4, 'gdxz003', '22', '计算机学院', '软件技术', '4班', '前端框架开发技术', '2024-03-06 08:10:00', '2024-03-06 23:59:59', '8-205');
INSERT INTO `courses` VALUES (5, 'gdxz003', '22', '计算机学院', '软件技术', '4班', '后端架构开发技术', '2024-03-07 08:10:00', '2024-03-07 23:59:59', '8-205');
INSERT INTO `courses` VALUES (6, 'gdxz003', '22', '计算机学院', '软件技术', '4班', '微信小程序', '2024-03-09 08:10:00', '2024-03-09 23:59:59', '8-205');
INSERT INTO `courses` VALUES (7, 'gdxz003', '22', '计算机学院', '软件技术', '4班', 'linux系统管理', '2024-03-10 00:10:00', '2024-03-10 23:59:59', '8-205');

-- ----------------------------
-- Table structure for majors
-- ----------------------------
DROP TABLE IF EXISTS `majors`;
CREATE TABLE `majors`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `college_id` int(0) NOT NULL,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of majors
-- ----------------------------
INSERT INTO `majors` VALUES (1, 1, '软件技术');
INSERT INTO `majors` VALUES (2, 1, '计算机网络技术');
INSERT INTO `majors` VALUES (3, 1, '物联网应用技术');
INSERT INTO `majors` VALUES (4, 1, '工业互联网技术');
INSERT INTO `majors` VALUES (5, 2, '国际经济与贸易');
INSERT INTO `majors` VALUES (6, 2, '市场营销');
INSERT INTO `majors` VALUES (7, 2, '商务管理');
INSERT INTO `majors` VALUES (8, 2, '会展策划与管理');
INSERT INTO `majors` VALUES (9, 2, '大数据与会计');
INSERT INTO `majors` VALUES (10, 2, '国际金融');
INSERT INTO `majors` VALUES (11, 3, '研学旅行管理与服务');
INSERT INTO `majors` VALUES (12, 3, '酒店管理与数字化运营');
INSERT INTO `majors` VALUES (13, 3, '公共文化服务与管理');
INSERT INTO `majors` VALUES (14, 4, '行政管理');
INSERT INTO `majors` VALUES (15, 4, '人力资源管理');
INSERT INTO `majors` VALUES (16, 4, '社区管理与服务');
INSERT INTO `majors` VALUES (17, 4, '青少年工作与管理');
INSERT INTO `majors` VALUES (18, 4, '法律事务');

-- ----------------------------
-- Table structure for records
-- ----------------------------
DROP TABLE IF EXISTS `records`;
CREATE TABLE `records`  (
  `id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `teacher_card` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `grade` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `college` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `major` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `class` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `course_id` int(0) NOT NULL,
  `student_card` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `create_time` datetime(0) NOT NULL,
  `update_time` datetime(0) NOT NULL,
  `status` int(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `index_name`(`student_card`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of records
-- ----------------------------
INSERT INTO `records` VALUES ('292918ec-f594-11ee-b2f2-f66add776e2d', 'gdxz003', '22', '计算机学院', '软件技术', '4班', 3, '22203020402', '陈锐斌', '2024-04-08 08:00:00', '2024-04-08 18:38:39', 2);
INSERT INTO `records` VALUES ('293af9c9-fba0-11ee-84b7-f66add776e2d', 'gdxz003', '23', '计算机学院', '软件技术', '2班', 1, '23203020201', '曾嘉豪', '2024-04-16 08:10:00', '2024-04-16 11:19:40', 1);
INSERT INTO `records` VALUES ('2a2483c5-fba0-11ee-84b7-f66add776e2d', 'gdxz003', '23', '计算机学院', '软件技术', '2班', 1, '23203020202', '陈栋宇', '2024-04-16 08:10:00', '2024-04-16 11:19:42', 1);
INSERT INTO `records` VALUES ('2a868279-fba0-11ee-84b7-f66add776e2d', 'gdxz003', '23', '计算机学院', '软件技术', '2班', 1, '23203020203', '陈麒先', '2024-04-16 08:10:00', '2024-04-16 11:19:42', 1);
INSERT INTO `records` VALUES ('2aeb06b2-fba0-11ee-84b7-f66add776e2d', 'gdxz003', '23', '计算机学院', '软件技术', '2班', 1, '23203020204', '陈晓敏', '2024-04-16 08:10:00', '2024-04-16 11:19:43', 1);
INSERT INTO `records` VALUES ('2d9c5101-fba0-11ee-84b7-f66add776e2d', 'gdxz003', '23', '计算机学院', '软件技术', '2班', 1, '23203020210', '郭毅岚', '2024-04-16 08:10:00', '2024-04-16 11:19:47', 2);
INSERT INTO `records` VALUES ('2dfc77db-fba0-11ee-84b7-f66add776e2d', 'gdxz003', '23', '计算机学院', '软件技术', '2班', 1, '23203020211', '韩宗攸', '2024-04-16 08:10:00', '2024-04-16 11:19:48', 2);
INSERT INTO `records` VALUES ('2e4a34d9-fba0-11ee-84b7-f66add776e2d', 'gdxz003', '23', '计算机学院', '软件技术', '2班', 1, '23203020212', '何铭棋', '2024-04-16 08:10:00', '2024-04-16 11:19:49', 2);
INSERT INTO `records` VALUES ('3549b3ca-f576-11ee-b5c4-f66add776e2d', 'gdxz003', '22', '计算机学院', '软件技术', '4班', 3, '22203020401', '陈广明', '2024-04-08 08:00:00', '2024-04-08 15:04:14', 1);
INSERT INTO `records` VALUES ('4702d712-f59a-11ee-bda9-f66add776e2d', 'gdxz003', '22', '计算机学院', '软件技术', '4班', 3, '22203020411', '黄宇轩', '2024-04-08 08:00:00', '2024-04-08 19:22:26', 2);
INSERT INTO `records` VALUES ('58d7f0ea-f6d4-11ee-8bdf-f66add776e2d', 'gdxz003', '22', '计算机学院', '软件技术', '4班', 4, '22203020401', '陈广明', '2024-04-10 08:10:00', '2024-04-10 08:50:38', 1);
INSERT INTO `records` VALUES ('5991a452-f6d4-11ee-8bdf-f66add776e2d', 'gdxz003', '22', '计算机学院', '软件技术', '4班', 4, '22203020403', '陈鑫', '2024-04-10 08:10:00', '2024-04-10 08:50:39', 2);
INSERT INTO `records` VALUES ('5a3ec573-f6d4-11ee-8bdf-f66add776e2d', 'gdxz003', '22', '计算机学院', '软件技术', '4班', 4, '22203020405', '戴苑莹', '2024-04-10 08:10:00', '2024-04-10 08:50:40', 1);
INSERT INTO `records` VALUES ('7d0da1dd-0895-11ef-9631-00155d70a558', '', '22', '计算机学院', '软件技术', '4班', 5, '22203020410', '黄爽清', '2024-05-02 08:10:00', '2024-05-02 23:06:01', 1);
INSERT INTO `records` VALUES ('92109322-09ec-11ef-a4b8-f66add776e2d', '', '22', '计算机学院', '软件技术', '4班', 6, '22203020410', '黄爽清', '2024-05-04 08:10:00', '2024-05-04 16:01:54', 1);
INSERT INTO `records` VALUES ('b301e1a4-fb3f-11ee-84b7-f66add776e2d', 'gdxz003', '22', '计算机学院', '软件技术', '4班', 3, '22203020401', '陈广明', '2024-04-15 08:00:00', '2024-04-15 23:49:10', 1);
INSERT INTO `records` VALUES ('b3bada16-fb3f-11ee-84b7-f66add776e2d', 'gdxz003', '22', '计算机学院', '软件技术', '4班', 3, '22203020403', '陈鑫', '2024-04-15 08:00:00', '2024-04-15 23:49:11', 1);
INSERT INTO `records` VALUES ('b471199e-fb3f-11ee-84b7-f66add776e2d', 'gdxz003', '22', '计算机学院', '软件技术', '4班', 3, '22203020405', '戴苑莹', '2024-04-15 08:00:00', '2024-04-15 23:49:12', 1);
INSERT INTO `records` VALUES ('b5d92f2a-fb3f-11ee-84b7-f66add776e2d', 'gdxz003', '22', '计算机学院', '软件技术', '4班', 3, '22203020410', '黄爽清', '2024-04-15 08:00:00', '2024-04-15 23:49:15', 1);
INSERT INTO `records` VALUES ('b5fe032b-0ce5-11ef-aa41-00155dea43e0', 'gdxz003', '22', '计算机学院', '软件技术', '4班', 4, '22203020410', '黄爽清', '2024-05-08 08:10:00', '2024-05-08 10:50:21', 1);

-- ----------------------------
-- Table structure for students
-- ----------------------------
DROP TABLE IF EXISTS `students`;
CREATE TABLE `students`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `grade` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `college` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `major` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `class` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `student_card` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `status` int(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `index_name`(`student_card`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 87 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of students
-- ----------------------------
INSERT INTO `students` VALUES (1, '22', '计算机学院', '软件技术', '4班', '22203020401', '陈广明', 1);
INSERT INTO `students` VALUES (2, '22', '计算机学院', '软件技术', '4班', '22203020402', '陈锐斌', 1);
INSERT INTO `students` VALUES (3, '22', '计算机学院', '软件技术', '4班', '22203020403', '陈鑫', 1);
INSERT INTO `students` VALUES (4, '22', '计算机学院', '软件技术', '4班', '22203020404', '陈子楼', 1);
INSERT INTO `students` VALUES (5, '22', '计算机学院', '软件技术', '4班', '22203020405', '戴苑莹', 1);
INSERT INTO `students` VALUES (6, '22', '计算机学院', '软件技术', '4班', '22203020406', '符杰森', 1);
INSERT INTO `students` VALUES (7, '22', '计算机学院', '软件技术', '4班', '22203020407', '傅康', 1);
INSERT INTO `students` VALUES (8, '22', '计算机学院', '软件技术', '4班', '22203020408', '洪晓辉', 1);
INSERT INTO `students` VALUES (9, '22', '计算机学院', '软件技术', '4班', '22203020409', '黄书成', 1);
INSERT INTO `students` VALUES (10, '22', '计算机学院', '软件技术', '4班', '22203020410', '黄爽清', 1);
INSERT INTO `students` VALUES (11, '22', '计算机学院', '软件技术', '4班', '22203020411', '黄宇轩', 1);
INSERT INTO `students` VALUES (12, '22', '计算机学院', '软件技术', '4班', '22203020412', '柯梓联', 1);
INSERT INTO `students` VALUES (13, '22', '计算机学院', '软件技术', '4班', '22203020413', '李焕城', 1);
INSERT INTO `students` VALUES (14, '22', '计算机学院', '软件技术', '4班', '22203020414', '李怡', 1);
INSERT INTO `students` VALUES (15, '22', '计算机学院', '软件技术', '4班', '22203020415', '练如风', 1);
INSERT INTO `students` VALUES (16, '22', '计算机学院', '软件技术', '4班', '22203020416', '林炯', 1);
INSERT INTO `students` VALUES (17, '22', '计算机学院', '软件技术', '4班', '22203020417', '林依烔', 1);
INSERT INTO `students` VALUES (18, '22', '计算机学院', '软件技术', '4班', '22203020418', '刘海峰', 1);
INSERT INTO `students` VALUES (19, '22', '计算机学院', '软件技术', '4班', '22203020419', '刘家炤', 1);
INSERT INTO `students` VALUES (20, '22', '计算机学院', '软件技术', '4班', '22203020420', '罗楚峰', 1);
INSERT INTO `students` VALUES (21, '22', '计算机学院', '软件技术', '4班', '22203020421', '陈奕冲', 1);
INSERT INTO `students` VALUES (22, '22', '计算机学院', '软件技术', '4班', '22203020422', '覃晓艳', 1);
INSERT INTO `students` VALUES (23, '22', '计算机学院', '软件技术', '4班', '22203020423', '王昊炫', 1);
INSERT INTO `students` VALUES (24, '22', '计算机学院', '软件技术', '4班', '22203020424', '王晓彤', 1);
INSERT INTO `students` VALUES (25, '22', '计算机学院', '软件技术', '4班', '22203020425', '王晔岑', 1);
INSERT INTO `students` VALUES (26, '22', '计算机学院', '软件技术', '4班', '22203020426', '吴凯槟', 1);
INSERT INTO `students` VALUES (27, '22', '计算机学院', '软件技术', '4班', '22203020427', '吴丽霞', 1);
INSERT INTO `students` VALUES (28, '22', '计算机学院', '软件技术', '4班', '22203020428', '吴烨栩', 1);
INSERT INTO `students` VALUES (29, '22', '计算机学院', '软件技术', '4班', '22203020429', '冼颖瑶', 1);
INSERT INTO `students` VALUES (30, '22', '计算机学院', '软件技术', '4班', '22203020430', '杨文杰', 1);
INSERT INTO `students` VALUES (31, '22', '计算机学院', '软件技术', '4班', '22203020431', '杨志炜', 1);
INSERT INTO `students` VALUES (32, '22', '计算机学院', '软件技术', '4班', '22203020432', '姚文康', 1);
INSERT INTO `students` VALUES (33, '22', '计算机学院', '软件技术', '4班', '22203020433', '易鹏', 1);
INSERT INTO `students` VALUES (34, '22', '计算机学院', '软件技术', '4班', '22203020434', '张丽婷', 1);
INSERT INTO `students` VALUES (35, '22', '计算机学院', '软件技术', '4班', '22203020435', '张荣姬', 1);
INSERT INTO `students` VALUES (36, '22', '计算机学院', '软件技术', '4班', '22203020437', '钟敏潼', 1);
INSERT INTO `students` VALUES (37, '22', '计算机学院', '软件技术', '4班', '22203020438', '钟烨', 1);
INSERT INTO `students` VALUES (38, '22', '计算机学院', '软件技术', '4班', '22203020439', '朱盼', 1);
INSERT INTO `students` VALUES (39, '22', '计算机学院', '软件技术', '4班', '22203020440', '庄如娜', 1);
INSERT INTO `students` VALUES (40, '23', '计算机学院', '软件技术', '2班', '23203020201', '曾嘉豪', 1);
INSERT INTO `students` VALUES (41, '23', '计算机学院', '软件技术', '2班', '23203020202', '陈栋宇', 1);
INSERT INTO `students` VALUES (42, '23', '计算机学院', '软件技术', '2班', '23203020203', '陈麒先', 1);
INSERT INTO `students` VALUES (43, '23', '计算机学院', '软件技术', '2班', '23203020204', '陈晓敏', 1);
INSERT INTO `students` VALUES (44, '23', '计算机学院', '软件技术', '2班', '23203020205', '陈泽煌', 1);
INSERT INTO `students` VALUES (45, '23', '计算机学院', '软件技术', '2班', '23203020206', '崔道铭', 1);
INSERT INTO `students` VALUES (46, '23', '计算机学院', '软件技术', '2班', '23203020207', '邓楚', 1);
INSERT INTO `students` VALUES (47, '23', '计算机学院', '软件技术', '2班', '23203020208', '翟泳辉', 1);
INSERT INTO `students` VALUES (48, '23', '计算机学院', '软件技术', '2班', '23203020209', '符饶议', 1);
INSERT INTO `students` VALUES (49, '23', '计算机学院', '软件技术', '2班', '23203020210', '郭毅岚', 1);
INSERT INTO `students` VALUES (50, '23', '计算机学院', '软件技术', '2班', '23203020211', '韩宗攸', 1);
INSERT INTO `students` VALUES (51, '23', '计算机学院', '软件技术', '2班', '23203020212', '何铭棋', 1);
INSERT INTO `students` VALUES (52, '23', '计算机学院', '软件技术', '2班', '23203020213', '何悦豪', 1);
INSERT INTO `students` VALUES (53, '23', '计算机学院', '软件技术', '2班', '23203020214', '黄志阳', 1);
INSERT INTO `students` VALUES (54, '23', '计算机学院', '软件技术', '2班', '23203020215', '江文焯', 1);
INSERT INTO `students` VALUES (55, '23', '计算机学院', '软件技术', '2班', '23203020216', '李冰娜', 1);
INSERT INTO `students` VALUES (56, '23', '计算机学院', '软件技术', '2班', '23203020217', '李贺颖', 1);
INSERT INTO `students` VALUES (57, '23', '计算机学院', '软件技术', '2班', '23203020218', '李天逸', 1);
INSERT INTO `students` VALUES (58, '23', '计算机学院', '软件技术', '2班', '23203020219', '李昀曦', 1);
INSERT INTO `students` VALUES (59, '23', '计算机学院', '软件技术', '2班', '23203020220', '梁玮琦', 1);
INSERT INTO `students` VALUES (60, '23', '计算机学院', '软件技术', '2班', '23203020221', '梁雨楠', 1);
INSERT INTO `students` VALUES (61, '23', '计算机学院', '软件技术', '2班', '23203020222', '廖桂凤', 1);
INSERT INTO `students` VALUES (62, '23', '计算机学院', '软件技术', '2班', '23203020223', '林诗琪', 1);
INSERT INTO `students` VALUES (63, '23', '计算机学院', '软件技术', '2班', '23203020224', '林晓柔', 1);
INSERT INTO `students` VALUES (64, '23', '计算机学院', '软件技术', '2班', '23203020225', '林雨琪', 1);
INSERT INTO `students` VALUES (65, '23', '计算机学院', '软件技术', '2班', '23203020226', '刘家欣', 1);
INSERT INTO `students` VALUES (66, '23', '计算机学院', '软件技术', '2班', '23203020227', '刘峻瑜', 1);
INSERT INTO `students` VALUES (67, '23', '计算机学院', '软件技术', '2班', '23203020228', '刘谢彤', 1);
INSERT INTO `students` VALUES (68, '23', '计算机学院', '软件技术', '2班', '23203020229', '吕伊晨', 1);
INSERT INTO `students` VALUES (69, '23', '计算机学院', '软件技术', '2班', '23203020230', '莫迷迷', 1);
INSERT INTO `students` VALUES (70, '23', '计算机学院', '软件技术', '2班', '23203020231', '潘文业', 1);
INSERT INTO `students` VALUES (71, '23', '计算机学院', '软件技术', '2班', '23203020232', '庞瑾彤', 1);
INSERT INTO `students` VALUES (72, '23', '计算机学院', '软件技术', '2班', '23203020233', '苏仁轩', 1);
INSERT INTO `students` VALUES (73, '23', '计算机学院', '软件技术', '2班', '23203020234', '孙家伟', 1);
INSERT INTO `students` VALUES (74, '23', '计算机学院', '软件技术', '2班', '23203020235', '唐杨杰', 1);
INSERT INTO `students` VALUES (75, '23', '计算机学院', '软件技术', '2班', '23203020236', '吴锐阳', 1);
INSERT INTO `students` VALUES (76, '23', '计算机学院', '软件技术', '2班', '23203020237', '谢可欣', 1);
INSERT INTO `students` VALUES (77, '23', '计算机学院', '软件技术', '2班', '23203020238', '谢绮文', 1);
INSERT INTO `students` VALUES (78, '23', '计算机学院', '软件技术', '2班', '23203020239', '徐梦娉', 1);
INSERT INTO `students` VALUES (79, '23', '计算机学院', '软件技术', '2班', '23203020240', '许冰婷', 1);
INSERT INTO `students` VALUES (80, '23', '计算机学院', '软件技术', '2班', '23203020241', '颜慧瑜', 1);
INSERT INTO `students` VALUES (81, '23', '计算机学院', '软件技术', '2班', '23203020242', '杨浩怡', 1);
INSERT INTO `students` VALUES (82, '23', '计算机学院', '软件技术', '2班', '23203020243', '袁海峰', 1);
INSERT INTO `students` VALUES (83, '23', '计算机学院', '软件技术', '2班', '23203020244', '郑琪媛', 1);
INSERT INTO `students` VALUES (84, '23', '计算机学院', '软件技术', '2班', '23203020245', '郑少纯', 1);
INSERT INTO `students` VALUES (85, '23', '计算机学院', '软件技术', '2班', '23203020246', '周楷婷', 1);
INSERT INTO `students` VALUES (86, '23', '计算机学院', '软件技术', '2班', '23203020247', '邹小惠', 1);

-- ----------------------------
-- Table structure for teachers
-- ----------------------------
DROP TABLE IF EXISTS `teachers`;
CREATE TABLE `teachers`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `college` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `major` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `teacher_card` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `status` int(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of teachers
-- ----------------------------
INSERT INTO `teachers` VALUES (8, '计算机学院', '软件技术', 'gdxz003', '黄爽清', '15767966159', '0784960d7a1f2bba9c3db3276def80f2', 1);

SET FOREIGN_KEY_CHECKS = 1;
