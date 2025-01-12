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

--
-- Name: journal_entries_type; Type: TYPE; Schema: public; Owner: -
--

CREATE TYPE public.journal_entries_type AS ENUM (
    'DEBIT',
    'CREDIT'
);


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: accounts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.accounts (
    id character varying(100) NOT NULL,
    name character varying(50) NOT NULL,
    company_id character varying(255) NOT NULL,
    type character varying(255) NOT NULL,
    code character varying(20) NOT NULL,
    is_lock boolean NOT NULL,
    status boolean NOT NULL,
    created_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    updated_at timestamp with time zone
);


--
-- Name: categories; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.categories (
    id character varying(100) NOT NULL,
    name character varying(50) NOT NULL,
    company_id character varying(255) NOT NULL,
    type character varying(255) NOT NULL,
    code character varying(20) NOT NULL,
    status boolean NOT NULL,
    created_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    updated_at timestamp with time zone
);


--
-- Name: geographies; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.geographies (
    id character varying(100) NOT NULL,
    city_code character varying(20) NOT NULL,
    province_code character varying(20) NOT NULL,
    district_code character varying(20) NOT NULL,
    sub_district_code character varying(20) NOT NULL,
    province_name character varying(20) NOT NULL,
    district_name character varying(20) NOT NULL,
    city_name character varying(20) NOT NULL,
    sub_district_name character varying(20) NOT NULL
);


--
-- Name: journal_entries; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.journal_entries (
    id character varying(100) NOT NULL,
    amount numeric(20,2),
    account_id character varying(50) NOT NULL,
    company_id character varying(50) NOT NULL,
    type public.journal_entries_type NOT NULL,
    additional_data jsonb,
    note character varying(255),
    date timestamp with time zone NOT NULL,
    transaction_code character varying(255),
    created_at timestamp with time zone NOT NULL
);


--
-- Name: material_conversions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.material_conversions (
    id character varying(100) NOT NULL,
    material_product_id character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    quantity integer NOT NULL,
    deleted_at timestamp with time zone
);


--
-- Name: material_products; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.material_products (
    id character varying(100) NOT NULL,
    name character varying(255) NOT NULL,
    company_id character varying(255) NOT NULL,
    smallest_unit_id character varying(255) NOT NULL,
    sku character varying(255) NOT NULL,
    category_id character varying(255) NOT NULL,
    status boolean NOT NULL,
    current_quantity integer NOT NULL,
    deleted_at timestamp with time zone,
    updated_at timestamp with time zone,
    created_at timestamp with time zone
);


--
-- Name: material_stocks; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.material_stocks (
    id character varying(100) NOT NULL,
    material_product_id character varying(255) NOT NULL,
    product_type character varying(255) NOT NULL,
    quantity integer NOT NULL,
    current_quantity integer NOT NULL,
    expired_date timestamp with time zone NOT NULL,
    company_id character varying(255) NOT NULL
);


--
-- Name: promo_items; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.promo_items (
    id character varying(100) NOT NULL,
    promo_id character varying(255) NOT NULL,
    sellable_product_id character varying(255) NOT NULL
);


--
-- Name: promos; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.promos (
    id character varying(100) NOT NULL,
    name character varying(255) NOT NULL,
    quantity integer NOT NULL,
    type character varying(50) NOT NULL,
    start_date timestamp with time zone,
    end_date timestamp with time zone,
    is_all boolean DEFAULT false NOT NULL,
    deleted_at timestamp with time zone,
    company_id character varying(255) NOT NULL,
    amount integer NOT NULL
);


--
-- Name: receipt; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.receipt (
    id character varying(100) NOT NULL,
    sellable_product_id character varying(255) NOT NULL,
    material_product_id character varying(255) NOT NULL,
    quantity integer NOT NULL
);


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version character varying(128) NOT NULL
);


--
-- Name: sellable_conversions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.sellable_conversions (
    id character varying(100) NOT NULL,
    sellable_product_id character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    quantity integer NOT NULL,
    deleted_at timestamp with time zone
);


--
-- Name: sellable_products; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.sellable_products (
    id character varying(100) NOT NULL,
    name character varying(255) NOT NULL,
    company_id character varying(255) NOT NULL,
    smallest_unit_id character varying(255) NOT NULL,
    category_id character varying(255) NOT NULL,
    image character varying(255) NOT NULL,
    sku character varying(255) NOT NULL,
    description character varying(255) NOT NULL,
    status boolean NOT NULL,
    has_receipt boolean NOT NULL,
    current_quantity integer NOT NULL,
    deleted_at timestamp with time zone,
    updated_at timestamp with time zone,
    created_at timestamp with time zone,
    price numeric(15,2) NOT NULL
);


--
-- Name: sellable_stocks; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.sellable_stocks (
    id character varying(100) NOT NULL,
    sellable_product_id character varying(255) NOT NULL,
    product_type character varying(255) NOT NULL,
    quantity integer NOT NULL,
    current_quantity integer NOT NULL,
    expired_date timestamp with time zone NOT NULL,
    company_id character varying(255) NOT NULL
);


--
-- Name: tax; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.tax (
    id character varying(100) NOT NULL,
    name character varying(255) NOT NULL,
    precentage integer NOT NULL,
    deleted_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


--
-- Name: transactions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.transactions (
    id character varying(100) NOT NULL,
    title character varying(50) NOT NULL,
    name character varying(255) NOT NULL,
    additional_data jsonb,
    date date NOT NULL,
    due_date date,
    note character varying(255),
    transaction_record_code character varying(255),
    transaction_id character varying(100),
    payment_method character varying(255) NOT NULL,
    payment_type character varying(255) NOT NULL,
    amount numeric(20,2) NOT NULL,
    company_id character varying(100) NOT NULL,
    created_at timestamp with time zone NOT NULL
);


--
-- Name: units; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.units (
    id character varying(100) NOT NULL,
    name character varying(50) NOT NULL,
    company_id character varying(255) NOT NULL,
    code character varying(20) NOT NULL,
    status boolean NOT NULL,
    created_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    updated_at timestamp with time zone
);


--
-- Name: accounts accounts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (id);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: geographies geographies_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.geographies
    ADD CONSTRAINT geographies_pkey PRIMARY KEY (id);


--
-- Name: journal_entries journal_entries_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.journal_entries
    ADD CONSTRAINT journal_entries_pkey PRIMARY KEY (id);


--
-- Name: material_conversions material_conversions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.material_conversions
    ADD CONSTRAINT material_conversions_pkey PRIMARY KEY (id);


--
-- Name: material_products material_products_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.material_products
    ADD CONSTRAINT material_products_pkey PRIMARY KEY (id);


--
-- Name: material_stocks material_stocks_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.material_stocks
    ADD CONSTRAINT material_stocks_pkey PRIMARY KEY (id);


--
-- Name: promo_items promo_items_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.promo_items
    ADD CONSTRAINT promo_items_pkey PRIMARY KEY (id);


--
-- Name: promos promos_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.promos
    ADD CONSTRAINT promos_pkey PRIMARY KEY (id);


--
-- Name: receipt receipt_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.receipt
    ADD CONSTRAINT receipt_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: sellable_conversions sellable_conversions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sellable_conversions
    ADD CONSTRAINT sellable_conversions_pkey PRIMARY KEY (id);


--
-- Name: sellable_products sellable_products_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sellable_products
    ADD CONSTRAINT sellable_products_pkey PRIMARY KEY (id);


--
-- Name: sellable_stocks sellable_stocks_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.sellable_stocks
    ADD CONSTRAINT sellable_stocks_pkey PRIMARY KEY (id);


--
-- Name: tax tax_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tax
    ADD CONSTRAINT tax_pkey PRIMARY KEY (id);


--
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- Name: units units_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.units
    ADD CONSTRAINT units_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.schema_migrations (version) VALUES
    ('20241112094421'),
    ('20241112095711'),
    ('20241112101144'),
    ('20241112101751'),
    ('20241112102007'),
    ('20241112102059'),
    ('20241112102814'),
    ('20241112134927'),
    ('20241112135155'),
    ('20241112135616'),
    ('20241112135842'),
    ('20241112140014'),
    ('20241112140542'),
    ('20241112141157'),
    ('20241113074549'),
    ('20241113074956');
