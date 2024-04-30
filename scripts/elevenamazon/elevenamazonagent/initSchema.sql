

--
-- Name: item_id_seq; Type: SEQUENCE; Schema: bestprice; Owner: postgres
--

CREATE SEQUENCE elevenamazon.item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE elevenamazon.item_id_seq OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: items; Type: TABLE; Schema: bestprice; Owner: postgres
--

CREATE TABLE elevenamazon.items (
  id                   bigint DEFAULT nextval('elevenamazon.item_id_seq'::regclass) NOT NULL,
  batchId              bigint NOT NULL,
  rank                 integer NOT NULL,
  prdNo                character varying(255),
  prdNm                character varying(255),
  content_type         character varying(255),
  imageUrl             character varying(255),
  linkUrl              character varying(255),
  last_discount_price  integer,
  delivery_info        character varying(255)
);

CREATE TABLE elevenamazon.batch (
    id                   bigInt NOT NULL,
    createdat            timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE elevenamazon.latest (
    id                   integer NOT NULL,
    batchid               bigInt NOT NULL
);

ALTER TABLE elevenamazon.items OWNER TO postgres;
ALTER TABLE elevenamazon.batch OWNER TO postgres;
ALTER TABLE elevenamazon.latest OWNER TO postgres;

--
-- Name: item_id_seq; Type: SEQUENCE SET; Schema: bestprice; Owner: postgres
--

SELECT pg_catalog.setval('elevenamazon.item_id_seq', 1, true);


--
-- Name: items items_pkey; Type: CONSTRAINT; Schema: elevenamazon; Owner: postgres
--

ALTER TABLE ONLY elevenamazon.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);


ALTER TABLE ONLY elevenamazon.batch
    ADD CONSTRAINT batch_pkey PRIMARY KEY (id);

ALTER TABLE elevenamazon.items ADD FOREIGN KEY ("batchid") REFERENCES elevenamazon.batch ("id");    


-- 기본값 저장
INSERT INTO elevenamazon.latest VALUES (1, 1234)
