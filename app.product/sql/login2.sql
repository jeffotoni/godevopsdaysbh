--
-- PostgreSQL database dump
--

--
-- Name: user; Type: TABLE; Schema: public; Owner: gopher
--
-- SELECT uuid_generate_v4();

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE public.login2 (
    logi_id serial NOT NULL,
    logi_uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
	logi_email character varying(200) NOT NULL,
    logi_nome character varying(200) NOT NULL,
    logi_last_name character varying(200) NOT NULL,
    logi_avatar_dominio character varying(200) NOT NULL,
    logi_senha character varying(200) NOT NULL,
    logi_avatar_type character varying(200) NOT NULL,
    logi_data_criacao date DEFAULT current_date NOT NULL,
    login_codemp integer DEFAULT 10 NOT NULL,
    ativo boolean DEFAULT true
);


ALTER TABLE public.login2 OWNER TO gopher;

--
-- Name: user user_pkey; Type: CONSTRAINT; Schema: public; Owner: gopher
--

ALTER TABLE ONLY public.login2
    ADD CONSTRAINT user_pkey PRIMARY KEY (logi_id);

--
-- Name: fki_codemp; Type: INDEX; Schema: public; Owner: gopher
--

CREATE INDEX fki_email ON public.login2 USING btree (logi_email);

