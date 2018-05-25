/*
 Navicat Premium Data Transfer

 Source Server         : Docker_postgres
 Source Server Type    : PostgreSQL
 Source Server Version : 90603
 Source Host           : local.rit.aws.regeneron.com
 Source Database       : coke_development
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 90603
 File Encoding         : utf-8

 Date: 06/15/2017 15:50:31 PM
*/

-- ----------------------------
--  Table structure for schema_migration
-- ----------------------------
DROP TABLE IF EXISTS "schema_migration";
CREATE TABLE "schema_migration" (
	"version" varchar(255) NOT NULL COLLATE "default"
)
WITH (OIDS=FALSE);

-- ----------------------------
--  Records of schema_migration
-- ----------------------------
BEGIN;
INSERT INTO "schema_migration" VALUES ('20170609173045');
COMMIT;

-- ----------------------------
--  Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "users";
CREATE TABLE "users" (
	"created_at" timestamp(6) NOT NULL,
	"updated_at" timestamp(6) NOT NULL,
	"id" uuid NOT NULL,
	"name" varchar(255) NOT NULL COLLATE "default",
	"email" varchar(255) NOT NULL COLLATE "default",
	"bio" text COLLATE "default"
)
WITH (OIDS=FALSE);

-- ----------------------------
--  Records of users
-- ----------------------------
BEGIN;
INSERT INTO "users" VALUES ('2017-06-15 11:49:33.16436', '2017-06-15 11:54:45.615226', '65c6d6ef-3a74-4aa0-9524-cb6832ca25d9', 'Romel Campbell', 'romelcampbell@gmail.com', 'Cyclistngdlinf''[vm
]vk

');
COMMIT;

-- ----------------------------
--  Indexes structure for table schema_migration
-- ----------------------------
CREATE UNIQUE INDEX  "version_idx" ON "schema_migration" USING btree("version" COLLATE "default" "pg_catalog"."text_ops" ASC NULLS LAST);

-- ----------------------------
--  Primary key structure for table users
-- ----------------------------
ALTER TABLE "users" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;

