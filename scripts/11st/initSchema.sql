

--
-- Name: item_id_seq; Type: SEQUENCE; Schema: bestprice; Owner: postgres
--

CREATE SEQUENCE elevenst.item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE elevenst.item_id_seq OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: items; Type: TABLE; Schema: bestprice; Owner: postgres
--

CREATE TABLE elevenst.items (
  id                   bigint DEFAULT nextval('elevenst.item_id_seq'::regclass) NOT NULL,
  batchId              bigint NOT NULL,
  productId            character varying(255),
  rank                 integer,
  title                character varying(255),
  linkInfo             character varying(255),
  imageLink            character varying(255),
  originalPrice        character varying(255),
  discountPrice        character varying(255),
  discountRate         character varying(255),
  shippingText         character varying(255)
);

CREATE TABLE elevenst.batch (
    id                   bigInt NOT NULL,
    createdat            timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE elevenst.latest (
    id                   integer NOT NULL,
    batchid                bigInt NOT NULL
);

ALTER TABLE elevenst.items OWNER TO postgres;
ALTER TABLE elevenst.batch OWNER TO postgres;
ALTER TABLE elevenst.latest OWNER TO postgres;

--
-- Name: item_id_seq; Type: SEQUENCE SET; Schema: bestprice; Owner: postgres
--

SELECT pg_catalog.setval('elevenst.item_id_seq', 1, true);


--
-- Name: items items_pkey; Type: CONSTRAINT; Schema: elevenst; Owner: postgres
--

ALTER TABLE ONLY elevenst.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);


ALTER TABLE ONLY elevenst.batch
    ADD CONSTRAINT batch_pkey PRIMARY KEY (id);

ALTER TABLE elevenst.items ADD FOREIGN KEY ("batchid") REFERENCES elevenst.batch ("id");    


-- 기본값 저장
INSERT INTO elevenst.latest VALUES (1, 1234)
