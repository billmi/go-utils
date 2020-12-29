/*
Navicat MySQL Data Transfer

Source Server         : 127.0.0.1
Source Server Version : 50553
Source Host           : localhost:3306
Source Database       : bill

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2019-04-24 14:57:36
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `test`
-- ----------------------------
DROP TABLE IF EXISTS `test`;
CREATE TABLE `test` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `device_type` int(11) NOT NULL DEFAULT '0' COMMENT '设备类型：1-',
  `device_type_name` varchar(64) NOT NULL,
  `device_type_title` varchar(64) NOT NULL DEFAULT '' COMMENT '设备类型标题',
  `r_id` int(11) NOT NULL,
  `device_screen_width` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '屏宽',
  `device_screen_height` int(11) NOT NULL DEFAULT '0' COMMENT '屏高',
  `update_time` int(11) NOT NULL DEFAULT '0',
  `create_time` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `i_device_type` (`device_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COMMENT='设备类型表';

-- ----------------------------
-- Records of test
-- ----------------------------
INSERT INTO `test` VALUES ('1', '1', 'aa', 'aa', '1', '1920', '1080', '1541647758', '1541646082');
INSERT INTO `test` VALUES ('2', '2', 'bb', 'bb', '2', '1920', '1080', '1541746764', '1541746764');
INSERT INTO `test` VALUES ('3', '3', 'cc', 'cc', '3', '1920', '1080', '1541753141', '1541749904');
INSERT INTO `test` VALUES ('4', '4', 'dd', 'dd', '4', '1080', '1920', '1541753141', '1541753141');
INSERT INTO `test` VALUES ('5', '5', 'ee', 'ee', '5', '1920', '1080', '1541753141', '1541753141');
INSERT INTO `test` VALUES ('6', '6', 'ff', 'ff', '6', '1920', '360', '1541753141', '1541753141');
INSERT INTO `test` VALUES ('7', '7', 'gg', 'gg', '7', '800', '1280', '1541753141', '1541753141');
INSERT INTO `test` VALUES ('8', '8', 'hh', 'hh', '8', '1080', '1920', '1541753141', '1541753141');
INSERT INTO `test` VALUES ('10', '10', 'ii', 'ii', '9', '1920', '1080', '1541753141', '1541753141');

-- ----------------------------
-- Table structure for `test1`
-- ----------------------------
DROP TABLE IF EXISTS `test1`;
CREATE TABLE `test1` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(50) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `create_time` int(11) DEFAULT NULL,
  `update_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of test1
-- ----------------------------
INSERT INTO `test1` VALUES ('1', 'relation1', '1', '0', '0');
INSERT INTO `test1` VALUES ('2', 'relation2', '1', '0', '0');
INSERT INTO `test1` VALUES ('3', 'relation3', '1', '0', '0');
INSERT INTO `test1` VALUES ('4', 'relation4', '1', '0', '0');
INSERT INTO `test1` VALUES ('5', 'relation5', '1', '0', '0');
INSERT INTO `test1` VALUES ('6', 'relation6', '1', '0', '0');
INSERT INTO `test1` VALUES ('7', 'relation7', '1', '0', '0');
INSERT INTO `test1` VALUES ('8', 'relation8', '1', '0', '0');
