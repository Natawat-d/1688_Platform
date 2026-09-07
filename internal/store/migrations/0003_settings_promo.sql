-- Whether customer quotes may use a supplier's promotional price.
--
-- The catalogue carries a promotionPrice next to every wholesale price, but the
-- documentation never says an order preview will honour it, and against the
-- gateway it did not: we quoted a tier's promotional price, the preview charged
-- the wholesale price, and the difference would have come out of our margin.
-- Off by default. Turn it on only once live previews are seen to apply it.
INSERT INTO settings (key, value) VALUES ('use_promo_prices', '0')
ON CONFLICT (key) DO NOTHING;
