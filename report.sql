-- MySQL dump 10.13  Distrib 5.7.18, for Linux (x86_64)
--
-- Host: localhost    Database: smart_backstage
-- ------------------------------------------------------
-- Server version	5.7.18-0ubuntu0.16.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `report`
--

DROP TABLE IF EXISTS `report`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `report` (
  `id` varchar(100) NOT NULL,
  `pasin` varchar(100) DEFAULT NULL COMMENT 'Asin(父)',
  `asin` varchar(100) DEFAULT NULL COMMENT 'Asin(子)',
  `title` varchar(100) DEFAULT NULL COMMENT '商品名称',
  `uv` int(11) DEFAULT NULL COMMENT '买家访问次数',
  `uvb` varchar(100) DEFAULT NULL COMMENT '该日买家访问次数百分比',
  `pv` int(11) DEFAULT NULL COMMENT '页面浏览次数',
  `pvb` varchar(100) DEFAULT NULL COMMENT '该日页面浏览次数百分比',
  `bpvb` varchar(100) DEFAULT NULL COMMENT '该日购物车占比',
  `on` int(11) DEFAULT NULL COMMENT '该日已订购商品数量',
  `onr` varchar(100) DEFAULT NULL COMMENT '该日订单商品数量转化率',
  `v` double DEFAULT NULL COMMENT '该日已订购商品销售额',
  `c` int(11) DEFAULT NULL COMMENT '该日订单数',
  `d` varchar(100) DEFAULT NULL COMMENT '日期',
  `aws` varchar(100) DEFAULT NULL COMMENT '店铺名',
  `status` tinyint(4) DEFAULT NULL COMMENT 'Asin(父)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `report`
--

LOCK TABLES `report` WRITE;
/*!40000 ALTER TABLE `report` DISABLE KEYS */;
/*!40000 ALTER TABLE `report` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-07-16 14:15:37
