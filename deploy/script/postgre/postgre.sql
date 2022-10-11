-- DROP SCHEMA hospital_v1n;

CREATE SCHEMA hospital_v1n AUTHORIZATION pgadmin;

-- DROP SEQUENCE hospital_v1n.pd_hospital_id_seq;

CREATE SEQUENCE hospital_v1n.pd_hospital_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START 1
	CACHE 1
	NO CYCLE;-- hospital_v1n.pd_hospital definition

-- Drop table

-- DROP TABLE hospital_v1n.pd_hospital;

CREATE TABLE hospital_v1n.pd_hospital (
	id serial4 NOT NULL,
	"name" varchar(120) NOT NULL,
	logo text NOT NULL,
	welfare bool NOT NULL,
	t_date date NOT NULL,
	t_time time NOT NULL,
	CONSTRAINT pd_hospital_pkey PRIMARY KEY (id)
);

COMMENT ON TABLE hospital_v1n.pd_hospital IS 'hospital information.';