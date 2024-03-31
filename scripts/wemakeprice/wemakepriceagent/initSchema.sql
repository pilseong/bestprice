

--
-- Name: item_id_seq; Type: SEQUENCE; Schema: bestprice; Owner: postgres
--

CREATE SEQUENCE wemakeprice.item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE wemakeprice.item_id_seq OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: items; Type: TABLE; Schema: bestprice; Owner: postgres
--

CREATE TABLE wemakeprice.items (
  id                   bigint DEFAULT nextval('wemakeprice.item_id_seq'::regclass) NOT NULL,
  batchId              bigint NOT NULL,
  rank                 integer NOT NULL,
  linkInfo             bigint NOT NULL,
  linkType             character varying(255),
  dispNm               character varying(255),
  mediumImgUrl         character varying(255),
  originPrice          integer,
  salePrice            integer,
  discountPrice        integer,
  discountType         character varying(255),
  discountText         character varying(255),
  discountRate         integer,
  salesCount           integer,
  shipText             character varying(255),
  autoLabel1           character varying(255),
  reviewCount          integer,
  reviewAvgPoint       float
);

CREATE TABLE wemakeprice.batch (
    id                   bigInt NOT NULL,
    createdat            timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE wemakeprice.latest (
    id                   integer NOT NULL,
    batchid               bigInt NOT NULL
);

ALTER TABLE wemakeprice.items OWNER TO postgres;
ALTER TABLE wemakeprice.batch OWNER TO postgres;
ALTER TABLE wemakeprice.latest OWNER TO postgres;

--
-- Name: item_id_seq; Type: SEQUENCE SET; Schema: bestprice; Owner: postgres
--

SELECT pg_catalog.setval('wemakeprice.item_id_seq', 1, true);


--
-- Name: items items_pkey; Type: CONSTRAINT; Schema: wemakeprice; Owner: postgres
--

ALTER TABLE ONLY wemakeprice.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);


ALTER TABLE ONLY wemakeprice.batch
    ADD CONSTRAINT batch_pkey PRIMARY KEY (id);

ALTER TABLE wemakeprice.items ADD FOREIGN KEY ("batchid") REFERENCES wemakeprice.batch ("id");    


-- 기본값 저장
INSERT INTO wemakeprice.latest VALUES (1, 1234)
