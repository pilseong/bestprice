

--
-- Name: item_id_seq; Type: SEQUENCE; Schema: bestprice; Owner: postgres
--

CREATE SEQUENCE tmon.item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE tmon.item_id_seq OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: items; Type: TABLE; Schema: bestprice; Owner: postgres
--

CREATE TABLE tmon.items (
  id                   bigint DEFAULT nextval('tmon.item_id_seq'::regclass) NOT NULL,
	batchId              bigint NOT NULL,
  rank                 integer NOT NULL,
	startDate	           character varying(255),
	endDate	             character varying(255),
	title	               character varying(255),
  dealNo               bigint NOT NULL,
	deliveryFeePolicy	   character varying(255),
	deliveryFee          integer,
	categoryName         character varying(255),
	pc3ColImageUrl       character varying(255),
	originalPrice        integer,
	price                integer,
	discountRate         integer
);

CREATE TABLE tmon.batch (
    id                   bigInt NOT NULL,
    createdat            timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE tmon.latest (
    id                   integer NOT NULL,
    batchid                bigInt NOT NULL
);

ALTER TABLE tmon.items OWNER TO postgres;
ALTER TABLE tmon.batch OWNER TO postgres;
ALTER TABLE tmon.latest OWNER TO postgres;

--
-- Name: item_id_seq; Type: SEQUENCE SET; Schema: bestprice; Owner: postgres
--

SELECT pg_catalog.setval('tmon.item_id_seq', 1, true);


--
-- Name: items items_pkey; Type: CONSTRAINT; Schema: tmon; Owner: postgres
--

ALTER TABLE ONLY tmon.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);


ALTER TABLE ONLY tmon.batch
    ADD CONSTRAINT batch_pkey PRIMARY KEY (id);

ALTER TABLE tmon.items ADD FOREIGN KEY ("batchid") REFERENCES tmon.batch ("id");    


-- 기본값 저장
INSERT INTO tmon.latest VALUES (1, 1234)
