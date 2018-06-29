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
-- Data for Name: auth_identities; Type: TABLE DATA; Schema: public; Owner: penunse
--

COPY public.auth_identities (id, created_at, updated_at, deleted_at, provider, uid, encrypted_password, user_id, confirmed_at, sign_logs) FROM stdin;
\.


--
-- Data for Name: tags; Type: TABLE DATA; Schema: public; Owner: penunse
--

COPY public.tags (id, created_at, updated_at, deleted_at, name) FROM stdin;
1	2018-06-29 20:09:26.228177+00	2018-06-29 20:09:26.228177+00	\N	ikea
\.


--
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: penunse
--

COPY public.transactions (id, "user", amount, note, created, updated, deleted) FROM stdin;
1	1	4.2	Fischburger	2018-05-22 00:00:00+00	2018-05-23 00:00:00+00	\N
2	0	8.9	Dinner at Ikea	2018-06-29 00:00:00+00	2018-06-29 00:00:00+00	\N
3	0	36.96	Electronics from Ikea	2018-06-29 19:51:35.651877+00	2018-06-29 19:51:35.651877+00	\N
\.


--
-- Data for Name: transactions_tags; Type: TABLE DATA; Schema: public; Owner: penunse
--

COPY public.transactions_tags (transaction_id, tag_id) FROM stdin;
2	1
3	1
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: penunse
--

COPY public.users (id, login, first, created, updated, deleted, pass) FROM stdin;
\.


--
-- Name: auth_identities_id_seq; Type: SEQUENCE SET; Schema: public; Owner: penunse
--

SELECT pg_catalog.setval('public.auth_identities_id_seq', 1, false);


--
-- Name: tags_id_seq; Type: SEQUENCE SET; Schema: public; Owner: penunse
--

SELECT pg_catalog.setval('public.tags_id_seq', 1, true);


--
-- Name: transactions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: penunse
--

SELECT pg_catalog.setval('public.transactions_id_seq', 3, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: penunse
--

SELECT pg_catalog.setval('public.users_id_seq', 1, false);


--
-- PostgreSQL database dump complete
--
