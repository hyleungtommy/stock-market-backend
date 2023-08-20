CREATE OR REPLACE PROCEDURE buy_stock(
    p_buyer_user_id INTEGER,
    p_transaction_id INTEGER
)
AS $$
DECLARE
    v_buyer_funds DECIMAL(10,0);
    v_seller_user_id INTEGER;
    v_sell_price DECIMAL(10,0);
    v_stock_id INTEGER;
    v_sell_amt INTEGER;
BEGIN
    
    SELECT funds INTO v_buyer_funds
    FROM users WHERE user_id = p_buyer_user_id;

    SELECT sell_price * sell_amt, user_id, sell_amt, stock_id INTO v_sell_price, v_seller_user_id, v_sell_amt,v_stock_id
    FROM transactions WHERE id = p_transaction_id;
    
    --check if buyer has enough money to buy
    IF v_buyer_funds >= v_sell_price THEN
        --add the amount of stock into buyer's holds
        INSERT INTO holds (user_id, stock_id, hold_no)
        VALUES (p_buyer_user_id, v_stock_id, v_sell_amt)
        ON CONFLICT (user_id, stock_id)
        DO UPDATE SET hold_no = holds.hold_no + excluded.hold_no;

        --add the amount of money into seller's funds
        UPDATE users
        SET funds = funds + v_sell_price
        WHERE user_id = v_seller_user_id;

        --subtract the amount of money from buyer's funds
        UPDATE users
        SET funds = funds - v_sell_price
        WHERE user_id = p_buyer_user_id;

        --insert into transaction log
        INSERT INTO transaction_logs  (user_id, stock_id, sell_price, sell_amt)
        SELECT user_id, stock_id, sell_price, sell_amt
        FROM transactions
        WHERE id = p_transaction_id;

        --remove the transaction
        DELETE FROM transactions
        WHERE id = p_transaction_id;

    ELSE
        RAISE EXCEPTION 'User does not have enough money to buy the stock';
    END IF;
EXCEPTION
    WHEN OTHERS THEN
        -- Rethrow the exception for handling at a higher level
        RAISE;
END;$$
LANGUAGE plpgsql;