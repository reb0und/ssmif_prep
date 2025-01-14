USE ssmif_db;

CREATE Table data (
	ticker String,
	price Float32,
	timestamp DateTime DEFAULT now(),
) Engine = MergeTree()
PARTITION BY toYYYYMMDD(timestamp)
ORDER BY (ticker, timestamp);
