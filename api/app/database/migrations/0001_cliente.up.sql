CREATE TABLE "cliente" (
    "uuid" varchar(100) PRIMARY KEY NOT NULL,
    "nome" varchar(100) DEFAULT NULL,
    "endereco" varchar(200) DEFAULT NULL,
    "cadastrado_em" timestamp without time zone DEFAULT NULL,
    "atualizado_em" timestamp without time zone DEFAULT NULL);

