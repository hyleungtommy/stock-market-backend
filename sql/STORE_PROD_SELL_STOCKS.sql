CREATE OR REPLACE PROCEDURE sell_stock(
    p_user_id INTEGER,
    p_stock_id INTEGER,
    p_sell_amount INTEGER,
    p_selling_price DECIMAL(10, 2)
)
AS $$
DECLARE
    v_hold_no INTEGER;
BEGIN
    -- Retrieve the hold_no from the holds table
    SELECT hold_no INTO v_hold_no
    FROM holds
    WHERE user_id = p_user_id AND stock_id = p_stock_id;

    -- Check if there's enough stock to sell
    IF v_hold_no >= p_sell_amount THEN
        -- Subtract the sold stock from the hold_no
        UPDATE holds
        SET hold_no = hold_no - p_sell_amount
        WHERE user_id = p_user_id AND stock_id = p_stock_id;

        -- Insert a new transaction record
        INSERT INTO transactions (user_id, stock_id, sell_amt, sell_price)
        VALUES (p_user_id, p_stock_id, p_sell_amount, p_selling_price);

    ELSE
        RAISE EXCEPTION 'Not enough stock to sell';
    END IF;
EXCEPTION
    WHEN OTHERS THEN
        -- Rethrow the exception for handling at a higher level
        RAISE;
END;$$
LANGUAGE plpgsql;