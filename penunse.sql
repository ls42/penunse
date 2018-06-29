--
-- PostgreSQL database dump
--

-- Dumped from database version 10.3
-- Dumped by pg_dump version 10.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner:
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner:
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: auth_identities; Type: TABLE; Schema: public; Owner: penunse
--

CREATE TABLE public.auth_identities (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    provider text,
    uid text,
    encrypted_password text,
    user_id text,
    confirmed_at timestamp with time zone,
    sign_logs text
);


ALTER TABLE public.auth_identities OWNER TO penunse;

--
-- Name: auth_identities_id_seq; Type: SEQUENCE; Schema: public; Owner: penunse
--

CREATE SEQUENCE public.auth_identities_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.auth_identities_id_seq OWNER TO penunse;

--
-- Name: auth_identities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: penunse
--

ALTER SEQUENCE public.auth_identities_id_seq OWNED BY public.auth_identities.id;


--
-- Name: tags; Type: TABLE; Schema: public; Owner: penunse
--

CREATE TABLE public.tags (
    id integer NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    name text
);


ALTER TABLE public.tags OWNER TO penunse;

--
-- Name: tags_id_seq; Type: SEQUENCE; Schema: public; Owner: penunse
--

CREATE SEQUENCE public.tags_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tags_id_seq OWNER TO penunse;

--
-- Name: tags_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: penunse
--

ALTER SEQUENCE public.tags_id_seq OWNED BY public.tags.id;


--
-- Name: transactions; Type: TABLE; Schema: public; Owner: penunse
--

CREATE TABLE public.transactions (
    id integer NOT NULL,
    "user" integer NOT NULL,
    amount numeric NOT NULL,
    note text NOT NULL,
    created timestamp with time zone DEFAULT now(),
    updated timestamp with time zone DEFAULT now(),
    deleted timestamp with time zone
);


ALTER TABLE public.transactions OWNER TO penunse;

--
-- Name: transactions_id_seq; Type: SEQUENCE; Schema: public; Owner: penunse
--

CREATE SEQUENCE public.transactions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transactions_id_seq OWNER TO penunse;

--
-- Name: transactions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: penunse
--

ALTER SEQUENCE public.transactions_id_seq OWNED BY public.transactions.id;


--
-- Name: transactions_tags; Type: TABLE; Schema: public; Owner: penunse
--

CREATE TABLE public.transactions_tags (
    transaction_id integer NOT NULL,
    tag_id integer NOT NULL
);


ALTER TABLE public.transactions_tags OWNER TO penunse;

--
-- Name: users; Type: TABLE; Schema: public; Owner: penunse
--

CREATE TABLE public.users (
    id integer NOT NULL,
    login text,
    first text,
    created timestamp with time zone,
    updated timestamp with time zone,
    deleted timestamp with time zone,
    pass bytea
);


ALTER TABLE public.users OWNER TO penunse;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: penunse
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO penunse;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: penunse
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: auth_identities id; Type: DEFAULT; Schema: public; Owner: penunse
--

ALTER TABLE ONLY public.auth_identities ALTER COLUMN id SET DEFAULT nextval('public.auth_identities_id_seq'::regclass);


--
-- Name: tags id; Type: DEFAULT; Schema: public; Owner: penunse
--

ALTER TABLE ONLY public.tags ALTER COLUMN id SET DEFAULT nextval('public.tags_id_seq'::regclass);


--
-- Name: transactions id; Type: DEFAULT; Schema: public; Owner: penunse
--

ALTER TABLE ONLY public.transactions ALTER COLUMN id SET DEFAULT nextval('public.transactions_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: penunse
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: auth_identities auth_identities_pkey; Type: CONSTRAINT; Schema: public; Owner: penunse
--

ALTER TABLE ONLY public.auth_identities
    ADD CONSTRAINT auth_identities_pkey PRIMARY KEY (id);


--
-- Name: tags tags_name_key; Type: CONSTRAINT; Schema: public; Owner: penunse
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_name_key UNIQUE (name);


--
-- Name: tags tags_pkey; Type: CONSTRAINT; Schema: public; Owner: penunse
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_pkey PRIMARY KEY (id);


--
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: penunse
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- Name: transactions_tags transactions_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: penunse
--

ALTER TABLE ONLY public.transactions_tags
    ADD CONSTRAINT transactions_tags_pkey PRIMARY KEY (transaction_id, tag_id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: penunse
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_auth_identities_deleted_at; Type: INDEX; Schema: public; Owner: penunse
--

CREATE INDEX idx_auth_identities_deleted_at ON public.auth_identities USING btree (deleted_at);


--
-- Name: idx_tags_deleted_at; Type: INDEX; Schema: public; Owner: penunse
--

CREATE INDEX idx_tags_deleted_at ON public.tags USING btree (deleted_at);


--
-- Name: transactions_tags tags_id; Type: FK CONSTRAINT; Schema: public; Owner: penunse
--

ALTER TABLE ONLY public.transactions_tags
    ADD CONSTRAINT tags_id FOREIGN KEY (tag_id) REFERENCES public.tags(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: transactions_tags transactions_id; Type: FK CONSTRAINT; Schema: public; Owner: penunse
--

ALTER TABLE ONLY public.transactions_tags
    ADD CONSTRAINT transactions_id FOREIGN KEY (transaction_id) REFERENCES public.transactions(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--
