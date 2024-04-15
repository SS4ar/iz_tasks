

CREATE TABLE public.products (
    id integer NOT NULL,
    name text NOT NULL,
    price numeric(10,2) NOT NULL,
    image_url text,
    description text
);


ALTER TABLE public.products OWNER TO postgres;

CREATE SEQUENCE public.products_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.products_id_seq OWNER TO postgres;

ALTER SEQUENCE public.products_id_seq OWNED BY public.products.id;

CREATE TABLE public.user_products (
    id integer NOT NULL,
    user_id integer NOT NULL,
    product_id integer NOT NULL,
    purchase_price numeric(10,2)
);


ALTER TABLE public.user_products OWNER TO postgres;

CREATE SEQUENCE public.user_products_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_products_id_seq OWNER TO postgres;

ALTER SEQUENCE public.user_products_id_seq OWNED BY public.user_products.id;

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(50) NOT NULL,
    password text NOT NULL,
    balance numeric(10,2)
);


ALTER TABLE public.users OWNER TO postgres;

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;

ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_id_seq'::regclass);

ALTER TABLE ONLY public.user_products ALTER COLUMN id SET DEFAULT nextval('public.user_products_id_seq'::regclass);

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);

COPY public.products (id, name, price, image_url, description) FROM stdin;
1	Mystic	14.00	/resources/1.jpg	One who seeks a deeper understanding of the divine through mystical experiences.
2	Poetic	11.00	/resources/9.jpg	Someone who expresses themselves in a creative and imaginative way through language.
3	Philosophical	16.00	/resources/10.jpg	A person who engages in the study and exploration of the fundamental nature of knowledge, reality, and existence.
4	Spiritual	9.00	/resources/11.jpg	Someone who seeks a connection with the divine or higher power beyond the physical world.
5	Holy	10.00	/resources/12.jpg	A person or object that is considered sacred or blessed by a particular religion or belief system.
6	Teacher	10.00	/resources/13.jpg	Someone who imparts knowledge and skills to others through instruction and guidance.
7	flag	1337.00	/resources/8.jpg	kubanctf{d0_u_11k3_c4py84r45_0r_83av32s?}
8	Sage	23.00	/resources/2.jpg	An individual who possesses great wisdom and insight gained through experience and contemplation.
9	Prophet	16.00	/resources/3.jpg	A person who receives divine revelations and communicates them to others.
10	Yogi	5.00	/resources/4.jpg	Someone who practices yoga as a means of achieving spiritual and physical well-being.
11	Sufi	10.00	/resources/5.jpg	A Muslim mystic who seeks a direct personal experience of the divine through prayer and meditation.
12	Tantric	30.00	/resources/6.jpg	A practitioner of tantra, a spiritual path that emphasizes the integration of mind, body, and spirit.
13	Ascetic	19.00	/resources/7.jpg	Someone who renounces worldly pleasures and possessions in pursuit of spiritual enlightenment.
\.

SELECT pg_catalog.setval('public.products_id_seq', 10, true);


SELECT pg_catalog.setval('public.user_products_id_seq', 32, true);


SELECT pg_catalog.setval('public.users_id_seq', 43, true);

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.user_products
    ADD CONSTRAINT user_products_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);



ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


ALTER TABLE ONLY public.user_products
    ADD CONSTRAINT user_products_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id);


ALTER TABLE ONLY public.user_products
    ADD CONSTRAINT user_products_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);

