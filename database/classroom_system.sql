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

 Date: 11/03/2024 21:58:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP Database IF EXISTS `classroom_system`;
CREATE Database IF NOT EXISTS `classroom_system` DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

USE `classroom_system`;

-- ----------------------------
-- Table structure for classes
-- ----------------------------
DROP TABLE IF EXISTS `classes`;
CREATE TABLE `classes`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `class` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of classes
-- ----------------------------
INSERT INTO `classes` VALUES (1, '1班');
INSERT INTO `classes` VALUES (2, '2班');
INSERT INTO `classes` VALUES (3, '3班');
INSERT INTO `classes` VALUES (4, '4班');

-- ----------------------------
-- Table structure for classrooms
-- ----------------------------
DROP TABLE IF EXISTS `classrooms`;
CREATE TABLE `classrooms`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `classroom` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of classrooms
-- ----------------------------
INSERT INTO `classrooms` VALUES (1, '3-403');
INSERT INTO `classrooms` VALUES (2, '8-205');
INSERT INTO `classrooms` VALUES (3, '8-204');
INSERT INTO `classrooms` VALUES (4, '9-604');
INSERT INTO `classrooms` VALUES (5, '3-301');
INSERT INTO `classrooms` VALUES (6, '2-302');
INSERT INTO `classrooms` VALUES (7, '2-101');
INSERT INTO `classrooms` VALUES (8, '4-202');
INSERT INTO `classrooms` VALUES (9, '3-501');
INSERT INTO `classrooms` VALUES (10, '8-206');
INSERT INTO `classrooms` VALUES (11, '2-105');
INSERT INTO `classrooms` VALUES (12, '2-301');

-- ----------------------------
-- Table structure for colleges
-- ----------------------------
DROP TABLE IF EXISTS `colleges`;
CREATE TABLE `colleges`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `college` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `campus` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

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
  `teacher_id` int(0) NOT NULL,
  `grade_id` int(0) NOT NULL,
  `major_id` int(0) NOT NULL,
  `classe_id` int(0) NOT NULL,
  `course` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `start_time` datetime(0) NULL DEFAULT NULL,
  `end_time` datetime(0) NULL DEFAULT NULL,
  `classroom_id` int(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of courses
-- ----------------------------
INSERT INTO `courses` VALUES (1, 1, 2, 1, 2, 'Java程序设计', '2024-03-05 08:10:00', '2024-03-05 23:59:59', 3);
INSERT INTO `courses` VALUES (2, 1, 2, 1, 2, '数据结构与算法', '2024-03-08 08:10:00', '2024-03-08 23:59:59', 3);
INSERT INTO `courses` VALUES (3, 1, 1, 1, 4, 'Web应用开发', '2024-03-04 08:00:00', '2024-03-04 23:59:59', 2);
INSERT INTO `courses` VALUES (4, 1, 1, 1, 4, '前端框架开发技术', '2024-03-06 08:10:00', '2024-03-06 23:59:59', 2);
INSERT INTO `courses` VALUES (5, 1, 1, 1, 4, '后端架构开发技术', '2024-03-07 08:10:00', '2024-03-07 23:59:59', 2);
INSERT INTO `courses` VALUES (6, 1, 1, 1, 4, '微信小程序', '2024-03-09 08:10:00', '2024-03-09 23:59:59', 3);
INSERT INTO `courses` VALUES (7, 1, 1, 1, 4, 'linux系统管理', '2024-03-10 08:10:00', '2024-03-10 23:59:59', 3);

-- ----------------------------
-- Table structure for grades
-- ----------------------------
DROP TABLE IF EXISTS `grades`;
CREATE TABLE `grades`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `grade` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of grades
-- ----------------------------
INSERT INTO `grades` VALUES (1, '22');
INSERT INTO `grades` VALUES (2, '23');

-- ----------------------------
-- Table structure for majors
-- ----------------------------
DROP TABLE IF EXISTS `majors`;
CREATE TABLE `majors`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `college_id` int(0) NOT NULL,
  `major` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

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
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `uuid` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `course_id` int(0) NOT NULL,
  `students_id` int(0) NOT NULL,
  `create_time` datetime(0) NOT NULL,
  `update_time` datetime(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of records
-- ----------------------------

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of students
-- ----------------------------
INSERT INTO `students` VALUES (1, '22', '计算机学院', '软件技术', '4班', '22203020401', '陈广明');
INSERT INTO `students` VALUES (2, '22', '计算机学院', '软件技术', '4班', '22203020402', '陈锐斌');
INSERT INTO `students` VALUES (3, '22', '计算机学院', '软件技术', '4班', '22203020403', '陈鑫');
INSERT INTO `students` VALUES (4, '22', '计算机学院', '软件技术', '4班', '22203020404', '陈子楼');
INSERT INTO `students` VALUES (5, '22', '计算机学院', '软件技术', '4班', '22203020405', '戴苑莹');
INSERT INTO `students` VALUES (6, '22', '计算机学院', '软件技术', '4班', '22203020406', '符杰森');
INSERT INTO `students` VALUES (7, '22', '计算机学院', '软件技术', '4班', '22203020407', '傅康');
INSERT INTO `students` VALUES (8, '22', '计算机学院', '软件技术', '4班', '22203020408', '洪晓辉');
INSERT INTO `students` VALUES (9, '22', '计算机学院', '软件技术', '4班', '22203020409', '黄书成');
INSERT INTO `students` VALUES (10, '22', '计算机学院', '软件技术', '4班', '22203020410', '黄爽清');
INSERT INTO `students` VALUES (11, '22', '计算机学院', '软件技术', '4班', '22203020411', '黄宇轩');
INSERT INTO `students` VALUES (12, '22', '计算机学院', '软件技术', '4班', '22203020412', '柯梓联');
INSERT INTO `students` VALUES (13, '22', '计算机学院', '软件技术', '4班', '22203020413', '李焕城');
INSERT INTO `students` VALUES (14, '22', '计算机学院', '软件技术', '4班', '22203020414', '李怡');
INSERT INTO `students` VALUES (15, '22', '计算机学院', '软件技术', '4班', '22203020415', '练如风');
INSERT INTO `students` VALUES (16, '22', '计算机学院', '软件技术', '4班', '22203020416', '林炯');
INSERT INTO `students` VALUES (17, '22', '计算机学院', '软件技术', '4班', '22203020417', '林依烔');
INSERT INTO `students` VALUES (18, '22', '计算机学院', '软件技术', '4班', '22203020418', '刘海峰');
INSERT INTO `students` VALUES (19, '22', '计算机学院', '软件技术', '4班', '22203020419', '刘家炤');
INSERT INTO `students` VALUES (20, '22', '计算机学院', '软件技术', '4班', '22203020420', '罗楚峰');
INSERT INTO `students` VALUES (21, '22', '计算机学院', '软件技术', '4班', '22203020421', '陈奕冲');
INSERT INTO `students` VALUES (22, '22', '计算机学院', '软件技术', '4班', '22203020422', '覃晓艳');
INSERT INTO `students` VALUES (23, '22', '计算机学院', '软件技术', '4班', '22203020423', '王昊炫');
INSERT INTO `students` VALUES (24, '22', '计算机学院', '软件技术', '4班', '22203020424', '王晓彤');
INSERT INTO `students` VALUES (25, '22', '计算机学院', '软件技术', '4班', '22203020425', '王晔岑');
INSERT INTO `students` VALUES (26, '22', '计算机学院', '软件技术', '4班', '22203020426', '吴凯槟');
INSERT INTO `students` VALUES (27, '22', '计算机学院', '软件技术', '4班', '22203020427', '吴丽霞');
INSERT INTO `students` VALUES (28, '22', '计算机学院', '软件技术', '4班', '22203020428', '吴烨栩');
INSERT INTO `students` VALUES (29, '22', '计算机学院', '软件技术', '4班', '22203020429', '冼颖瑶');
INSERT INTO `students` VALUES (30, '22', '计算机学院', '软件技术', '4班', '22203020430', '杨文杰');
INSERT INTO `students` VALUES (31, '22', '计算机学院', '软件技术', '4班', '22203020431', '杨志炜');
INSERT INTO `students` VALUES (32, '22', '计算机学院', '软件技术', '4班', '22203020432', '姚文康');
INSERT INTO `students` VALUES (33, '22', '计算机学院', '软件技术', '4班', '22203020433', '易鹏');
INSERT INTO `students` VALUES (34, '22', '计算机学院', '软件技术', '4班', '22203020434', '张丽婷');
INSERT INTO `students` VALUES (35, '22', '计算机学院', '软件技术', '4班', '22203020435', '张荣姬');
INSERT INTO `students` VALUES (36, '22', '计算机学院', '软件技术', '4班', '22203020437', '钟敏潼');
INSERT INTO `students` VALUES (37, '22', '计算机学院', '软件技术', '4班', '22203020438', '钟烨');
INSERT INTO `students` VALUES (38, '22', '计算机学院', '软件技术', '4班', '22203020439', '朱盼');
INSERT INTO `students` VALUES (39, '22', '计算机学院', '软件技术', '4班', '22203020440', '庄如娜');
INSERT INTO `students` VALUES (40, '23', '计算机学院', '软件技术', '2班', '23203020201', '曾嘉豪');
INSERT INTO `students` VALUES (41, '23', '计算机学院', '软件技术', '2班', '23203020202', '陈栋宇');
INSERT INTO `students` VALUES (42, '23', '计算机学院', '软件技术', '2班', '23203020203', '陈麒先');
INSERT INTO `students` VALUES (43, '23', '计算机学院', '软件技术', '2班', '23203020204', '陈晓敏');
INSERT INTO `students` VALUES (44, '23', '计算机学院', '软件技术', '2班', '23203020205', '陈泽煌');
INSERT INTO `students` VALUES (45, '23', '计算机学院', '软件技术', '2班', '23203020206', '崔道铭');
INSERT INTO `students` VALUES (46, '23', '计算机学院', '软件技术', '2班', '23203020207', '邓楚');
INSERT INTO `students` VALUES (47, '23', '计算机学院', '软件技术', '2班', '23203020208', '翟泳辉');
INSERT INTO `students` VALUES (48, '23', '计算机学院', '软件技术', '2班', '23203020209', '符饶议');
INSERT INTO `students` VALUES (49, '23', '计算机学院', '软件技术', '2班', '23203020210', '郭毅岚');
INSERT INTO `students` VALUES (50, '23', '计算机学院', '软件技术', '2班', '23203020211', '韩宗攸');
INSERT INTO `students` VALUES (51, '23', '计算机学院', '软件技术', '2班', '23203020212', '何铭棋');
INSERT INTO `students` VALUES (52, '23', '计算机学院', '软件技术', '2班', '23203020213', '何悦豪');
INSERT INTO `students` VALUES (53, '23', '计算机学院', '软件技术', '2班', '23203020214', '黄志阳');
INSERT INTO `students` VALUES (54, '23', '计算机学院', '软件技术', '2班', '23203020215', '江文焯');
INSERT INTO `students` VALUES (55, '23', '计算机学院', '软件技术', '2班', '23203020216', '李冰娜');
INSERT INTO `students` VALUES (56, '23', '计算机学院', '软件技术', '2班', '23203020217', '李贺颖');
INSERT INTO `students` VALUES (57, '23', '计算机学院', '软件技术', '2班', '23203020218', '李天逸');
INSERT INTO `students` VALUES (58, '23', '计算机学院', '软件技术', '2班', '23203020219', '李昀曦');
INSERT INTO `students` VALUES (59, '23', '计算机学院', '软件技术', '2班', '23203020220', '梁玮琦');
INSERT INTO `students` VALUES (60, '23', '计算机学院', '软件技术', '2班', '23203020221', '梁雨楠');
INSERT INTO `students` VALUES (61, '23', '计算机学院', '软件技术', '2班', '23203020222', '廖桂凤');
INSERT INTO `students` VALUES (62, '23', '计算机学院', '软件技术', '2班', '23203020223', '林诗琪');
INSERT INTO `students` VALUES (63, '23', '计算机学院', '软件技术', '2班', '23203020224', '林晓柔');
INSERT INTO `students` VALUES (64, '23', '计算机学院', '软件技术', '2班', '23203020225', '林雨琪');
INSERT INTO `students` VALUES (65, '23', '计算机学院', '软件技术', '2班', '23203020226', '刘家欣');
INSERT INTO `students` VALUES (66, '23', '计算机学院', '软件技术', '2班', '23203020227', '刘峻瑜');
INSERT INTO `students` VALUES (67, '23', '计算机学院', '软件技术', '2班', '23203020228', '刘谢彤');
INSERT INTO `students` VALUES (68, '23', '计算机学院', '软件技术', '2班', '23203020229', '吕伊晨');
INSERT INTO `students` VALUES (69, '23', '计算机学院', '软件技术', '2班', '23203020230', '莫迷迷');
INSERT INTO `students` VALUES (70, '23', '计算机学院', '软件技术', '2班', '23203020231', '潘文业');
INSERT INTO `students` VALUES (71, '23', '计算机学院', '软件技术', '2班', '23203020232', '庞瑾彤');
INSERT INTO `students` VALUES (72, '23', '计算机学院', '软件技术', '2班', '23203020233', '苏仁轩');
INSERT INTO `students` VALUES (73, '23', '计算机学院', '软件技术', '2班', '23203020234', '孙家伟');
INSERT INTO `students` VALUES (74, '23', '计算机学院', '软件技术', '2班', '23203020235', '唐杨杰');
INSERT INTO `students` VALUES (75, '23', '计算机学院', '软件技术', '2班', '23203020236', '吴锐阳');
INSERT INTO `students` VALUES (76, '23', '计算机学院', '软件技术', '2班', '23203020237', '谢可欣');
INSERT INTO `students` VALUES (77, '23', '计算机学院', '软件技术', '2班', '23203020238', '谢绮文');
INSERT INTO `students` VALUES (78, '23', '计算机学院', '软件技术', '2班', '23203020239', '徐梦娉');
INSERT INTO `students` VALUES (79, '23', '计算机学院', '软件技术', '2班', '23203020240', '许冰婷');
INSERT INTO `students` VALUES (80, '23', '计算机学院', '软件技术', '2班', '23203020241', '颜慧瑜');
INSERT INTO `students` VALUES (81, '23', '计算机学院', '软件技术', '2班', '23203020242', '杨浩怡');
INSERT INTO `students` VALUES (82, '23', '计算机学院', '软件技术', '2班', '23203020243', '袁海峰');
INSERT INTO `students` VALUES (83, '23', '计算机学院', '软件技术', '2班', '23203020244', '郑琪媛');
INSERT INTO `students` VALUES (84, '23', '计算机学院', '软件技术', '2班', '23203020245', '郑少纯');
INSERT INTO `students` VALUES (85, '23', '计算机学院', '软件技术', '2班', '23203020246', '周楷婷');
INSERT INTO `students` VALUES (86, '23', '计算机学院', '软件技术', '2班', '23203020247', '邹小惠');

-- ----------------------------
-- Table structure for teachers
-- ----------------------------
DROP TABLE IF EXISTS `teachers`;
CREATE TABLE `teachers`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `teacher_card` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of teachers
-- ----------------------------
INSERT INTO `teachers` VALUES (1, 'gdxz003', "12345", '黄爽清', NULL);

SET FOREIGN_KEY_CHECKS = 1;
