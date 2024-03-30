

--
-- Name: item_id_seq; Type: SEQUENCE; Schema: bestprice; Owner: postgres
--

CREATE SEQUENCE gmarket.item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE gmarket.item_id_seq OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: items; Type: TABLE; Schema: bestprice; Owner: postgres
--

CREATE TABLE gmarket.items (
    id                   bigint DEFAULT nextval('gmarket.item_id_seq'::regclass) NOT NULL,
	batchId              bigint NOT NULL,
    rank                 integer NOT NULL,
	goodsCode            character varying(255),
	goodsName            character varying(255),
	linkUrl              character varying(255),
	discountPrice        character varying(255),
	price                character varying(255),
	deliveryInfo         character varying(255),
	itemPriceType        character varying(255),
	imageUrl             character varying(255),
	discountRate         integer,
	expressShippingText  character varying(255),
	deliveryText         character varying(255),
	consultingPeriod     character varying(255),
	isPriceExpose        boolean,
	isStartPrice         boolean,
	isFreeShipping       boolean,
	isSmileShipping      boolean,
	isExpressShipping    boolean,
	isCartVisible        boolean,
	isBigSmileItem       boolean,
	ImageChgDt           bigint
);

CREATE TABLE gmarket.batch (
    id                   bigInt NOT NULL,
    createdat            timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE gmarket.latest (
    id                   integer NOT NULL,
    batchid                bigInt NOT NULL
);

ALTER TABLE gmarket.items OWNER TO postgres;
ALTER TABLE gmarket.batch OWNER TO postgres;
ALTER TABLE gmarket.latest OWNER TO postgres;

--
-- Name: item_id_seq; Type: SEQUENCE SET; Schema: bestprice; Owner: postgres
--

SELECT pg_catalog.setval('gmarket.item_id_seq', 1, true);


--
-- Name: items items_pkey; Type: CONSTRAINT; Schema: gmarket; Owner: postgres
--

ALTER TABLE ONLY gmarket.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);


ALTER TABLE ONLY gmarket.batch
    ADD CONSTRAINT batch_pkey PRIMARY KEY (id);

ALTER TABLE gmarket.items ADD FOREIGN KEY ("batchid") REFERENCES gmarket.batch ("id");    


-- 기본값 저장
INSERT INTO gmarket.latest VALUES (1, 1234)
