CREATE OR REPLACE PROCEDURE buy_init_stock(
    p_user_id INTEGER,
    p_stock_id INTEGER,
    p_buy_amt INTEGER
)
AS $$
DECLARE
    v_total_price INTEGER;
    v_stock_init_amt INTEGER;
    v_user_funds DECIMAL(10,0);
BEGIN
    SELECT remain_stock INTO v_stock_init_amt
    FROM stocks WHERE id = p_stock_id;

    SELECT p_buy_amt * current_price INTO v_total_price
    FROM stocks WHERE id = p_stock_id;

    SELECT funds INTO v_user_funds
    FROM users WHERE user_id = p_user_id;

    --check if init_stock have enough to buy and if user have enough money
    IF v_stock_init_amt >= p_buy_amt AND v_total_price <= v_user_funds THEN
        --subtract init_stock from stocks
        UPDATE stocks
        SET remain_stock = remain_stock - p_buy_amt
        WHERE id = p_stock_id;

        --subtract funds from user
        UPDATE users
        SET funds = funds - v_total_price
        WHERE user_id = p_user_id;

        --upsert into holds
        INSERT INTO holds (user_id, stock_id, hold_no)
        VALUES (p_user_id, p_stock_id, p_buy_amt)
        ON CONFLICT (user_id, stock_id)
        DO UPDATE SET hold_no = holds.hold_no + excluded.hold_no;
    ELSE
        RAISE EXCEPTION 'Not enough stock to buy or not enough money';
    END IF;
EXCEPTION
    WHEN OTHERS THEN
        -- Rethrow the exception for handling at a higher level
        RAISE;
END;$$
LANGUAGE plpgsql;