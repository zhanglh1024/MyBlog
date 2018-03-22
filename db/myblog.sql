/*
Navicat MySQL Data Transfer

Source Server         : Myblog
Source Server Version : 50720
Source Host           : localhost:3306
Source Database       : myblog

Target Server Type    : MYSQL
Target Server Version : 50720
File Encoding         : 65001

Date: 2018-03-22 17:26:43
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Uid` int(11) DEFAULT NULL,
  `Content` text,
  `Attachment` varchar(255) DEFAULT NULL,
  `Created` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `Auther` varchar(255) DEFAULT NULL,
  `ReplayCount` int(11) DEFAULT '0',
  `ReplayId` int(11) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of message
-- ----------------------------
INSERT INTO `message` VALUES ('1', '1', '如何才能把网站的前段做的漂亮点', null, null, 'zlh', '0', null);
INSERT INTO `message` VALUES ('2', '0', '事在人为', '', '2018-03-22 17:20:51', 'zlhuuu', '0', '0');
INSERT INTO `message` VALUES ('3', '0', '简单来说，设置 core.autocrlf=true 后，我们工作区的文件都应该用 CRLF 来换行。如果改动文件时引入了 LF，或者设置 core.autocrlf 之前，工作区已经有 LF 换行符。提交改动时，git 会警告你哪些文件不是纯 CRLF 文件，但 git 不会擅自修改工作区的那些文件，而是对暂存区（我们对工作区的改动）进行修改。也因此，当我们进行 git add 的操作时，只要 git 发现改动的内容里有 LF 换行符，就还会出现这个警告。\r\n\r\n作者：Andy Deng\r\n链接：https://www.zhihu.com/question/50862500/answer/123197258\r\n来源：知乎\r\n著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。', '', '2018-03-22 17:24:37', 'zlh', '0', '0');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Username` varchar(255) NOT NULL COMMENT '用户名',
  `Pwd` varchar(11) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', 'zlh', '262');
INSERT INTO `user` VALUES ('2', 'desire', '3333');
INSERT INTO `user` VALUES ('3', 'dd', 'ddd');
INSERT INTO `user` VALUES ('4', 'tianaishang', 'zlh');
INSERT INTO `user` VALUES ('5', 'zlhuuu', '262');
