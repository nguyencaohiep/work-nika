CREATE TABLE reviews (
	id bigserial NOT NULL,
	accountid bigint NULL,
	"content" text NULL,
	star int8 NULL,
	productid bigint NOT NULL,
	createddate timestamp NOT NULL,
	updateddate timestamp NOT NULL
);
CREATE INDEX idx_reviews_accountid ON reviews (accountid);
CREATE INDEX idx_reviews_productid ON reviews (productid);
CREATE INDEX idx_reviews_star ON review (star);

CREATE TABLE replies (
	id bigserial NOT NULL,
	reviewid bigint NULL,
	accountid bigint NULL,
	content text NULL,
    productid bigint NOT NULL,
	createddate timestamp NOT NULL,
	updateddate timestamp NOT NULL
);
CREATE INDEX idx_replies_reviewid ON replies (reviewid);
CREATE INDEX idx_replies_accountid ON replies (accountid);
CREATE INDEX idx_replies_productid ON replies (productid);

CREATE TABLE reactions (
	commentid bigint NOT NULL,
	type bigint NOT NULL,
	accountid bigint NOT NULL,
	reactiontype varchar NOT NULL,
    productid bigint NOT NULL,
	createddate timestamp NOT NULL,
	updateddate timestamp NOT NULL
);
CREATE INDEX idx_reactions_commentid ON reactions (commentid);
CREATE INDEX idx_reactions_type ON reactions (type);
CREATE INDEX idx_replies_accountid ON reactions (accountid);
CREATE INDEX idx_reactions_productid ON reactions (productid);