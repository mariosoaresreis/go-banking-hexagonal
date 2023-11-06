-- MySQL dump 10.13  Distrib 5.6.27, for Linux (x86_64)
--
-- Host: localhost    Database: banking
-- ------------------------------------------------------
-- Server version	5.6.27

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
-- Table structure for table `accounts`
--
CREATE DATABASE banking;
USE banking;

DROP TABLE IF EXISTS clientes;
CREATE TABLE clientes (
                             `id` int(11) NOT NULL AUTO_INCREMENT,
                             `nome` varchar(100) NOT NULL,
                             `data_nascimento` date NOT NULL,
                             `cidade` varchar(100) NOT NULL,
                             `cep` varchar(10) NOT NULL,
                             `status` tinyint(1) NOT NULL DEFAULT '1',
                             PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=2006 DEFAULT CHARSET=latin1;

LOCK TABLES clientes WRITE;
/*!40000 ALTER TABLE clientes DISABLE KEYS */;
INSERT INTO clientes VALUES
                            (2000,'Steve','1978-12-15','Delhi','110075',1),
                            (2001,'Arian','1988-05-21','Newburgh, NY','12550',1),
                            (2002,'Hadley','1988-04-30','Englewood, NJ','07631',1),
                            (2003,'Ben','1988-01-04','Manchester, NH','03102',0),
                            (2004,'Nina','1988-05-14','Clarkston, MI','48348',1),
                            (2005,'Osman','1988-11-08','Hyattsville, MD','20782',0);


UNLOCK TABLES;

DROP TABLE IF EXISTS contas_bancarias;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE contas_bancarias (
                            `id` int(11) NOT NULL AUTO_INCREMENT,
                            `cliente_id` int(11) NOT NULL,
                            `data_abertura` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            `account_type` varchar(10) NOT NULL,
                            `amount` decimal(10,2) NOT NULL,
                            `status` tinyint(1) NOT NULL DEFAULT '1',
                            PRIMARY KEY (`id`),
                            KEY clienteFK (`cliente_id`),
                            CONSTRAINT clienteFK FOREIGN KEY (`cliente_id`) REFERENCES clientes (id)
) ENGINE=InnoDB AUTO_INCREMENT=95471 DEFAULT CHARSET=latin1;
--
-- Dumping data for table `accounts`
--

LOCK TABLES contas_bancarias WRITE;
/*!40000 ALTER TABLE contas_bancarias DISABLE KEYS */;
INSERT INTO contas_bancarias VALUES
                           (95470,2000,'2020-08-22 10:20:06', 'saving', 6823.23, 1),
                           (95471,2002,'2020-08-09 10:27:22', 'checking', 3342.96, 1),
                           (95472,2001,'2020-08-09 10:35:22', 'saving', 7000, 1),
                           (95473,2001,'2020-08-09 10:38:22', 'saving', 5861.86, 1);
/*!40000 ALTER TABLE contas_bancarias ENABLE KEYS */;

UNLOCK TABLES;

DROP TABLE IF EXISTS transacoes;

CREATE TABLE transacoes (
                                id int(11) NOT NULL AUTO_INCREMENT,
                                conta_bancaria_id int(11) NOT NULL,
                                valor decimal(10,2) NOT NULL,
                                tipo_transacao varchar(10) NOT NULL,
                                data_transacao datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                PRIMARY KEY (id),
                                KEY `transacoes_FK` (conta_bancaria_id),
                                CONSTRAINT `transacoes_FK` FOREIGN KEY (conta_bancaria_id) REFERENCES contas_bancarias (id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

UNLOCK TABLES;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios (
                         login varchar(20) NOT NULL,
                         `password` varchar(20) NOT NULL,
                         `tipo` varchar(20) NOT NULL,
                         `cliente_id` int(11) DEFAULT NULL,
                         `data_criacao` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         PRIMARY KEY (login)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
--
-- Dumping data for table `users`
--

LOCK TABLES usuarios WRITE;
/*!40000 ALTER TABLE usuarios DISABLE KEYS */;
INSERT INTO usuarios VALUES
                        ('admin','pass123','admin', NULL, '2020-08-09 10:27:22'),
                        ('2001','pass123','usuario', 2001, '2020-08-09 10:27:22'),
                        ('2000','pass123','usuario', 2000, '2020-08-09 10:27:22');
/*!40000 ALTER TABLE usuarios ENABLE KEYS */;

UNLOCK TABLES;

DROP TABLE IF EXISTS token;

CREATE TABLE token (
                                       token varchar(300) NOT NULL,
                                       data_criacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                       PRIMARY KEY (token)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-08-31 10:25:19