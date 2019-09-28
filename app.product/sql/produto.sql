--
-- PostgreSQL database dump
--

--
-- Name: produto; Type: TABLE; Schema: public; Owner: gopher
--

CREATE TABLE public.produto (
    codprod character varying(20) NOT NULL,
    descrprod character varying(150) NOT NULL,
    caracteristicas text,
    ativo boolean,
    excluido boolean,
    estoque numeric,
    preco numeric,
    preco_promocional numeric,
    custo numeric,
    peso numeric,
    altura numeric,
    largura numeric,
    profundidade numeric,
    codgrupoprod integer NOT NULL,
    marca character varying(40),
    ean character varying(13),
    nbm character varying(20),
    codemp integer NOT NULL
);


ALTER TABLE public.produto OWNER TO gopher;

--
-- Name: produto produto_pkey; Type: CONSTRAINT; Schema: public; Owner: gopher
--

ALTER TABLE ONLY public.produto
    ADD CONSTRAINT produto_pkey PRIMARY KEY (codprod, codemp);

--
-- Name: fki_codemp; Type: INDEX; Schema: public; Owner: gopher
--

CREATE INDEX fki_codemp ON public.produto USING btree (codemp);



