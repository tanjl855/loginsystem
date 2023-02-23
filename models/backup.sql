--
-- PostgreSQL database dump
--

-- Dumped from database version 12.13 (Debian 12.13-1.pgdg110+1)
-- Dumped by pg_dump version 12.13 (Debian 12.13-1.pgdg110+1)

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
-- Name: userinfo_uid_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.userinfo_uid_seq
    START WITH 6000001
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.userinfo_uid_seq OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: userinfo; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.userinfo (
                                 uid character varying(32) DEFAULT nextval('public.userinfo_uid_seq'::regclass) NOT NULL,
                                 name character varying(50),
                                 password character varying(80),
                                 email text,
                                 phone character varying(20),
                                 created_at timestamp(6) without time zone NOT NULL,
                                 deleted_at character varying(6) DEFAULT '0'::character varying NOT NULL,
                                 updated_at timestamp(6) without time zone,
                                 status character varying DEFAULT '0'::character varying NOT NULL
);


ALTER TABLE public.userinfo OWNER TO postgres;

--
-- Data for Name: userinfo; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.userinfo (uid, name, password, email, phone, created_at, deleted_at, updated_at, status) FROM stdin;
6000032 root    63a9f0ea7bb98050796b649e85481845        root@root.com   123123  2022-11-17 17:21:06.440635   0       \N      0
\.


--
-- Name: userinfo_uid_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.userinfo_uid_seq', 6000048, true);


--
-- Name: userinfo uid_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.userinfo
    ADD CONSTRAINT uid_name UNIQUE (name, uid);


--
-- Name: userinfo userinfo_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.userinfo
    ADD CONSTRAINT userinfo_pkey PRIMARY KEY (uid);


--
-- Name: name_index; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX name_index ON public.userinfo USING btree (name);


--
-- Name: uid_index; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX uid_index ON public.userinfo USING btree (uid);


--
-- PostgreSQL database dump complete
--
