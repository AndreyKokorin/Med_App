--
-- PostgreSQL database dump
--

-- Dumped from database version 17.2
-- Dumped by pg_dump version 17.2

-- Started on 2025-02-25 14:08:06

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
-- TOC entry 220 (class 1259 OID 16470)
-- Name: appointments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.appointments (
    id integer NOT NULL,
    patient_id integer NOT NULL,
    doctor_id integer NOT NULL,
    appointment_time timestamp without time zone NOT NULL,
    status character varying(50) NOT NULL,
    schedule_id integer DEFAULT 0 NOT NULL,
    CONSTRAINT appointments_status_check CHECK (((status)::text = ANY ((ARRAY['Pending'::character varying, 'Confirmed'::character varying, 'Completed'::character varying, 'Cancelled'::character varying])::text[])))
);


ALTER TABLE public.appointments OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 16469)
-- Name: appointments_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.appointments_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.appointments_id_seq OWNER TO postgres;

--
-- TOC entry 4900 (class 0 OID 0)
-- Dependencies: 219
-- Name: appointments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.appointments_id_seq OWNED BY public.appointments.id;


--
-- TOC entry 222 (class 1259 OID 16490)
-- Name: medical_records; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.medical_records (
    id integer NOT NULL,
    patient_id integer NOT NULL,
    diagnosis character varying(100) NOT NULL,
    recomendation text NOT NULL,
    doctor_id integer,
    created_time timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.medical_records OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 16489)
-- Name: medical_records_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.medical_records_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.medical_records_id_seq OWNER TO postgres;

--
-- TOC entry 4901 (class 0 OID 0)
-- Dependencies: 221
-- Name: medical_records_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.medical_records_id_seq OWNED BY public.medical_records.id;


--
-- TOC entry 224 (class 1259 OID 16517)
-- Name: schedules; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schedules (
    id integer NOT NULL,
    doctor_id integer,
    start_time timestamp without time zone,
    end_time timestamp without time zone,
    capacity integer DEFAULT 1,
    booked_slots integer DEFAULT 0,
    status character varying(10) DEFAULT 'active'::character varying NOT NULL
);


ALTER TABLE public.schedules OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 16516)
-- Name: schedules_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.schedules_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.schedules_id_seq OWNER TO postgres;

--
-- TOC entry 4902 (class 0 OID 0)
-- Dependencies: 223
-- Name: schedules_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.schedules_id_seq OWNED BY public.schedules.id;


--
-- TOC entry 218 (class 1259 OID 16457)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    age integer NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    roles character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 16456)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 4903 (class 0 OID 0)
-- Dependencies: 217
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 4713 (class 2604 OID 16473)
-- Name: appointments id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.appointments ALTER COLUMN id SET DEFAULT nextval('public.appointments_id_seq'::regclass);


--
-- TOC entry 4715 (class 2604 OID 16493)
-- Name: medical_records id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.medical_records ALTER COLUMN id SET DEFAULT nextval('public.medical_records_id_seq'::regclass);


--
-- TOC entry 4717 (class 2604 OID 16520)
-- Name: schedules id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schedules ALTER COLUMN id SET DEFAULT nextval('public.schedules_id_seq'::regclass);


--
-- TOC entry 4710 (class 2604 OID 16460)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 4890 (class 0 OID 16470)
-- Dependencies: 220
-- Data for Name: appointments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.appointments (id, patient_id, doctor_id, appointment_time, status, schedule_id) FROM stdin;
13	6	7	2025-02-20 14:00:00	Pending	5
\.


--
-- TOC entry 4892 (class 0 OID 16490)
-- Dependencies: 222
-- Data for Name: medical_records; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.medical_records (id, patient_id, diagnosis, recomendation, doctor_id, created_time) FROM stdin;
3	6	Ангина	Не сосать хуй	3	2025-02-17 11:32:19.445564
4	6	1	Не сосать	3	2025-02-17 11:32:19.445564
5	6	Артроз	нехуй ходить так много	3	2025-02-17 16:37:34.69996
\.


--
-- TOC entry 4894 (class 0 OID 16517)
-- Dependencies: 224
-- Data for Name: schedules; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.schedules (id, doctor_id, start_time, end_time, capacity, booked_slots, status) FROM stdin;
0	3	2025-02-18 09:00:00	2025-02-18 12:00:00	10	0	archived
1	3	2025-02-18 11:23:09.575256	2025-02-18 11:23:09.575256	1	0	archived
3	3	2025-02-18 09:00:00	2025-02-18 12:00:00	10	0	archived
4	3	2025-02-20 09:00:00	2025-02-20 12:00:00	5	0	archived
5	7	2025-02-20 12:00:00	2025-02-20 15:00:00	1	0	archived
\.


--
-- TOC entry 4888 (class 0 OID 16457)
-- Dependencies: 218
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, age, email, password, roles, created_at, updated_at) FROM stdin;
3	John	3	doctor1@example.com	$2a$10$6JN/kHOAnjsyk.lzi6syyOi7NzYkAzxHLJJgyZ2IRVYsq.RKCbgM6	doctor	2025-02-13 08:51:17.798531	2025-02-13 08:51:17.798531
2	Andrey	28	aaaa@gmail.com	$2a$10$nzv.Sn5g4xJc07d8skDCqe5hLee57TFusoN5g9FoA/xFtOIRjlz0G	admin	2025-02-12 15:48:57.399223	2025-02-12 15:48:57.399223
6	Andrey	21	andtey2003$@gmail.com	$2a$10$jKcr6RqYcH9oBXwlawveX.GbsteWY52OLXY3SFcuATDSrJC7h1sMK	user	2025-02-17 09:41:47.731169	2025-02-17 09:41:47.731169
7	doctor	25	doctor@gmail.com	$2a$10$SYSJ6FT.mHmaB5NrAkeht.Ty9wVaWvYVg23vHRElSPRkoQqrwM7B6	doctor	2025-02-20 11:37:06.224944	2025-02-20 11:37:06.224944
8	Ivan	21	andrey.kokorin.007@gmail.com	$2a$10$cJiUdwgba7DLANpoxo/g1eEaUfG7LHGtpRNiI1C9IiNt0/oTlbZOO	user	2025-02-20 15:23:44.09919	2025-02-20 15:23:44.09919
9	annananannana	21	abcdefg@gmail.com	$2a$10$Gp6wJJWDSHpMwfPi9np.MOZrY4v040ueMpu6D7jKK3SnAmH8Ya.OO	user	2025-02-20 15:28:34.461091	2025-02-20 15:28:34.461091
10	jdkfjsdjf	21	anton@gmail.com	$2a$10$jiL90sBCYIuLsmmWinCKheDUd/rrDdfegvUDHYQffqJD4YXvqNtW2	user	2025-02-20 15:31:11.771823	2025-02-20 15:31:11.771823
\.


--
-- TOC entry 4904 (class 0 OID 0)
-- Dependencies: 219
-- Name: appointments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.appointments_id_seq', 13, true);


--
-- TOC entry 4905 (class 0 OID 0)
-- Dependencies: 221
-- Name: medical_records_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.medical_records_id_seq', 5, true);


--
-- TOC entry 4906 (class 0 OID 0)
-- Dependencies: 223
-- Name: schedules_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.schedules_id_seq', 5, true);


--
-- TOC entry 4907 (class 0 OID 0)
-- Dependencies: 217
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 10, true);


--
-- TOC entry 4727 (class 2606 OID 16488)
-- Name: appointments appointment_time; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.appointments
    ADD CONSTRAINT appointment_time UNIQUE (doctor_id, appointment_time);


--
-- TOC entry 4729 (class 2606 OID 16476)
-- Name: appointments appointments_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.appointments
    ADD CONSTRAINT appointments_pkey PRIMARY KEY (id);


--
-- TOC entry 4733 (class 2606 OID 16497)
-- Name: medical_records medical_records_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.medical_records
    ADD CONSTRAINT medical_records_pkey PRIMARY KEY (id);


--
-- TOC entry 4735 (class 2606 OID 16524)
-- Name: schedules schedules_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schedules
    ADD CONSTRAINT schedules_pkey PRIMARY KEY (id);


--
-- TOC entry 4731 (class 2606 OID 16540)
-- Name: appointments unique_schedule; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.appointments
    ADD CONSTRAINT unique_schedule UNIQUE (schedule_id, appointment_time);


--
-- TOC entry 4723 (class 2606 OID 16468)
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- TOC entry 4725 (class 2606 OID 16466)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 4736 (class 2606 OID 16482)
-- Name: appointments appointments_doctor_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.appointments
    ADD CONSTRAINT appointments_doctor_id_fkey FOREIGN KEY (doctor_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- TOC entry 4737 (class 2606 OID 16477)
-- Name: appointments appointments_patient_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.appointments
    ADD CONSTRAINT appointments_patient_id_fkey FOREIGN KEY (patient_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- TOC entry 4739 (class 2606 OID 16504)
-- Name: medical_records fk_doctor_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.medical_records
    ADD CONSTRAINT fk_doctor_id FOREIGN KEY (doctor_id) REFERENCES public.users(id);


--
-- TOC entry 4738 (class 2606 OID 16534)
-- Name: appointments fk_schedule; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.appointments
    ADD CONSTRAINT fk_schedule FOREIGN KEY (schedule_id) REFERENCES public.schedules(id) ON DELETE CASCADE;


--
-- TOC entry 4740 (class 2606 OID 16498)
-- Name: medical_records medical_records_patient_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.medical_records
    ADD CONSTRAINT medical_records_patient_id_fkey FOREIGN KEY (patient_id) REFERENCES public.users(id);


--
-- TOC entry 4741 (class 2606 OID 16525)
-- Name: schedules schedules_doctor_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schedules
    ADD CONSTRAINT schedules_doctor_id_fkey FOREIGN KEY (doctor_id) REFERENCES public.users(id);


-- Completed on 2025-02-25 14:08:06

--
-- PostgreSQL database dump complete
--

