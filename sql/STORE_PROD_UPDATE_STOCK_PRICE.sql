CREATE OR REPLACE PROCEDURE update_stock_price()
AS $$
DECLARE
    stock_cursor CURSOR For
        SELECT id FROM stocks;
    v_avg_price DECIMAL(10, 2);
    v_stock_id INTEGER;
BEGIN
    OPEN stock_cursor;
    -- For each stock
    LOOP
        FETCH stock_cursor INTO v_stock_id;
        EXIT WHEN NOT FOUND;
    -- Get all transaction history from previous day and calculate the average price
        SELECT AVG(sell_price) INTO v_avg_price
        FROM transaction_logs
        WHERE transaction_date >= now() - interval '24 hours' AND stock_id = v_stock_id;
    -- Update stock current price
        IF v_avg_price IS NOT NULL THEN
            UPDATE stocks SET current_price = v_avg_price
            WHERE id = v_stock_id;
        END IF;
    END LOOP;

    CLOSE stock_cursor;
EXCEPTION
    WHEN OTHERS THEN
        -- Rethrow the exception for handling at a higher level
        RAISE;
END;$$
LANGUAGE plpgsql;