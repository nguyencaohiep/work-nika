-- public.fund definition

-- Drop table

-- DROP TABLE public.fund;

CREATE TABLE public.fund (
	id uuid NULL DEFAULT gen_random_uuid(),
	projectid varchar NULL,
	ventureid varchar NULL,
	fundstagecode varchar NULL,
	fundstagename varchar NULL,
	fundamount float8 NULL,
	fundamountusd float8 NULL,
	funddate timestamp NULL,
	description text NULL,
	announcementurl varchar NULL,
	valuation float8 NULL,
	sourceurl varchar NULL,
	createddate timestamp NULL DEFAULT now(),
	updateddate timestamp NULL DEFAULT now(),
	projectlogo varchar NULL,
	projectcode varchar NULL,
	projectname varchar NULL
);
CREATE INDEX fund_fundamount_idx ON public.fund USING btree (fundamount);
CREATE INDEX fund_projectid_idx ON public.fund USING btree (projectid);
CREATE INDEX fund_ventureid_fundstagecode_idx ON public.fund USING btree (ventureid, fundstagecode);
CREATE INDEX fund_ventureid_idx ON public.fund USING btree (ventureid);

-- public.venture definition

-- Drop table

-- DROP TABLE public.venture;

CREATE TABLE public.venture (
	id uuid NULL DEFAULT gen_random_uuid(),
	ventureid varchar NULL,
	venturesrc varchar NULL,
	venturecode varchar NULL,
	venturename varchar NULL,
	venturelogo varchar NULL,
	yearfounded int8 NULL,
	"location" varchar NULL,
	description text NULL,
	socials json NULL,
	sourceurl varchar NULL,
	createddate timestamp NULL DEFAULT now(),
	updateddate timestamp NULL DEFAULT now(),
	star float8 NULL DEFAULT 0,
	"notice" varchar NULL,
	website varchar NULL,
	reputation int8 NULL,
	subcategory varchar NULL,
	score int8 NULL,
	totalfund int8 NULL,
	statistics_category json NULL,
	statistics_month json NULL,
	statistics_country json NULL
);
CREATE INDEX venture_location_desc_idx ON public.venture USING btree (location DESC NULLS LAST);
CREATE INDEX venture_score_idx ON public.venture USING btree (score);
CREATE INDEX venture_ventureid_idx ON public.venture USING btree (ventureid);
CREATE INDEX venture_venturename_idx ON public.venture USING btree (venturename);
CREATE INDEX venture_yearfounded_desc_idx ON public.venture USING btree (yearfounded DESC NULLS LAST);



CREATE OR REPLACE VIEW public.fund_round
AS SELECT fund.ventureid AS fventureid,
    sum(
        CASE
            WHEN fund.fundstagecode::text = 'seed'::text THEN fund.fundamount
            ELSE 0::double precision
        END) AS seed,
    sum(
        CASE
            WHEN fund.fundstagecode::text = 'series-a'::text THEN fund.fundamount
            ELSE 0::double precision
        END) AS seriesa,
    sum(
        CASE
            WHEN fund.fundstagecode::text = 'series-b'::text THEN fund.fundamount
            ELSE 0::double precision
        END) AS seriesb,
    sum(
        CASE
            WHEN fund.fundstagecode::text = 'series-c'::text THEN fund.fundamount
            ELSE 0::double precision
        END) AS seriesc,
    sum(
        CASE
            WHEN fund.fundstagecode::text = 'strategic'::text THEN fund.fundamount
            ELSE 0::double precision
        END) AS strategic,
    sum(
        CASE
            WHEN fund.fundstagecode::text = 'ico'::text THEN fund.fundamount
            ELSE 0::double precision
        END) AS ico,
    sum(fund.fundamount) AS totalfund
   FROM fund
  GROUP BY fund.ventureid;



CREATE OR REPLACE VIEW public.compact_venture
AS SELECT (venture.ventureId), venture.ventureName, venture.ventureLogo, 
venture.yearFounded, venture.location, venture.description, 
venture.star, venture.notice, venture.subcategory,  venture.score,
fund_round.fventureid,
fund_round.seed,
fund_round.seriesa,
fund_round.seriesb,
fund_round.strategic,
fund_round.seriesc,
fund_round.ico,
fund_round.totalfund 
from venture left join fund_round on venture.ventureid = fund_round.fventureid;