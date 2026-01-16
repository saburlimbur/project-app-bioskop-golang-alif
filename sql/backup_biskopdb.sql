--
-- PostgreSQL database dump
--

-- Dumped from database version 17.4
-- Dumped by pg_dump version 17.4

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
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
-- Name: bookings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bookings (
    id integer NOT NULL,
    user_id integer NOT NULL,
    showtime_id integer NOT NULL,
    seat_id integer NOT NULL,
    booking_code character varying(20) NOT NULL,
    booking_status character varying(20) DEFAULT 'pending'::character varying,
    total_price numeric(10,2) NOT NULL,
    expired_at timestamp without time zone NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.bookings OWNER TO postgres;

--
-- Name: bookings_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bookings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.bookings_id_seq OWNER TO postgres;

--
-- Name: bookings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bookings_id_seq OWNED BY public.bookings.id;


--
-- Name: cinemas; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cinemas (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    address text NOT NULL,
    city character varying(50) NOT NULL,
    phone_number character varying(20),
    total_seats integer DEFAULT 50 NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.cinemas OWNER TO postgres;

--
-- Name: cinemas_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.cinemas_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.cinemas_id_seq OWNER TO postgres;

--
-- Name: cinemas_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.cinemas_id_seq OWNED BY public.cinemas.id;


--
-- Name: movies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.movies (
    id integer NOT NULL,
    title character varying(200) NOT NULL,
    genre character varying(100),
    duration integer NOT NULL,
    rating character varying(10),
    description text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    thumbnail text,
    poster text
);


ALTER TABLE public.movies OWNER TO postgres;

--
-- Name: movies_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.movies_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.movies_id_seq OWNER TO postgres;

--
-- Name: movies_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.movies_id_seq OWNED BY public.movies.id;


--
-- Name: otp_verifications; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.otp_verifications (
    id integer NOT NULL,
    user_id integer NOT NULL,
    otp_code character varying(6) NOT NULL,
    expires_at timestamp without time zone NOT NULL,
    is_used boolean DEFAULT false,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.otp_verifications OWNER TO postgres;

--
-- Name: otp_verifications_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.otp_verifications_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.otp_verifications_id_seq OWNER TO postgres;

--
-- Name: otp_verifications_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.otp_verifications_id_seq OWNED BY public.otp_verifications.id;


--
-- Name: payment_methods; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.payment_methods (
    id integer NOT NULL,
    name character varying(50) NOT NULL,
    code character varying(20) NOT NULL,
    description text,
    is_active boolean DEFAULT true,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.payment_methods OWNER TO postgres;

--
-- Name: payment_methods_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.payment_methods_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.payment_methods_id_seq OWNER TO postgres;

--
-- Name: payment_methods_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.payment_methods_id_seq OWNED BY public.payment_methods.id;


--
-- Name: payments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.payments (
    id integer NOT NULL,
    booking_id integer NOT NULL,
    payment_method_id integer NOT NULL,
    amount numeric(10,2) NOT NULL,
    payment_status character varying(20) DEFAULT 'pending'::character varying,
    payment_date timestamp without time zone,
    transaction_id character varying(100),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.payments OWNER TO postgres;

--
-- Name: payments_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.payments_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.payments_id_seq OWNER TO postgres;

--
-- Name: payments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.payments_id_seq OWNED BY public.payments.id;


--
-- Name: seats; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.seats (
    id integer NOT NULL,
    cinema_id integer NOT NULL,
    seat_number character varying(10) NOT NULL,
    row_number character varying(5) NOT NULL,
    seat_type character varying(20) DEFAULT 'regular'::character varying,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.seats OWNER TO postgres;

--
-- Name: seats_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.seats_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.seats_id_seq OWNER TO postgres;

--
-- Name: seats_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.seats_id_seq OWNED BY public.seats.id;


--
-- Name: sessions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sessions (
    id integer NOT NULL,
    user_id integer NOT NULL,
    token character varying(255) NOT NULL,
    expires_at timestamp without time zone NOT NULL,
    ip_address character varying(45),
    device_info text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    revoked_at timestamp without time zone
);


ALTER TABLE public.sessions OWNER TO postgres;

--
-- Name: sessions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sessions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sessions_id_seq OWNER TO postgres;

--
-- Name: sessions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sessions_id_seq OWNED BY public.sessions.id;


--
-- Name: showtimes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.showtimes (
    id integer NOT NULL,
    cinema_id integer NOT NULL,
    movie_id integer NOT NULL,
    show_date date NOT NULL,
    show_time time without time zone NOT NULL,
    price numeric(10,2) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.showtimes OWNER TO postgres;

--
-- Name: showtimes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.showtimes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.showtimes_id_seq OWNER TO postgres;

--
-- Name: showtimes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.showtimes_id_seq OWNED BY public.showtimes.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(50) NOT NULL,
    email character varying(100) NOT NULL,
    password character varying(255) NOT NULL,
    full_name character varying(100) NOT NULL,
    phone_number character varying(20),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    is_verified boolean DEFAULT false,
    email_verified_at timestamp without time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
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
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: bookings id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings ALTER COLUMN id SET DEFAULT nextval('public.bookings_id_seq'::regclass);


--
-- Name: cinemas id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cinemas ALTER COLUMN id SET DEFAULT nextval('public.cinemas_id_seq'::regclass);


--
-- Name: movies id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.movies ALTER COLUMN id SET DEFAULT nextval('public.movies_id_seq'::regclass);


--
-- Name: otp_verifications id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.otp_verifications ALTER COLUMN id SET DEFAULT nextval('public.otp_verifications_id_seq'::regclass);


--
-- Name: payment_methods id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payment_methods ALTER COLUMN id SET DEFAULT nextval('public.payment_methods_id_seq'::regclass);


--
-- Name: payments id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payments ALTER COLUMN id SET DEFAULT nextval('public.payments_id_seq'::regclass);


--
-- Name: seats id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.seats ALTER COLUMN id SET DEFAULT nextval('public.seats_id_seq'::regclass);


--
-- Name: sessions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions ALTER COLUMN id SET DEFAULT nextval('public.sessions_id_seq'::regclass);


--
-- Name: showtimes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.showtimes ALTER COLUMN id SET DEFAULT nextval('public.showtimes_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: bookings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bookings (id, user_id, showtime_id, seat_id, booking_code, booking_status, total_price, expired_at, created_at, updated_at) FROM stdin;
2	5	1	1	BK-859738	confirmed	50000.00	2026-01-13 17:26:26.004018	2026-01-13 17:11:25.999761	2026-01-15 13:11:33.32591
3	5	1	2	BK-304134	confirmed	50000.00	2026-01-13 17:28:20.958997	2026-01-13 17:13:20.957059	2026-01-15 13:32:53.959207
4	8	1	3	BK-363330	confirmed	50000.00	2026-01-16 16:30:22.024545	2026-01-16 16:15:22.023738	2026-01-16 16:22:49.680189
\.


--
-- Data for Name: cinemas; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.cinemas (id, name, address, city, phone_number, total_seats, created_at, updated_at) FROM stdin;
1	Cineplex Galaxy	Jl. Sudirman No. 10	Jakarta	021-5551234	150	2026-01-12 16:17:29.931439	2026-01-12 16:17:29.931439
2	Transmart Cinema	Jl. Raya Bogor No. 45	Depok	021-5555678	120	2026-01-12 16:17:29.931439	2026-01-12 16:17:29.931439
3	Mall Taman Anggrek XXI	Jl. Letjen S. Parman No. 21	Jakarta	021-5559012	200	2026-01-12 16:17:29.931439	2026-01-12 16:17:29.931439
4	CineOne BSD	Jl. BSD Raya Utama No. 88	Tangerang	021-5553456	180	2026-01-12 16:17:29.931439	2026-01-12 16:17:29.931439
5	Plaza Semanggi XXI	Jl. Gatot Subroto No. 50	Jakarta	021-5557890	160	2026-01-12 16:17:29.931439	2026-01-12 16:17:29.931439
6	Cinepolis Bekasi	Jl. Ahmad Yani No. 99	Bekasi	021-5552345	140	2026-01-12 16:17:29.931439	2026-01-12 16:17:29.931439
\.


--
-- Data for Name: movies; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.movies (id, title, genre, duration, rating, description, created_at, updated_at, thumbnail, poster) FROM stdin;
1	Avengers: Endgame	Action	180	PG-13	Superhero movie	2026-01-12 23:24:46.104073	2026-01-12 23:24:46.104073	https://images2.alphacoders.com/131/thumb-1920-1315111.jpg	https://id-live-01.slatic.net/p/7b170ba1d8970b4082afa04fb26cb666.jpg
\.


--
-- Data for Name: otp_verifications; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.otp_verifications (id, user_id, otp_code, expires_at, is_used, created_at) FROM stdin;
1	6	253422	2026-01-16 14:00:41.655997	f	2026-01-16 13:50:41.64513
2	7	397831	2026-01-16 14:11:43.985911	t	2026-01-16 14:01:43.982927
3	8	060582	2026-01-16 15:55:26.441057	t	2026-01-16 15:45:26.428015
\.


--
-- Data for Name: payment_methods; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.payment_methods (id, name, code, description, is_active, created_at) FROM stdin;
1	Transfer Bank	BANK_TRANSFER	Transfer via bank	t	2026-01-13 20:50:47.925487
2	OVO	OVO	Dompet digital OVO	t	2026-01-13 20:50:47.925487
3	GoPay	GOPAY	Dompet digital GoPay	t	2026-01-13 20:50:47.925487
\.


--
-- Data for Name: payments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.payments (id, booking_id, payment_method_id, amount, payment_status, payment_date, transaction_id, created_at, updated_at) FROM stdin;
6	2	1	50000.00	success	2026-01-15 13:11:33.335509	TRX-20260115131133-989341	2026-01-15 13:11:33.32591	2026-01-15 13:11:33.32591
7	3	1	50000.00	success	2026-01-15 13:32:53.961777	TRX-20260115133253-478425	2026-01-15 13:32:53.959207	2026-01-15 13:32:53.959207
8	4	1	50000.00	success	2026-01-16 16:22:49.686959	TRX-20260116162249-735135	2026-01-16 16:22:49.680189	2026-01-16 16:22:49.680189
\.


--
-- Data for Name: seats; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.seats (id, cinema_id, seat_number, row_number, seat_type, created_at) FROM stdin;
1	1	A1	A	regular	2026-01-12 23:26:28.882653
2	1	A2	A	regular	2026-01-12 23:26:28.882653
3	1	A3	A	regular	2026-01-12 23:26:28.882653
4	1	B1	B	vip	2026-01-12 23:26:28.882653
5	1	B2	B	vip	2026-01-12 23:26:28.882653
\.


--
-- Data for Name: sessions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sessions (id, user_id, token, expires_at, ip_address, device_info, created_at, revoked_at) FROM stdin;
1	5	598fef96-0b5a-4024-b552-aa79315f24c4	2026-01-13 13:44:56.322548	[::1]:56259	PostmanRuntime/7.51.0	2026-01-12 13:44:56.329797	\N
2	5	3729e2df-a02e-4509-a9d2-f324405877a6	2026-01-13 15:15:01.932079	[::1]:57459	PostmanRuntime/7.51.0	2026-01-12 15:15:01.938606	2026-01-12 15:15:21.05886
3	5	7d4c79fe-3f54-4355-8fb5-70c0833f18e0	2026-01-13 16:09:26.274613	[::1]:58005	PostmanRuntime/7.51.0	2026-01-12 16:09:26.281619	\N
4	5	61ae3c66-71ce-40c4-aeac-55f71c208ffc	2026-01-13 16:11:13.669187	[::1]:58030	PostmanRuntime/7.51.0	2026-01-12 16:11:13.669469	\N
5	5	d31c5bc7-fd3f-4746-a3a3-57d51f91489b	2026-01-13 22:38:40.75686	[::1]:56866	PostmanRuntime/7.51.0	2026-01-12 22:38:40.765737	2026-01-12 22:39:15.328794
6	5	7c65cea2-3bc2-4c9b-96ed-a4319cea0e0b	2026-01-14 17:05:35.420796	[::1]:50806	PostmanRuntime/7.51.0	2026-01-13 17:05:35.427567	\N
7	5	f24824c1-d65d-4511-b170-7bdb9eef7650	2026-01-15 17:33:03.128146	[::1]:64429	PostmanRuntime/7.51.0	2026-01-14 17:33:03.134589	\N
8	5	cb796023-1fe8-4853-8128-a60cba68186a	2026-01-16 12:53:47.300602	[::1]:56550	PostmanRuntime/7.51.0	2026-01-15 12:53:47.308198	\N
9	5	090f12ec-d1c5-43ab-ab15-fae9bbfa154c	2026-01-17 13:50:09.246163	[::1]:50822	PostmanRuntime/7.51.0	2026-01-16 13:50:09.252521	\N
10	7	48b4c975-8281-407e-9c98-2733bb46663e	2026-01-17 14:47:12.820312	[::1]:52999	PostmanRuntime/7.51.0	2026-01-16 14:47:12.831046	\N
11	7	fb948baf-8c6b-4f0b-af9c-3b9315d7f2e0	2026-01-17 14:47:38.264936	[::1]:52999	PostmanRuntime/7.51.0	2026-01-16 14:47:38.265896	\N
12	8	994c80fc-46d1-4e0b-b1d1-3d31a8a25ca7	2026-01-17 16:01:13.249984	[::1]:55506	PostmanRuntime/7.51.0	2026-01-16 16:01:13.251566	2026-01-16 16:09:05.589471
13	8	528bef45-db37-4900-8fc9-aa1a13aacaa0	2026-01-17 16:10:48.579582	[::1]:55914	PostmanRuntime/7.51.0	2026-01-16 16:10:48.580466	\N
\.


--
-- Data for Name: showtimes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.showtimes (id, cinema_id, movie_id, show_date, show_time, price, created_at, updated_at) FROM stdin;
1	1	1	2026-01-20	19:00:00	50000.00	2026-01-12 23:25:34.874601	2026-01-12 23:25:34.874601
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, username, email, password, full_name, phone_number, created_at, updated_at, is_verified, email_verified_at) FROM stdin;
1	lennon	lennon@mail.com	$2a$10$5hzAofJncqAHacZe7dnwq.46hEbpYNCvlqekZw.7SPZMn7QogmBIe			2026-01-10 21:59:52.99487	2026-01-10 21:59:52.99487	f	\N
2	bahlil	bahlil@mail.com	$2a$10$PwtezUD63bZEP/UKLKrAz.Z70H4GPTiY1NMhjxDz1xBTcwQCvpy9.	John Bahlil	081234568870	2026-01-10 22:11:47.43619	2026-01-10 22:11:47.43619	f	\N
3	bahlul	bahlul@mail.com	$2a$10$2cD0SbqdTCiUpDukm10qre3r5zwKuqcbqDWNz.TmWTA3IXzhyWvJa	John Bahlul	081234568871	2026-01-11 00:26:18.792011	2026-01-11 00:26:18.792011	f	\N
4	bahlol	bahlol@mail.com	$2a$10$UupQYShz5L7oLhGM/bPwPefxNntcln.aEQku1H8ZmsDgb8HPCC85u	John Bahlol	081234568877	2026-01-11 21:08:50.746829	2026-01-11 21:08:50.746829	f	\N
5	mccartney	mccartney@mail.com	$2a$10$upKtvO06WE9K5i7MbAMc9eWhbFy72LxxSV2gFcx8vjGsEbBQqP/0u	Mc Cartney	081234568800	2026-01-12 13:44:24.458367	2026-01-12 13:44:24.458367	f	\N
6	ringgo	ringgo@mail.com	$2a$10$oGF/C1dDUrIC.Mx5V6WP4evFJTj7gFnPEwq1cxw6mKRxE6SwL2rSm	Ringgo Stars	081234568100	2026-01-16 13:50:41.64513	2026-01-16 13:50:41.64513	f	\N
7	jokowi	jokowi@mail.com	$2a$10$Ueqn1CbavXNtXdeRrgSte.dADosqzpsQmjwVdwJkEeef/6.ZZMXpO	Jokowi Stars	081234568120	2026-01-16 14:01:43.982927	2026-01-16 14:24:59.071046	f	2026-01-16 14:24:59.071046
8	gibran	gibran@mail.com	$2a$10$lE0pLRjQxNVRb0C2z2sB2eliu1Ivue9P7nVL9gSuVatBZKm3oLxga	Gibran Stars	081234568122	2026-01-16 15:45:26.428015	2026-01-16 16:00:17.780178	t	2026-01-16 16:00:17.780178
\.


--
-- Name: bookings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bookings_id_seq', 4, true);


--
-- Name: cinemas_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.cinemas_id_seq', 6, true);


--
-- Name: movies_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.movies_id_seq', 1, true);


--
-- Name: otp_verifications_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.otp_verifications_id_seq', 3, true);


--
-- Name: payment_methods_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.payment_methods_id_seq', 3, true);


--
-- Name: payments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.payments_id_seq', 8, true);


--
-- Name: seats_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.seats_id_seq', 5, true);


--
-- Name: sessions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sessions_id_seq', 13, true);


--
-- Name: showtimes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.showtimes_id_seq', 1, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 8, true);


--
-- Name: bookings bookings_booking_code_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT bookings_booking_code_key UNIQUE (booking_code);


--
-- Name: bookings bookings_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT bookings_pkey PRIMARY KEY (id);


--
-- Name: bookings bookings_showtime_id_seat_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT bookings_showtime_id_seat_id_key UNIQUE (showtime_id, seat_id);


--
-- Name: cinemas cinemas_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cinemas
    ADD CONSTRAINT cinemas_pkey PRIMARY KEY (id);


--
-- Name: movies movies_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.movies
    ADD CONSTRAINT movies_pkey PRIMARY KEY (id);


--
-- Name: otp_verifications otp_verifications_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.otp_verifications
    ADD CONSTRAINT otp_verifications_pkey PRIMARY KEY (id);


--
-- Name: payment_methods payment_methods_code_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payment_methods
    ADD CONSTRAINT payment_methods_code_key UNIQUE (code);


--
-- Name: payment_methods payment_methods_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payment_methods
    ADD CONSTRAINT payment_methods_pkey PRIMARY KEY (id);


--
-- Name: payments payments_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payments
    ADD CONSTRAINT payments_pkey PRIMARY KEY (id);


--
-- Name: payments payments_transaction_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payments
    ADD CONSTRAINT payments_transaction_id_key UNIQUE (transaction_id);


--
-- Name: seats seats_cinema_id_seat_number_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.seats
    ADD CONSTRAINT seats_cinema_id_seat_number_key UNIQUE (cinema_id, seat_number);


--
-- Name: seats seats_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.seats
    ADD CONSTRAINT seats_pkey PRIMARY KEY (id);


--
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);


--
-- Name: sessions sessions_token_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_token_key UNIQUE (token);


--
-- Name: showtimes showtimes_cinema_id_movie_id_show_date_show_time_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.showtimes
    ADD CONSTRAINT showtimes_cinema_id_movie_id_show_date_show_time_key UNIQUE (cinema_id, movie_id, show_date, show_time);


--
-- Name: showtimes showtimes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.showtimes
    ADD CONSTRAINT showtimes_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- Name: idx_bookings_showtime; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bookings_showtime ON public.bookings USING btree (showtime_id);


--
-- Name: idx_bookings_user; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bookings_user ON public.bookings USING btree (user_id);


--
-- Name: idx_otp_user; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_otp_user ON public.otp_verifications USING btree (user_id);


--
-- Name: idx_payments_booking; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_payments_booking ON public.payments USING btree (booking_id);


--
-- Name: idx_seats_cinema; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_seats_cinema ON public.seats USING btree (cinema_id);


--
-- Name: idx_sessions_token; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sessions_token ON public.sessions USING btree (token);


--
-- Name: idx_sessions_user_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sessions_user_id ON public.sessions USING btree (user_id);


--
-- Name: idx_showtimes_cinema_date; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_showtimes_cinema_date ON public.showtimes USING btree (cinema_id, show_date);


--
-- Name: bookings bookings_seat_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT bookings_seat_id_fkey FOREIGN KEY (seat_id) REFERENCES public.seats(id) ON DELETE CASCADE;


--
-- Name: bookings bookings_showtime_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT bookings_showtime_id_fkey FOREIGN KEY (showtime_id) REFERENCES public.showtimes(id) ON DELETE CASCADE;


--
-- Name: bookings bookings_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT bookings_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: otp_verifications otp_verifications_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.otp_verifications
    ADD CONSTRAINT otp_verifications_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: payments payments_booking_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payments
    ADD CONSTRAINT payments_booking_id_fkey FOREIGN KEY (booking_id) REFERENCES public.bookings(id) ON DELETE CASCADE;


--
-- Name: payments payments_payment_method_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payments
    ADD CONSTRAINT payments_payment_method_id_fkey FOREIGN KEY (payment_method_id) REFERENCES public.payment_methods(id);


--
-- Name: seats seats_cinema_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.seats
    ADD CONSTRAINT seats_cinema_id_fkey FOREIGN KEY (cinema_id) REFERENCES public.cinemas(id) ON DELETE CASCADE;


--
-- Name: sessions sessions_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: showtimes showtimes_cinema_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.showtimes
    ADD CONSTRAINT showtimes_cinema_id_fkey FOREIGN KEY (cinema_id) REFERENCES public.cinemas(id) ON DELETE CASCADE;


--
-- Name: showtimes showtimes_movie_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.showtimes
    ADD CONSTRAINT showtimes_movie_id_fkey FOREIGN KEY (movie_id) REFERENCES public.movies(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

