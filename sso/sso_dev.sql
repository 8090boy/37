/*
Navicat MySQL Data Transfer

Source Server         : loca
Source Server Version : 50710
Source Host           : localhost:3306
Source Database       : sso_dev

Target Server Type    : MYSQL
Target Server Version : 50710
File Encoding         : 65001

Date: 2016-01-31 17:44:56
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for state
-- ----------------------------
DROP TABLE IF EXISTS `state`;
CREATE TABLE `state` (
  `token` varchar(255) DEFAULT NULL,
  `userjson` varchar(255) DEFAULT NULL,
  `overdue` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of state
-- ----------------------------
INSERT INTO `state` VALUES ('2cf8b4a6-5ff7-487d-9200-fa8c15044446', '18620131416_3', '2016-01-31 23:58:58');
INSERT INTO `state` VALUES ('59816c49-cf1b-4905-b800-40a9c8758f02', '18620131416_3', '2016-01-31 23:59:27');
INSERT INTO `state` VALUES ('7047e95f-8de6-4790-b200-144613b6c96e', '18620131416_3', '2016-02-01 00:00:08');
INSERT INTO `state` VALUES ('e516127a-3ae7-4aec-96e1-af262fd7e0ad', '18620131416_3', '2016-02-01 00:12:30');
INSERT INTO `state` VALUES ('6fa2f4a3-6bce-4b4f-9780-f97e59dac20d', '18620131416_3', '2016-02-01 00:12:34');
INSERT INTO `state` VALUES ('0080f664-e9cc-41a6-bf43-b5d359cebe66', '18620131416_3', '2016-02-01 00:18:17');
INSERT INTO `state` VALUES ('df3b13d1-46b3-4ab8-84c5-c09a5cb5e515', '18620131416_3', '2016-02-01 21:10:09');
INSERT INTO `state` VALUES ('01cf6d4c-cb83-45f3-b7a5-e6c657f685e5', '18620131416_3', '2016-02-01 21:10:17');
INSERT INTO `state` VALUES ('739aea98-4210-4c8b-b96a-dc981105c7ed', '18620131416_3', '2016-02-01 21:10:18');
INSERT INTO `state` VALUES ('4898474c-1845-4dc5-ba24-89f393793e80', '18620131416_3', '2016-02-01 21:10:56');
INSERT INTO `state` VALUES ('04737e3c-399c-4375-80a8-9ad472902882', '18620131416_3', '2016-02-01 21:11:00');
INSERT INTO `state` VALUES ('dc88cecb-c6ba-44ea-a356-91c74c1a0e8f', '18620131416_3', '2016-02-01 21:11:11');
INSERT INTO `state` VALUES ('f94dd388-03e3-418a-989a-245521ae3b77', '18620131416_3', '2016-02-02 15:34:27');
INSERT INTO `state` VALUES ('88776525-29f0-4730-b998-53e6ae782e27', '18620131416_3', '2016-02-02 15:35:13');
INSERT INTO `state` VALUES ('51925db7-a1ab-4c5b-8ba6-94ec661427c2', '18620131416_3', '2016-02-02 15:35:16');
INSERT INTO `state` VALUES ('e97c6c06-8a43-4187-ac42-f6eeffa2c76a', '18620131416_3', '2016-02-02 16:35:32');
INSERT INTO `state` VALUES ('a53718f0-9572-4253-b664-35ba17f948ad', '13559938898_10', '2016-02-02 16:39:30');
INSERT INTO `state` VALUES ('e934e705-a755-47bc-b4c9-2146259afae5', '18620131416_3', '2016-02-02 16:51:17');
INSERT INTO `state` VALUES ('076ed883-474e-4afe-b835-779f0aab6b1e', '18620131416_3', '2016-02-02 16:57:06');
INSERT INTO `state` VALUES ('7536c9d1-865f-4d06-937a-f741aa40d0f1', '18620131416_3', '2016-02-02 17:11:17');
INSERT INTO `state` VALUES ('fdc8271b-5358-4037-b3e9-d174759ec625', '13559938898_10', '2016-02-02 17:16:48');
INSERT INTO `state` VALUES ('78568b20-308b-424b-ab97-8b866a73baa8', '18620131416_3', '2016-02-02 17:45:11');
INSERT INTO `state` VALUES ('49aab234-747b-46ee-a4e0-89dc766c0bf6', '18620131416_3', '2016-02-02 17:46:10');
INSERT INTO `state` VALUES ('083054d3-f749-41bc-8912-cb42440b530a', '18620131416_3', '2016-02-02 18:09:41');
INSERT INTO `state` VALUES ('b8d6c1e6-e87a-4ac8-8286-313c412b4e17', '13355558888_12', '2016-02-02 18:10:21');
INSERT INTO `state` VALUES ('0885a970-5077-46fc-b0a7-33773c111092', '18620131416_3', '2016-02-02 21:49:55');
INSERT INTO `state` VALUES ('265c436a-a7d7-49f0-b983-8c2408e6eeac', '18620131416_3', '2016-02-02 22:20:08');
INSERT INTO `state` VALUES ('c1a66294-825e-4580-8a3b-ef5e986cd5e8', '18620131416_3', '2016-02-02 23:32:04');
INSERT INTO `state` VALUES ('0b1314ee-e1d0-4afb-b877-49b56d9d0e96', '18620131416_3', '2016-02-02 23:36:02');
INSERT INTO `state` VALUES ('9d176448-7fd5-4f52-ab61-9aec945521a3', '18620131416_3', '2016-02-03 13:02:44');
INSERT INTO `state` VALUES ('656f3018-4d7c-4f57-9114-9ebd0b2e04ab', '18620131416_3', '2016-02-03 13:03:01');
INSERT INTO `state` VALUES ('68a9398a-0f53-4ba8-bb1d-d00b2bb481f0', '18620131416_3', '2016-02-03 13:04:19');
INSERT INTO `state` VALUES ('090c0639-8763-45e2-8fa7-b8286fca3b68', '18620131416_3', '2016-02-03 13:05:25');
INSERT INTO `state` VALUES ('ff6c582d-fe32-42cf-bf94-5d92d914ef41', '13888888888_8', '2016-02-03 13:12:51');
INSERT INTO `state` VALUES ('40f92281-e0b2-4724-af84-b3f9a6617dd9', '18620131417_4', '2016-02-03 14:54:43');
INSERT INTO `state` VALUES ('dad86240-b9c9-4b7a-a8e3-919507627524', '18620131419_6', '2016-02-03 15:16:25');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `q_q` int(11) DEFAULT NULL,
  `sex` tinyint(1) DEFAULT NULL,
  `username` varchar(20) DEFAULT NULL,
  `password` varchar(8) DEFAULT NULL,
  `alias` varchar(50) DEFAULT NULL,
  `mobile` varchar(11) DEFAULT NULL,
  `alipay` varchar(50) DEFAULT NULL,
  `wechat` varchar(50) DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL,
  `city` varchar(20) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `identity` varchar(20) DEFAULT NULL,
  `create` datetime DEFAULT NULL,
  `last` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', '611041314', '1', 'roc', '0b14728a', '2013', '18620131415', 'alipay@qq.com', 'Wechat001', '611041314@qq.com', '汉中', '坂田街道岗塔偶', '612325101122', '2016-01-31 13:51:04', '2016-01-31 13:51:04');
INSERT INTO `user` VALUES ('2', '611041314', '1', '0099', '0b14728a', '0099', '13678560099', 'alipay@qq.com', 'Wechat001', '13678560099@qq.com', '汉中', '坂田街道岗塔偶', '1312312313', '2016-01-31 13:51:04', '2016-01-31 13:51:04');
INSERT INTO `user` VALUES ('3', '611041314', '1', '1416', '0b14728a', '1416', '18620131416', 'alipay@qq.com', 'Wechat001', '1416@qq.com', '汉中', '坂田街道岗塔偶', '1312312313', '2016-01-31 13:51:04', '2016-01-31 13:51:04');
INSERT INTO `user` VALUES ('4', '0', '0', '', '0b14728a', '', '18620131417', '', 'abcd_1234', '', '', '', '', '2016-01-31 13:51:59', '2016-01-31 13:51:59');
INSERT INTO `user` VALUES ('5', '0', '0', '', '0b14728a', '', '18620131418', '', 'abcd_1234', '', '', '', '', '2016-01-31 13:56:55', '2016-01-31 13:56:55');
INSERT INTO `user` VALUES ('6', '0', '0', '', '0b14728a', '', '18620131419', '', 'abcd_1234', '', '', '', '', '2016-01-31 13:59:15', '2016-01-31 13:59:15');
