#!/bin/bash

PGPASSWORD=password psql -U admindb -d postgres -c "
                                     -- PostgreSQL database dump
                                     --

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
                                     -- Name: tari; Type: TABLE; Schema: public; Owner: admindb
                                     --

                                    CREATE TABLE public.tari (
                                        id bigint NOT NULL,
                                        nume_tara text,
                                        latitudine double precision,
                                        longitudine double precision,
                                        CONSTRAINT id_valid CHECK (id > 0),
                                        CONSTRAINT nume_tara_valid CHECK (nume_tara <> ''::text),
                                        CONSTRAINT tari_pkey PRIMARY KEY (id),
                                        CONSTRAINT tari_nume_tara_key UNIQUE (nume_tara)
                                    );

                                     ALTER TABLE public.tari OWNER TO admindb;

                                    CREATE TABLE public.orase (
                                        id bigint NOT NULL,
                                        id_tara bigint NOT NULL,
                                        nume_oras text,
                                        latitudine double precision,
                                        longitudine double precision,
                                        CONSTRAINT id_valid CHECK (id > 0),
                                        CONSTRAINT nume_oras_valid CHECK (nume_oras <> ''::text),
                                        CONSTRAINT orase_pkey PRIMARY KEY (id),
                                        CONSTRAINT orase_id_tara_fkey FOREIGN KEY (id_tara) REFERENCES public.tari(id) ON DELETE CASCADE,
                                        CONSTRAINT orase_nume_oras_key UNIQUE (id_tara, nume_oras)
                                    );

                                     ALTER TABLE public.orase OWNER TO admindb;

                                    CREATE TABLE public.temperaturi (
                                        id bigint NOT NULL,
                                        id_oras bigint NOT NULL,
                                        valoare double precision NOT NULL,
                                        timestamp timestamp NOT NULL,
                                        CONSTRAINT id_valid CHECK (id > 0),
                                        CONSTRAINT temperaturi_pkey PRIMARY KEY (id),
                                        CONSTRAINT temperaturi_id_oras_fkey FOREIGN KEY (id_oras) REFERENCES public.orase(id) ON DELETE CASCADE,
                                        CONSTRAINT temperaturi_timestamp_key UNIQUE (id_oras, timestamp)
                                    );

                                     ALTER TABLE public.temperaturi OWNER TO admindb;

                                     --
                                     -- Name: tari_id_seq; Type: SEQUENCE; Schema: public; Owner: admindb
                                     --

                                    CREATE SEQUENCE public.tari_id_seq
                                        START WITH 1
                                        INCREMENT BY 1
                                        NO MINVALUE
                                        NO MAXVALUE
                                        CACHE 1;

                                    ALTER TABLE public.tari_id_seq OWNER TO admindb;

                                    --
                                    -- Name: tari_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admindb
                                    --

                                    ALTER SEQUENCE public.tari_id_seq OWNED BY public.tari.id;

                                    --
                                    -- Name: orase_id_seq; Type: SEQUENCE; Schema: public; Owner: admindb
                                    --

                                    CREATE SEQUENCE public.orase_id_seq
                                        START WITH 1
                                        INCREMENT BY 1
                                        NO MINVALUE
                                        NO MAXVALUE
                                        CACHE 1;

                                    ALTER TABLE public.orase_id_seq OWNER TO admindb;

                                    --
                                    -- Name: orase_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admindb
                                    --

                                    ALTER SEQUENCE public.orase_id_seq OWNED BY public.orase.id;

                                    --
                                    -- Name: temperaturi_id_seq; Type: SEQUENCE; Schema: public; Owner: admindb
                                    --

                                    CREATE SEQUENCE public.temperaturi_id_seq
                                        START WITH 1
                                        INCREMENT BY 1
                                        NO MINVALUE
                                        NO MAXVALUE
                                        CACHE 1;

                                    ALTER TABLE public.temperaturi_id_seq OWNER TO admindb;

                                    --
                                    -- Name: temperaturi_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admindb
                                    --

                                    ALTER SEQUENCE public.temperaturi_id_seq OWNED BY public.temperaturi.id;

                                    --
                                    -- Name: orase id; Type: DEFAULT; Schema: public; Owner: admindb
                                    --

                                    ALTER TABLE ONLY public.orase ALTER COLUMN id SET DEFAULT nextval('public.orase_id_seq'::regclass);

                                    --
                                    -- Name: tari id; Type: DEFAULT; Schema: public; Owner: admindb
                                    --

                                    ALTER TABLE ONLY public.tari ALTER COLUMN id SET DEFAULT nextval('public.tari_id_seq'::regclass);

                                    --
                                    -- Name: temperaturi id; Type: DEFAULT; Schema: public; Owner: admindb
                                    --

                                    ALTER TABLE ONLY public.temperaturi ALTER COLUMN id SET DEFAULT nextval('public.temperaturi_id_seq'::regclass);

                                    --
                                    -- PostgreSQL database dump complete
                                    --"
