/*
Navicat MySQL Data Transfer

Source Server         : Myblog
Source Server Version : 50720
Source Host           : localhost:3306
Source Database       : myblog

Target Server Type    : MYSQL
Target Server Version : 50720
File Encoding         : 65001

Date: 2018-03-23 18:01:57
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for artical
-- ----------------------------
DROP TABLE IF EXISTS `artical`;
CREATE TABLE `artical` (
  `Id` bigint(11) NOT NULL AUTO_INCREMENT,
  `Title` char(255) DEFAULT NULL COMMENT '用户文章标题',
  `Auther` char(20) CHARACTER SET utf8 COLLATE utf8_latvian_ci DEFAULT NULL COMMENT '文章作者',
  `Content` text COMMENT '文章内容',
  `WritDate` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of artical
-- ----------------------------
INSERT INTO `artical` VALUES ('1', '古代凶兽', 'copy', '狴犴，又名宪章，样子像虎，有威力，好狱讼，却又有威力，人们便将其刻铸在监狱门上，故民间有虎头牢的说法。狱门上部那虎头形的装饰便是其遗像。又相传它主持正义，能明是非，秉公而断，再加上它的形象威风凛凛,因此它也被安在衙门大堂两则以及官员出巡回避的牌上端，以维护公堂的肃然之气。', '2018-03-23 17:51:06');
INSERT INTO `artical` VALUES ('2', '睚眦', 'break', '睚眦，平生好斗喜杀，刀环、刀柄、龙吞口便是它的遗像。这些武器装饰了龙的形象后，更增添了慑人的力量。它不仅装饰在沙场名将的兵器上，更大量地用在仪仗和宫殿守卫者武器上，从而更显得威严庄重。', null);

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message` (
  `Id` bigint(11) NOT NULL AUTO_INCREMENT,
  `Uid` bigint(11) DEFAULT NULL,
  `Content` text,
  `Attachment` varchar(255) DEFAULT NULL,
  `Created` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `Auther` varchar(255) DEFAULT NULL,
  `ReplayCount` bigint(11) DEFAULT '0',
  `ReplayId` bigint(11) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of message
-- ----------------------------
INSERT INTO `message` VALUES ('1', '1', '如何才能把网站的前段做的漂亮点', null, null, 'zlh', '0', null);
INSERT INTO `message` VALUES ('2', '0', '事在人为', '', '2018-03-22 17:20:51', 'zlhuuu', '0', '0');
INSERT INTO `message` VALUES ('3', '0', '简单来说，设置 core.autocrlf=true 后，我们工作区的文件都应该用 CRLF 来换行。如果改动文件时引入了 LF，或者设置 core.autocrlf 之前，工作区已经有 LF 换行符。提交改动时，git 会警告你哪些文件不是纯 CRLF 文件，但 git 不会擅自修改工作区的那些文件，而是对暂存区（我们对工作区的改动）进行修改。也因此，当我们进行 git add 的操作时，只要 git 发现改动的内容里有 LF 换行符，就还会出现这个警告。\r\n\r\n作者：Andy Deng\r\n链接：https://www.zhihu.com/question/50862500/answer/123197258\r\n来源：知乎\r\n著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。', '', '2018-03-22 17:24:37', 'zlh', '0', '0');
INSERT INTO `message` VALUES ('4', '0', '好黑哈哈\r\n', '', '2018-03-23 15:42:59', '浪里淘沙', '0', '0');
INSERT INTO `message` VALUES ('5', '0', '打工的法国大革', '', '2018-03-23 15:46:05', 'dsfdsf', '0', '0');
INSERT INTO `message` VALUES ('6', '0', '打工的法国大革vv', '', '2018-03-23 15:48:21', 'dsfdsf', '0', '0');
INSERT INTO `message` VALUES ('7', '0', '打工的法国大革vv', '', '2018-03-23 15:48:33', 'dsfdsf', '0', '0');
INSERT INTO `message` VALUES ('8', '0', '神烦狗单调', '', '2018-03-23 15:54:28', '多少发给', '0', '0');
INSERT INTO `message` VALUES ('9', '0', '大是大非的是多少多少', '', '2018-03-23 16:01:02', 'hello', '0', '0');
INSERT INTO `message` VALUES ('10', '0', '斯蒂芬斯蒂芬我 ', '', '2018-03-23 16:38:39', '的师傅的师傅', '0', '0');
INSERT INTO `message` VALUES ('11', '0', '人团圆46有人提议', '', '2018-03-23 16:45:27', '让他也太人员', '0', '0');
INSERT INTO `message` VALUES ('12', '0', '人团圆46有人提议', '', '2018-03-23 16:53:36', '让他也太人员', '0', '0');
INSERT INTO `message` VALUES ('13', '0', '说的风格的大风刮大风刮东风公司', '', '2018-03-23 16:53:47', '说的非官方的', '0', '0');

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
