--
-- PostgreSQL database dump
--

--
-- Name: acessolog; Type: TABLE; Schema: public; Owner: gopher
--

CREATE TABLE public.acessolog (
    acess_id serial NOT NULL,
    acess_ip character varying(200) NOT NULL,
    acess_user_up character varying(200) NOT NULL,
    acess_user character varying(200) NOT NULL,
    acess_count integer NOT NULL,
    acess_remoto character varying(200) NOT NULL,
    acess_client_msg character varying(200) NOT NULL,
    acess_client integer NOT NULL,
    logi_data_criacao date DEFAULT current_date NOT NULL
);


ALTER TABLE public.acessolog OWNER TO gopher;

--
-- Name: acessolog acessolog_pkey; Type: CONSTRAINT; Schema: public; Owner: gopher
--

ALTER TABLE ONLY public.acessolog
    ADD CONSTRAINT acessolog_pkey PRIMARY KEY (acess_id);

