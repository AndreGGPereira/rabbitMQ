CREATE TABLE "client" (
    "uuid" varchar(100) PRIMARY KEY NOT NULL,
    "name" varchar(100) DEFAULT NULL,
    "address" varchar(200) DEFAULT NULL,
    "created_at" timestamp without time zone DEFAULT NULL,
    "updated_at" timestamp without time zone DEFAULT NULL);

