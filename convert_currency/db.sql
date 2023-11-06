create database convert;

CREATE TABLE public.crypto (
	id uuid NULL DEFAULT gen_random_uuid(),
	cryptoid varchar NULL,
	cryptocode varchar NULL,
	"name" varchar NULL,
	symbol varchar NULL,
	"decimal" int8 NULL,
	address varchar NULL,
	thumblogo varchar NULL,
	smalllogo varchar NULL,
	biglogo varchar NULL,
	chainid varchar NULL,
	chainname varchar NULL,
	description text NULL,
	score int8 NULL,
	socials jsonb NULL,
	website varchar NULL,
	explorer varchar NULL,
	multichain jsonb NULL,
	marketcapusd float8 NULL,
	totalsupply varchar NULL,
	priceusd float8 NULL,
	createddate timestamp NULL DEFAULT now(),
	updateddate timestamp NULL DEFAULT now(),
	totallpusd numeric NULL,
	"type" varchar NULL,
	pricepercentchange24h float8 NULL,
	totalvolume float8 NULL,
	high24h float8 NULL,
	low24h float8 NULL,
	pricechange24h float8 NULL,
	pricechangepercentage24h float8 NULL,
	marketcapchange24h float8 NULL,
	marketcapchangepercentage24h float8 NULL,
	ath float8 NULL,
	athchangepercentage float8 NULL,
	athdate varchar NULL,
	atl float8 NULL,
	atlchangepercentage float8 NULL,
	atldate varchar NULL
);
CREATE INDEX crypto_address_chainname_idx ON public.crypto USING btree (address, chainname);
CREATE INDEX crypto_chainname_idx ON public.crypto USING btree (chainname);
CREATE INDEX crypto_cryptocode_idx ON public.crypto USING btree (cryptocode);
CREATE INDEX crypto_cryptoid_idx ON public.crypto USING btree (cryptoid);
CREATE INDEX crypto_idex_crypto_createddate ON public.crypto USING btree (createddate);
CREATE INDEX crypto_marketcapusd_desc_idx ON public.crypto USING btree (marketcapusd DESC NULLS LAST);
CREATE INDEX crypto_name_idx ON public.crypto USING btree (name);
CREATE INDEX crypto_priceusd_desc_idx ON public.crypto USING btree (priceusd DESC NULLS LAST);
CREATE INDEX crypto_score_desc_idx ON public.crypto USING btree (score DESC NULLS LAST);
CREATE INDEX crypto_total_lp_usd_idx ON public.crypto USING btree (totallpusd);


-- public.currency definition

-- Drop table

-- DROP TABLE public.currency;

CREATE TABLE public.currency (
	"name" varchar NULL,
	symbol varchar NULL,
	image varchar NULL,
	createddate timestamp NULL,
	updateddate timestamp NULL
);