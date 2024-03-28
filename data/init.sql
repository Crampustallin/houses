--
-- PostgreSQL database dump
--

-- Dumped from database version 14.10 (Homebrew)
-- Dumped by pg_dump version 14.10 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: addresses; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.addresses (
    address_id integer NOT NULL,
    city character varying(100) NOT NULL,
    street character varying(100) NOT NULL,
    house_number character varying(10) NOT NULL,
    apartment_number character varying(10),
    postal_code character varying(20)
);


--
-- Name: addresses_address_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.addresses_address_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: addresses_address_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.addresses_address_id_seq OWNED BY public.addresses.address_id;


--
-- Name: properties; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.properties (
    property_id integer NOT NULL,
    property_type_id integer NOT NULL,
    address_id integer NOT NULL,
    price numeric(20,2) NOT NULL,
    rooms integer,
    area numeric(10,2),
    description text
);


--
-- Name: properties_property_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.properties_property_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: properties_property_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.properties_property_id_seq OWNED BY public.properties.property_id;


--
-- Name: property_types; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.property_types (
    property_type_id integer NOT NULL,
    property_type_name character varying(100) NOT NULL
);


--
-- Name: property_types_property_type_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.property_types_property_type_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: property_types_property_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.property_types_property_type_id_seq OWNED BY public.property_types.property_type_id;


--
-- Name: addresses address_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.addresses ALTER COLUMN address_id SET DEFAULT nextval('public.addresses_address_id_seq'::regclass);


--
-- Name: properties property_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.properties ALTER COLUMN property_id SET DEFAULT nextval('public.properties_property_id_seq'::regclass);


--
-- Name: property_types property_type_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.property_types ALTER COLUMN property_type_id SET DEFAULT nextval('public.property_types_property_type_id_seq'::regclass);


--
-- Data for Name: addresses; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.addresses (address_id, city, street, house_number, apartment_number, postal_code) FROM stdin;
2	Ваш_город	Ул. Ленина	5	3	\N
3	Ваш_город	Ул. Центральная	12	\N	\N
4	Ваш_город	СНТ "Рассвет"	25	\N	\N
5	Ваш_город	Проспект Победы	20	8	\N
6	Ваш_город	Ул. Садовая	7	\N	\N
7	Ваш_город	Деревня Солнечная	участок 10	\N	\N
8	Ваш_город	Проспект Мира	15	12	\N
9	Ваш_город	Ул. Новая	3	\N	\N
10	Ваш_город	Деревня Зеленая	участок 5	\N	\N
11	Ваш_город	Ул. Юбилейная	30	22	\N
12	Ваш_город	Ул. Лесная	14	\N	\N
13	Ваш_город	Проспект Свободы	7	15	\N
14	Ваш_город	Ул. Солнечная	25	\N	\N
15	Ваш_город	Дачное товарищество "Луч"	участок 18	\N	\N
16	Ваш_город	Ул. Звездная	8	42	\N
17	Ваш_город	Пер. Сиреневый	3	\N	\N
18	Ваш_город	Ул. Цветочная	14	30	\N
19	Ваш_город	Ул. Луговая	9	\N	\N
20	Ваш_город	СНТ "Солнечное"	участок 32	\N	\N
\.


--
-- Data for Name: properties; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.properties (property_id, property_type_id, address_id, price, rooms, area, description) FROM stdin;
3	1	2	1200000.00	1	45.50	Студия в новом доме с ремонтом и мебелью
7	1	3	1200000.00	1	45.50	
\.


--
-- Data for Name: property_types; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.property_types (property_type_id, property_type_name) FROM stdin;
1	Квартира
2	Дом
3	Земельный участок
\.


--
-- Name: addresses_address_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.addresses_address_id_seq', 20, true);


--
-- Name: properties_property_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.properties_property_id_seq', 8, true);


--
-- Name: property_types_property_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.property_types_property_type_id_seq', 3, true);


--
-- Name: addresses addresses_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.addresses
    ADD CONSTRAINT addresses_pkey PRIMARY KEY (address_id);


--
-- Name: properties properties_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.properties
    ADD CONSTRAINT properties_pkey PRIMARY KEY (property_id);


--
-- Name: property_types property_types_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.property_types
    ADD CONSTRAINT property_types_pkey PRIMARY KEY (property_type_id);


--
-- Name: properties properties_address_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.properties
    ADD CONSTRAINT properties_address_id_fkey FOREIGN KEY (address_id) REFERENCES public.addresses(address_id);


--
-- Name: properties properties_property_type_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.properties
    ADD CONSTRAINT properties_property_type_id_fkey FOREIGN KEY (property_type_id) REFERENCES public.property_types(property_type_id);
